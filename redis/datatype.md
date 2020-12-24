## key必须为字符串
## String
* 简单的key-value,适用于分布式锁，计数器(原子操作),分布式全局id
* 最大512M空间,使用预分配，惰性释放
## List
* 双向链表，最大2^32-1,可以实现栈，队列，有限集合，简单的消息队列
````
redis 127.0.0.1:6379> LPUSH runoobkey redis
(integer) 1
redis 127.0.0.1:6379> LPUSH runoobkey mongodb
(integer) 2
redis 127.0.0.1:6379> LPUSH runoobkey mysql
(integer) 3
redis 127.0.0.1:6379> LRANGE runoobkey 0 10

1) "mysql"
2) "mongodb"
3) "redis"
````
## Hash
* key依然为字符串，里面对应了field和value
````
127.0.0.1:6379>  HMSET runoobkey name "redis tutorial" description "redis basic commands for caching" likes 20 visitors 23000
OK
127.0.0.1:6379>  HGETALL runoobkey
1) "name"
2) "redis tutorial"
3) "description"
4) "redis basic commands for caching"
5) "likes"
6) "20"
7) "visitors"
8) "23000"
````
* 有点像把jasonString分开存储，不能嵌套其他类型
* 与string比可以减少key冲突，减少空间，一次获取相关数据，减少io。但是field不能单独设置超时，没有bit操作
* 数据多时是hash表，少时可以被压缩成数组。
* 是hash就意味着冲突，扩容。使用的类似go sync.Map读写分离和渐进式扩容
## Set
* 简化版Hash，没有field
* 不能重复
## ZSet
* 不能重复，每个value有一个分数，通过分数排序
* 使用了hash和跳表实现
* 可以用作积分排行榜，时间排序新闻，延迟队列
## Redis Geo
````
redis> GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
(integer) 2
redis> GEODIST Sicily Palermo Catania
"166274.1516"
redis> GEORADIUS Sicily 15 37 100 km
1) "Catania"
redis> GEORADIUS Sicily 15 37 200 km
1) "Palermo"
2) "Catania"
redis>
````
## HyperLogLog
* 计算集合的基数，适用于UV统计，错误率大概在 0.81%
* 比set节省内存
## Bitmap
* BitMap 的 offset 值上限 2^32 - 1
* key = 年份：用户id  offset = （今天是一年中的第几天） % （今年的天数）
* 使用日期作为 key，然后用户 id 为 offset 设置不同 offset 为 0 1 即可
## Bloom Filter
* 不存在的一定不存在，存在的不一定存在
* 一个元素加入时，通过K个散列函数映射成k个点(有效进行散列)，就位数组的对应位设置为1
* 查找时，可以快速通过散列到的位是否为0，来认为是否存在

