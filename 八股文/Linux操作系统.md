### 1、大端与小端的概念？各自的优势是什么？

<font color=red>【答】大端与小端是用来描述多字节数据在内存中的存放顺序，即字节序。大端（Big Endian）是指低地址端存放高位字节，小端（Little Endian）是指低地址端存放低位字节。</font>

* <font color=red>Big Endian：符号位的判定固定为第一个字节，容易判断正负。</font>
* <font color=red>Little Endian：长度为1，2，4字节的数，排列方式都是一样的，数据类型转换非常方便。</font>

### 2、select / poll / epoll 的区别

|   | <font color=red>select</font> | <font color=red>poll</font>    | <font color=red>epoll</font>  |
|-------:|:---:|:---------:|:-------:|
| <font color=red>操作方式</font>  | 遍历 | 遍历     | 回调 |
| <font color=red>底层实现</font> |  数组 | 链表      | 哈希表   |
| <font color=red>IO效率</font>  | 每次调用都进行线性遍历，时间复杂度为O(n)   | 每次调用都进行线性遍历，时间复杂度为O(n) | 事件通知方式，每当fd就绪，系统注册的回调函数就会被调用，将就绪fd放到rdllist里面。时间复杂度O(1)  |
| <font color=red>最大连接数</font> |    1024（x86）或 2048（x64）  |       无上限   |    无上限    |
| <font color=red>fd拷贝</font> |  每次调用select，都需要把fd集合从用户态拷贝到内核态   |   每次调用poll，都需要把fd集合从用户态拷贝到内核态      |  调用epoll_ctl时拷贝进内核并保存，之后每次epoll_wait不拷贝  |

### 3、系统调用与库函数的区别？

* <font color=red><b>系统调用：</b>操作系统为用户程序与硬件设备进行交互提供的一组接口，发生在内核地址空间。</font>
* <font color=red><b>库函数：</b>把一些常用的函数编写完放到一个文件里，编写应用程序时调用，这是由第三方提供的，发生在用户地址空间。</font>
* <font color=red>在移植性方面，不同操作系统的系统调用一般是不同的，移植性差；而在所有的ANSI C编译器版本中，C库函数是相同的。</font>
* <font color=red>在调用开销方面，系统调用需要在用户空间和内核环境间切换，开销较大；而库函数调用属于“过程调用”，开销较小。</font>

<font color=red>简单点说，库函数就是系统调用的上层封装，为了让应用程序在使用时更加方便。</font>

### 4、守护、僵尸、孤儿进程的概念

* <font color=red><b>守护进程：</b>运行在后台的一种特殊进程，独立于控制终端并周期性地执行某些任务。</font>
* <font color=red><b>僵尸进程：</b>一个进程 fork 子进程，子进程退出，而父进程没有 `wait`/`waitpid`子进程，那么子进程的进程描述符仍保存在系统中，这样的进程称为僵尸进程。</font>
* <font color=red><b>孤儿进程：</b>一个父进程退出，而它的一个或多个子进程还在运行，这些子进程称为孤儿进程。（孤儿进程将由 init 进程收养并对它们完成状态收集工作）</font>

### 5、linux下你常用的命令有哪些？

<font color=red>【答】ll、pwd、touch、rm、mkdir、rmdir、mv、cp、ln
cat、less、more、tail、vim、vimdiff、grep
tar、rz、sz
df、du、free、top、ethtool、sar、netstat、iostat、ps
ifconfig、ping、talnet……</font>

### 6、死锁产生的四个条件，死锁发生后怎么检测和恢复？

<font color=red>【答】死锁的四个必要条件：</font>

* <font color=red>互斥条件：一个资源每次只能被一个进程使用。</font>
* <font color=red>请求与保持条件：一个进程在申请新的资源的同时保持对原有资源的占有。</font>
* <font color=red>不剥夺条件:进程已获得的资源，在未使用完之前，不能强行剥夺。</font>
* <font color=red>循环等待条件:若干进程之间形成一种头尾相接的循环等待资源关系。</font>

<font color=red>死锁发生后：</font>

* <font color=red>检测死锁：首先为每个进程和每个资源指定一个唯一的号码，然后建立资源分配表和进程等待表。</font>
* <font color=red>解除死锁：当发现有进程死锁后，可以直接撤消死锁的进程或撤消代价最小的进程，直至有足够的资源可用，死锁状态消除为止。</font>

### 7、阻塞IO、非阻塞IO、同步IO、异步IO的区别？

<font color=red>【解】这里讨论的是Linux环境下的network IO。
在Richard Stevens的《UNIX® Network Programming Volume 1, Third Edition: The Sockets Networking》的第6.2节介绍了五种IO Model，并说明了各种IO的特点和区别。</font>

* <font color=red>blocking IO</font>
* <font color=red>non-blocking IO</font>
* <font color=red>IO multiplexing</font>
* <font color=red>signal driven IO（不讨论）</font>
* <font color=red>asynchronous IO</font>

<font color=red>当一个网络IO的 read 操作发生时，它会经历两个阶段：</font>

1. <font color=red>等待数据准备 (Waiting for the data to be ready)</font>
2. <font color=red>将数据从内核拷贝到进程中 (Copying the data from the kernel to the process)</font>

<font color=red>这些IO Model的区别就是在两个阶段上各有不同的情况:</font>

* <font color=red><b>阻塞IO（blocking IO）：</b>线程阻塞以等待数据，然后将数据从内核拷贝到进程，返回结果之后才解除阻塞状态。也就是说两个阶段都被block了。</font>


* <font color=red><b>非阻塞IO（non-blocking IO）：</b>当对一个非阻塞socket执行读操作时，如果kernel中的数据还没有准备好，那么它并不会block用户进程，而是立刻返回一个error。用户进程需要不断地主动进行read操作，一旦数据准备好了，就会把数据拷贝到用户内存。也就是说，第一阶段并不会阻塞线程，但第二阶段拷贝数据还是会阻塞线程。</font>


* <font color=red><b>IO复用（IO multiplexing）：</b>这种IO方式也称为event driven IO. 通过使用select/poll/epoll在单个进程中同时处理多个网络连接的IO。例如，当用户进程调用了select，那么整个进程会被block，通过不断地轮询所负责的所有socket，当某个socket的数据准备好了，select就会返回。这个时候用户进程再调用read操作，将数据从kernel拷贝到用户进程。在IO复用模型中，实际上对于每一个socket，一般都设置成为non-blocking，但是，整个用户进程其实是一直被block的，先是被select函数block，再是被socket IO第二阶段block。</font>


* <font color=red><b>同步IO（synchronous IO）：</b>POSIX中的同步IO定义是—— A synchronous I/O operation causes the requesting process to be blocked until that I/O operation completes。也就是说同步IO在IO操作完成之前会阻塞线程，按照这个定义，之前所述的blocking IO，non-blocking IO，IO multiplexing都属于synchronous IO。（non-blocking IO也属于同步IO是因为它在真正拷贝数据时也会阻塞线程）</font>


* <font color=red><b>异步IO（asynchronous IO）：</b>POSIX中的异步IO定义是—— An asynchronous I/O operation does not cause the requesting process to be blocked。在linux异步IO中，用户进程发起read操作之后，直接返回，去做其它的事。而另一方面，从kernel的角度，当它受到一个asynchronous read之后，kernel会等待数据准备完成，然后将数据拷贝到用户内存，当这一切都完成之后，kernel会给用户进程发送一个signal，告诉它read操作完成了。也就是说两个阶段都不会阻塞线程。它就像是用户进程将整个IO操作交给了他人（kernel）完成，然后他人做完后发信号通知。在此期间，用户进程不需要去检查IO操作的状态，也不需要主动的去拷贝数据。</font>

### 8、进程与线程的区别？

对操作系统来说，线程是最小的执行单元，进程是最小的资源管理单元：

* 进程，在一定的环境下，把静态的程序代码运行起来，通过使用不同的资源，来完成一定的任务。
* 而线程是进程的一部分，是cpu调度的最小单位，线程主抓cpu执行代码的过程，其余的资源的保护和管理由整个进程去完成。
* 线程自己不拥有系统资源，只拥有一点在运行中必不可少的资源(运行栈和程序计数器pc)，切换开销小
* 一个进程中的多个线程可以共享进程所拥有的全部资源。

### 9、线程与协程的区别？

1. 线程从属于进程，是程序的实际执行者。一个进程可以有多个线程。线程拥有自己的栈空间。
2. 协程(Coroutines)，是一种比线程更加轻量级的存在，一个线程可以有多个协程。
3. 协程在用户态，不受内核调度，由用户程序管理；而线程是由内核调度的。

### 10、进程间的通信方式有哪些？

进程间通信的方式有：

1. 管道（包括无名管道pipe和命名管道FIFO） 
2. 消息队列
3. 共享内存 
4. Socket
5. Streams等 
   
其中 Socket和Streams支持不同主机上的两个进程IPC（InterProcess Communication）。



