

### 1、简单说说协程底层原理 GPM 模型

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/go-GPM1.png?raw=true)

![](https://github.com/SongLee24/LeetCode-Go/blob/main/%E5%85%AB%E8%82%A1%E6%96%87/images/go-GPM2.jpeg?raw=true)

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




