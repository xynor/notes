## 模型
````
func handleConn(c net.Conn) {
    defer c.Close()
    for {
        // read from the connection
        // ... ...
        // write to the connection
        //... ...
    }
}

func main() {
    l, err := net.Listen("tcp", ":8888")
    if err != nil {
        fmt.Println("listen error:", err)
        return
    }

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
            break
        }
        // start a new goroutine to handle
        // the new connection.
        go handleConn(c)
    }
}
````
* socket都是non-block的，与runtime在逻辑上实现了以阻塞的方式写非阻塞的代码
## TCP的建立
* net.Dial与带超时的net.DialTimeout
#### 网络不可达或服务未启动
* 很快就会返回connection refused
#### 对方的listen backlog满
* server比较忙，来不及accept，导致backlog满了(connect是在listen queue中完成的，accept是从queue中取出一个conn)，导致Dial阻塞
* client可以一次性建立(somaxconn)个链接，当server队列满了之后，Dial阻塞(握手不成功)
* server一直不accept，则可能导致Dial超时operation timed out，与平台和内核设置有关
* 网络不好，丢包多导致握手不成功，Dial也会超时
* 如果需要严格感知Dial，则可以用DialTimeout
## Socket读
* 使用上和non-block的socket读类似，没有数据阻塞。有数据未超过buffer，返回buffer和len。有数据超过buffer，分多次返回。类似epoll水平触发
* socket关闭，当对方关闭socket(FIN)，read会先返回读缓冲区还有的数据，随后读到EOF，表示对方关闭
* 读取超时，时间敏感的情况下设置SetReadDeadline，则在规定时间没有读到数据，read返回错误i/o timeout。```是否会发送RST？```
## Socket写
* socket的写队列不满时可以非阻塞返回写了多少字节，当写队列满时，write阻塞
* 当在写的过程中，对方关闭socket，将收到RST报文，如果继续写返回错误broken pipe，则此次的write返回的n需要特殊处理
* 写超时，时间敏感的情况下设置SetWriteDeadline，也存在写入部分数据返回错误和n
* 只有在写出错，broken pipe或写超时的时候才会写入部分数据并且返回错误和n，```正常写是阻塞写完buffer，err==nil,n==len(buffer)```
## net.conn实现了io.Reader和io.Writer接口可以用bufio包下面的Writer和Reader、io/ioutil下的函数等。io.ReadFull(t.conn, preface)
## 并发安全
* 读和写都是有锁保护的，多个协程读可能使一个业务报文读到几个协程中。多个协程写，需要确保不能把一个业务包分别用几个协程写。这会导致业务包不连续
## 关闭链接
* 在己方已经关闭的socket上再读写，会得到use of closed network connection错误
* 在对方关闭socket，己方read会得到EOF
* 对方关闭socket，己方继续写，会收到RST报文，再写会出现broken pipe错误，如果中间路由断开链接，那么只有写超时，keepalive超时能检测
## 思考
* 读写携程分开
* 做了tcp的frame协议，在read的时候使用io.ReadFull(t.conn, preface)读frame长度
* 业务需要写的时候直接发到dequeue中(共享内存)，通知channel，通知写协程从dequeue中拿出报文写socket
## socket对错误的反馈能力是有限的
* 对端异常关闭(OS没有发FIN包)/网络延迟大，如果本端阻塞在read上，那么将永远等待，除非设置超时
* 本端先write了数据，并read等待。重传超时则read会收到ETIMEDOUT/EHOSTUNREACH/ENETUNREACH错误
* 对端恰好恢复，收到本端包。对端不能识别，返回RST，本端read返回ECONNREST
* write错误，最终通过read通知应用有点阴差阳错
* keepalive,应用层心跳
