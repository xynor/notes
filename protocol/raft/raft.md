## 演示
[演示](http://thesecretlivesofdata.com/raft/)  
[论文翻译](https://www.jianshu.com/p/2a2ba021f721?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation)
[论文解析](https://blog.csdn.net/rsy56640/article/details/89116768)    
## 原则
* 每个任期内只有一个Leader。
* Leader负责发送AppendEntries RPC(日志复制，心跳)消息。Leader不会修改和删除日志，只增加。
* Follower只要在没超时能收到AppendEntries RPC消息，就一直是Follower，心跳超时则认为Leader故障，可以竞选下一任期Leader。
* Leader如果收到了比自己term大的AppendEntries RPC消息，则认为自己过期了，转为Follower。
* **写消息需要Leader在操作半数Follower日志提交后再返回给client**。
* **读消息需要Leader确认一次自己是不是最新的Leader，即发送AppendEntries RPC消息得到半数确认**。

## 初始化启动
* A，B，C都是Follower并设置各自的超时时间(150-300ms)。
* 等待超时的时候如果收到了RequestVote RPC，则表示已经有其他节点进入了Candidate。投票给此Candidate，继续保持Follower状态。
* 如果没有收到RequestVote RPC，超时，Follower先增加自己的当前任期号并转成Candidate，投票给自己并广播RequestVote RPC。
* 得到多数投票，转成Leader。发送心跳(不包含日志条目的AppendEntries RPC)来维持自己的地位。
* 如果在成为Candidate并等待投票的时候，收到了AppendEntries RPC消息，说明有其他server成为了Leader。当此server的任期号不小于Candidate的任期号，则肯定其Leader地位，并
回到Follower，否则拒绝此RPC，继续Candidate状态。
* 可能票数被瓜分，没有server成为Leader，那么Candidate超时，进入Follower等待下一次选举(增加任期号)。

## 客户端交互
* Client发送所有的请求给Leader。client启动时随机连接server，如果不是Leader，则拒绝并redirect(AppendEntries 请求包含了 leader 的网络地址)。
* 如果Leader崩溃，则客户端请求超时，随机重试。
* Client写操作需要幂等。在Leader可能commit日志后奔溃没返回给client时发生。所以需要一条指令都赋予一个唯一的序列号。
* 只读需要Leader知道自己是不是最新的。否则可能出现过期数据。这需要在返回前先跟超过半数的Follower交换心跳。

## Leader崩溃
* 其他的Follow收不到心跳消息，超时，任期号加一后进入Candidate，投票给自己并发投票消息。
* Follow或Candidate收到投票消息后，如果任期号比自己小则返回false，如果对比其日志不比自己旧(可能是最新的)则投票。
* 超过半数投票后成为新的Leader。

## 原Leader恢复
* 崩溃的原Leader恢复后，同样先进入Follower，此时收到了Leader的心跳消息。currentTerm是？？？
* 从Leader的日志中同步丢失的entries。

## 分区容忍
* 当网络出现问题，导致网络分区后，比如5个server被分成3，2。失去Leader的多数将重新选择Leader，而少数有Leader的分区因为Leader不能commit日志而失败。
* 分区恢复后，不能工作的分区收到比自己term大的RPC消息而重新同步数据。

## 一些思考
* 适用范围在`broadcastTime ≪ electionTimeout ≪ MTBF（结点平均 crash 时间）`可信的局域网中。
* 不能解决拜占庭问题(包括网络延迟，分区和数据包丢失，重复和乱序)的安全性(不会返回不正确的结果)。