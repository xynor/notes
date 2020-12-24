## sql执行顺序
* from子句组装来自不同的数据源
* where子句基于指定条件对记录进行筛选
* group by子句将数据划分为多个组
* 使用聚集函数进行计算
* 使用having子句筛选分组(增加 HAVING 子句原因是，WHERE 关键字无法与合计函数一起使用) `
SELECT TeacherID, AVG(Age) AS AGE FROM Student
GROUP BY TeacherID
HAVING AVG(Age) > 12`
* 计算所有的表达式
* select字段
* 使用order by对数据进行排序
* 每一步产生一个虚拟表做为下一步的输入，虚拟表对客户端不可见 

## exits和in,not exits和not in
* exits是返回bool，loop外查询，再逐一判断是否返回true
* in是返回列，先查询in再在外查询中匹配，in有Null时报错
* 当外表大，内表小，用in效率更高。外表小，内表大，用exits效率高
* 无论哪个表大，用not exits比not in效率高，应为not in用全表扫描