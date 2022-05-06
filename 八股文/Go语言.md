

### 1、简单说说协程底层原理 GPM 模型

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/go-GPM1.png?raw=true)

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/go-GPM2.jpeg?raw=true)

调度器的设计策略：复用线程 ——> 避免频繁的创建、销毁线程，而是对线程的复用。

1. work stealing机制 
   * 当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程。

2. hand off机制 
   * 当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。

利用并行： GOMAXPROCS 设置P的数量，最多有 GOMAXPROCS 个线程分布在多个CPU上同时运行。 GOMAXPROCS 也限制了并发的程度，比如 GOMAXPROCS = 核数/2 ，则最多利用了一半的CPU核进行并行。

抢占：在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在Go中，一个goroutine最多占用CPU 10ms，防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方。

全局G队列：在新的调度器中依然有全局G队列，但功能已经被弱化了，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G。

参考：https://learnku.com/articles/41728

### 2、哪些方式可以实现线程安全的 map

* map+读写锁
* sync.Map

参考：https://zhuanlan.zhihu.com/p/449078860

### 3、golang 闭包

闭包其实就是匿名函数的调用

## 基础

### 1、golang 中 make 和 new 的区别？

简单的说，new只分配内存，make用于slice，map，和channel的初始化。

* make和new都是golang用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。
* make返回的还是引用类型本身；而new返回的是指向类型的指针。
* make只能用来分配及初始化类型为slice，map，channel的数据；new可以分配任意类型的数据。

### 2、slice底层怎么实现的？怎么扩容的？是在堆上还是栈上？

切片(slice)是 Golang 中一种比较特殊的数据结构，这种数据结构更便于使用和管理数据集合。切片是围绕动态数组的概念构建的，可以按需自动增长和缩小。

切片是一个很小的对象，它对底层的数组(内部是通过数组保存数据的)进行了抽象，并提供相关的操作方法。切片是一个有三个字段的数据结构，这些数据结构包含 Golang 需要操作底层数组的元数据：

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/slice-struct.jpg?raw=true)

可以通过`slice[i:j]`从一个切片创建新的切片：

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/slice-create-another.jpg?raw=true)

需要注意的是：现在两个切片**共享同一个底层数组**。

**切片扩容**：如果切片的底层数组没有足够的可用容量，append() 函数会创建一个新的底层数组，将被引用的现有的值复制到新数组里，再追加新的值。

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/slice-expand.jpg?raw=true)

函数 append() 会智能地处理底层数组的容量增长。在切片的容量小于 1000 个元素时，总是会成倍地增加容量。一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25%的容量(随着语言的演化，这种增长算法可能会有所改变)。

go语言中的容器类型，包括Slice、Map、Heap、List、Ring，底层存储都是在堆上。

参考：https://blog.csdn.net/stpeace/article/details/100695000

### 3、数组和切片的区别？

* 定义方式不一样
* 初始化方法不一样：
    * 数组需要指定大小，不指定也会根据初始化的自动推算出大小，不可改变
    * 切片不需要指定大小。
* 函数传递方式不同：数组是值传递，切片是地址传递。

### 4、变量在什么情况下会逃逸

参考：https://zhuanlan.zhihu.com/p/441593663

## context相关

### 1、context 结构是什么样的？

参考：https://blog.csdn.net/chinawangfei/article/details/86559975

### 2、context 使用场景和用途

## channel相关

### 1、channel 是否线程安全？锁用在什么地方？

Golang的Channel,发送一个数据到Channel 和 从Channel接收一个数据都是**原子性**的。
而且Go的设计思想就是:不要通过共享内存来通信，而是通过通信来共享内存，前者就是传统的加锁，后者就是Channel。
也就是说，设计Channel的主要目的就是在多任务间传递数据的，这当然是安全的。

channel操作分为四部分：创建、发送、接收和关闭。其中 发送、接收 会对 hchan 加锁。

### 2、go channel 的底层实现原理 （数据结构）

```
type hchan struct {
    qcount   uint           // 循环队列中数据数
    dataqsiz uint           // 循环队列的大小
    buf      unsafe.Pointer // 指向大小为dataqsize的包含数据元素的数组指针
    elemsize uint16         // 数据元素的大小
    closed   uint32         // 代表channel是否关闭   
    elemtype *_type         // _type代表Go的类型系统，elemtype代表channel中的元素类型
    sendx    uint           // 发送索引号，初始值为0
    recvx    uint           // 接收索引号，初始值为0
  recvq    waitq          // 接收等待队列，存储试图从channel接收数据(<-ch)的阻塞goroutines
    sendq    waitq          // 发送等待队列，存储试图发送数据(ch<-)到channel的阻塞goroutines

    lock mutex              // 加锁能保护hchan的所有字段，包括waitq中sudoq对象
}
```
参考：https://zhuanlan.zhihu.com/p/312041083

### 3、nil、关闭的 channel、有数据的 channel，再进行读、写、关闭会怎么样？（各类变种题型）

channel使用的注意事项：
* channel 中只能存放指定的数据类型。
* channle 的数据放满后，就不能再放入了。
* 如果从 channel 取出数据后，可以继续放入。
* 在没有使用协程的情况下，如果 channel 数据取完了，再取，就会报 dead lock。

参考：https://www.cnblogs.com/jiujuan/p/16014608.html

### 4、向 channel 发送数据和从 channel 读数据的流程是什么样的？

## map相关

### 1、map 使用注意的点，并发安全？

### 2、map 循环是有序的还是无序的？

无序的

### 3、 map 中删除一个 key，它的内存会释放么？

* 如果删除的元素是值类型，如int，float，bool，string以及数组和struct，map的内存不会自动释放
* 如果删除的元素是引用类型，如指针，slice，map，chan等，map的内存会自动释放，但释放的内存是子元素应用类型的内存占用

参考：https://blog.csdn.net/csdniter/article/details/103611783

### 4、怎么处理对 map 进行并发访问？有没有其他方案？ 区别是什么？

### 5、 nil map 和空 map 有何不同？

### 6、map 的数据结构是什么？是怎么实现扩容？

## GPM相关

### 1、什么是 GMP？（必问）

### 2、进程、线程、协程有什么区别？

### 3、抢占式调度是如何抢占的？

### 4、M 和 P 的数量问题？

## 锁相关

### 1、除了 mutex 以外还有那些方式安全读写共享变量？

通过 Channel 可以进行安全读写共享变量。

### 2、Go 如何实现原子操作？

### 3、Mutex 是悲观锁还是乐观锁？悲观锁、乐观锁是什么？

### 4、Mutex 有几种模式？

### 5、goroutine 的自旋占用资源如何解决

## 并发相关

### 1、怎么控制并发数？

### 2、多个 goroutine 对同一个 map 写会 panic，异常是否可以用 defer 捕获？

### 3、如何优雅的实现一个 goroutine 池（百度、手写代码）

## GC相关

### 1、go gc 是怎么实现的？（必问）

### 2、go 是 gc 算法是怎么实现的？ （得物，出现频率低）

### 3、GC 中 stw 时机，各个阶段是如何解决的？ （百度）

### 4、GC 的触发时机？

## 内存相关

### 1、谈谈内存泄露，什么情况下内存会泄露？怎么定位排查内存泄漏问题？

### 2、知道 golang 的内存逃逸吗？什么情况下会发生内存逃逸？

### 3、请简述 Go 是如何分配内存的？

Channel 分配在栈上还是堆上？哪些对象分配在堆上，哪些对象分配在栈上？

### 4、介绍一下大对象小对象，为什么小对象多了会造成 gc 压力？