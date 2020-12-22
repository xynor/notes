## ACID
* 原子性Atomicity
* 一致性Consistency
* 隔离性Isolation
* 持久性Durability
## 默认隔离级别可重复读 Repeatable read
* 脏读，A事务修改了数据但是还没有提交，B事务读取到了A未提交的数据
* 不可重复读，A事务读取到了一行数据，B事务修改了这一行数据并提交，A再读取同样数据前后两次读取不一样，行锁解决
* 幻读(Y)，A事务读取了N行数据，B事务添加了n行数据并提交，A再读取时变成N+n行，导致A插入(与B一样数据)失败，表锁解决
//todo 与锁的关系
## 事务实现原理
* redo log是用来恢复数据的，用于保障已提交事务的持久性(记录已经提交的)
* undo log是用来回滚数据的，用于保障未提交事务的原子性
* mvcc每一行多出两个隐藏字段，创建时间(事务ID)和删除时间(事务ID)
* 查询时，只会查找创建时间小于等于当前时间，保证在查询时行数据存在。和 行的删除版本要么未定义，要么大于当前事务版本，保证行数据在事务开始前未删除
* 通过mvcc解决了select读的幻读问题，但是在B事务删除某行后，A事务再update(写)会出错。
* 读是读的快照片，写需要最新版本数据。如果select要当前读则需要加锁select * from table where ? lock in share mode;select * from table where ? for update;