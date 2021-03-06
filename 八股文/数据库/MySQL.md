### 1、数据库索引的优缺点？

<font color=red>优点：提高数据检索的性能。
缺点：</font>

* <font color=red>索引会占据物理存储空间</font>
* <font color=red>当向表中添加/删除数据时，索引也需动态更新，降低了插入/删除的速度</font>

### 2、数据库的三范式？

* <font color=red>1NF：字段不可分（原子性）</font>
* <font color=red>2NF：有主键，非主键字段依赖主键（唯一性）</font>
* <font color=red>3NF：非主键字段不能相互依赖（每列都与主键有直接关系，不存在传递依赖）</font>

### 3、什么是事务？事务有哪些特性？

<font color=red>【解】事务是指作为单个逻辑工作单元执行的一系列操作，要么完全地执行，要么完全地不执行。</font>

* <font color=red>原子性：要么完全执行，要么完全不执行</font>
* <font color=red>一致性：事务完成时，所有数据保持一致</font>
* <font color=red>隔离性：多个事务作修改时，互相隔离</font>
* <font color=red>持久性：事务所作的修改是永久性的</font>

### 4、MySQL的引擎 InnoDB 和 MyISAM 的区别。

* <font color=red>InnoDB支持外键，MyISAM不支持；</font>
* <font color=red>InnoDB支持事务处理，MyISAM不支持；</font>
* <font color=red>InnoDB是行锁，MyISAM是表锁；</font>
* <font color=red>MyISAM是默认的存储引擎，强调性能。</font>