## GRPC
* 有断链重链，有PING消息(与WINDOW_UPDATE一起)
* 多个协程可以使用一个connection，server收到请求后，首先会在单个协程中完成http2的frame解码，如果是业务请求(dataframe)开启单个协程处理单个请求
````
//google.golang.org/grpc/server.go 
func (s *Server) serveStreams(st transport.ServerTransport) {
    defer st.Close()
    var wg sync.WaitGroup
    // HandleStreams 是注册 grpc server处理 http2 stream 数据的处理函数
    st.HandleStreams(func(stream *transport.Stream) { 
        wg.Add(1)
        //每次有新request时会调用这个方法， 这个方法就是开新的协程处理请求
        go func() {
            defer wg.Done()
            s.handleStream(st, stream, s.traceInfo(st, stream))
        }()  
    }, func(ctx context.Context, method string) context.Context {
        if !EnableTracing {
            return ctx
        }    
        tr := trace.New("grpc.Recv."+methodFamily(method), method)
        return trace.NewContext(ctx, tr)
    })   
    wg.Wait()
}
````
* streamId不可能被重复使用，使用完了会新开一条tcp链接，同一条stream逻辑上是同一次链接，可以close发送
* Handler在服务器上是开了协程调度，意味着server推送可以是个for死循环
* stream的recv和send可以在不同协程中，同一个stream的recv或send不能在不同协程中。多协程调用，需要```stream, err := client.RouteChat(ctx)```建立新的stream
* Handler入参，出参都是message，这是一元调用。入参是stream，出参是message则是client->server。入参是message，出参是stream则是server->client。
出参，入参是stream则是双向的
* 双向流可以是不同的message类型，在stream.Recv和stream.Send中区分类型
* 因为http2有url，所以client,server可以注册多个结构体