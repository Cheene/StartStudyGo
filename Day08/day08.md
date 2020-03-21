## 内容回顾

### channel
* 为什么需要 channel？
通过 channel，实现多个 gorotine 之间的通信

`CSP`:通过通信来共享内存   

* channel 是一种引用类型，所以需要 `make`  函数来初始化
* 声明： `var ch chan type`
* 初始化： `ch = make(chan type,[可选的缓冲区的大小]`
* 三种操作
    * 发送 `ch <- 100`
    * 接收 `var := <- ch`
    * 关闭 `close(ch) ` 但不是必须的，会自动回收掉的
* 带缓冲区的与无缓冲区的通道
    * 快递柜 快递到快递柜里 --》 带缓冲区的通道
    * 无快递柜 快递直接送到家 --》 无缓冲区的通道
* channel 的单向通道：
    * 只读通道： `<- chan int`
    * 只写通道： `chan int <-`
 * select
    * 多路复用，同一时刻有多个通道执行的时候
    * 每一次的执行都是随机的
### 并发 goroutine
并发和并行的区别

如何启动 `goroutine` : go 关键字

* 启动： 将要并发执行的任务包装成一个函数，

* 结束：gorotine 对应的函数执行完，gorotine 就结束了

        注 程序启动的时候创建一个goroutine 去执行一个 main 函数
      
main 函数结束后，程序就结束了，由该程序启动的其他所有的goroutine 也都结束了。

`sync.WiatGroup` 判断 goroutine 是否全部结束完成
 `var wg sync.WaitGroup` 结构体
* wg.Add(number) --> 外部调用之外（计数器+1）
* wg.Done() --> --> go 内部函数使用（计数器-1）
* wg.Wait() --> 等待结束（等）

### goroutine 的本质

### gorouttine 的调度模型 `GMP`

`M:N` 把 M 个 goroutine 分配个 N 个 操作系统线程

### goroutine 与 操作系统线程（OS 线程）的区别
goroutine 是用户态的线程，比内核态的线程更轻量，初始的时只占用2KB，所以可以轻松开始十万个goroutine

OS线程是 内核态的线程，至少需要2MB

### runtime.GOMAXPROCS 

Go1.5之后，是操作系统的逻辑核心数

默认就是跑满CPU

```
runtime.GOMAXPROCS(1) // 仅占用1个核
```

### work pool 模式
自己写一个方法，开启一定数量的 goroutine 去干活


GMP

go


### 

## 今日内容



### sync包

### context

### 网络编程