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
