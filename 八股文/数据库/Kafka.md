### 1、你说了解kafka，能简单描述一下Kafka吗？能画出它的架构图吗？

<font color=red>Kafka是一个高吞吐、易扩展的分布式发布-订阅消息系统，它能够将消息持久化到磁盘，用于批量的消费。Kafka中有以下几个概念：</font>

* <font color=red>Topic：特指Kafka处理的消息源（feeds of messages）的不同分类。</font>
* <font color=red>Partition：Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）。</font>
* <font color=red>Broker：Kafa集群中包含一台或多台服务器，这种服务器被称为broker。</font>
* <font color=red>Producer：生产者，向Kafka的一个topic发布消息。</font>
* <font color=red>Consumers：消费者，从kafka的某个topic读取消息。</font>

Kafka架构图如下：

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/kafka-structure.jpg?raw=true)

参考：https://www.infoq.cn/article/apache-kafka/

### 2、Kafka 分区的目的？

分区对于 Kafka 集群的好处是：实现负载均衡。
分区对于消费者来说，可以提高并发度，提高效率。

### 3、Kafka 是如何做到消息的有序性？

kafka 中的每个 partition 中的消息在写入时都是有序的，而且单独一个 partition 只能由一个消费者去消费，可以在里面保证消息的顺序性。但是分区之间的消息是不保证有序的。

### 4、Kafka 的高可靠性是怎么实现的？

参考：https://blog.csdn.net/weixin_38246518/article/details/108123878

### 5、Kafka 数据一致性是怎么实现的？

一致性是指不管老的 Leader 还是新选举的 Leader，Consumer 都能读到一样的数据。

kafka是通过 HW(High Water Mark) 机制（木桶原理）来保证数据的一致性。

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/kafka-high-water-mark.jpg?raw=true)

假设分区的副本为3，其中副本0是 Leader，副本1和副本2是 follower，并且在 ISR 列表里面。虽然副本0已经写入了 Message4，但是 Consumer 只能读取到 Message2。因为所有的 ISR 都同步了 Message2，只有 High Water Mark 以上的消息才支持 Consumer 读取，而 High Water Mark 取决于 ISR 列表里面偏移量最小的分区，对应于上图的副本2，这个很类似于木桶原理。

这样做的原因是还没有被足够多副本复制的消息被认为是“不安全”的，如果 Leader 发生崩溃，另一个副本成为新 Leader，那么这些消息很可能丢失了。如果我们允许消费者读取这些消息，可能就会破坏一致性。试想，一个消费者从当前 Leader（副本0） 读取并处理了 Message4，这个时候 Leader 挂掉了，选举了副本1为新的 Leader，这时候另一个消费者再去从新的 Leader 读取消息，发现这个消息其实并不存在，这就导致了数据不一致性问题。

当然，引入了 High Water Mark 机制，会导致 Broker 间的消息复制因为某些原因变慢，那么消息到达消费者的时间也会随之变长（因为我们会先等待消息复制完毕）。延迟时间可以通过参数 replica.lag.time.max.ms 参数配置，它指定了副本在复制消息时可被允许的最大延迟时间。

### 6、消费者和消费者组有什么关系？

每个消费者从属于消费组。




