#day09
## 分享
* 注释
* 日志
* 单元测试
* 要做就做专业的，从一开始就规范起来。

# 内容回顾
## 互斥锁
`sync.Mutex`
是结构体，是值类型 ==> 传递函数参数的时候，使用指针
### 两个方法
```go
var lock sync.Mutex
lock.Lock()
lock.UnLock()
````
###为什么使用锁？
保证一次仅有一个访问临界变量；

防止同一时刻多个goroutine操作同一个资源

## 读写互斥锁
适用于读多写少的场景
1 获取读锁的时候，后续读的goroutine都可以，后续写的goroutine不能写
2 写的goroutine，后面无论是读还是写都要等待

```go
var rwLock sync.RWMutex

rwLock.RLock() //获取读锁
rwLock.RUnlock() // 获取写锁


```

## 等待组

`sync.Waitgroup`
结构体，值类型 --》 需要传递指针类型。用来等待goroutine执行完结束

```go
var wg sync.Waitgroup

wg.Add(number int) //goroutine 并发执行之前
wg.Done() // goroutine 函数结束之前， 相当于计时器 -1 
wg.Wait() // 等待全部 goroutine 执行完成

```

### Sync.Once
某些函数只需要执行一次
### Sync.Map
并发操作一个map的时候，内置的map不是并发安全的。
开箱即用的安全的map

```go
var syncMap sync.Map
synaMap.Store(key,value)
syncMap.Load(key)
syncMap.Delete()
sync.Range()
````

### 原子操作
Go 语言内置了一些针对内置的基本数据类型的一些并发安全操作

`atomic.xxxx`




# 今日内容
## context
## 单元测试 -- 开发人员的测试
## pprof 调试工具