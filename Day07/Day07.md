###内容回顾
#### time
`2006-01-02 15:04:05 000`
#####时间类型
* time.Time: 
*时间戳：   
   * `time.Now().unix()`1970.1.1
   *  `time.Now().UnixNano()`        
#####时间间隔类型
* `time.Duration`: 毫秒，纳秒，微妙，秒，分钟，小时
##### 时间操作
时间对象 +/- 一个时间间隔的对象
time.
#####时间格式化
#####时间定时器
`time,Tick`

```go
timer := time.Tick(time.Second)
for t := range timer{
    fmt.Println(t)
}
````
##### 字符串转换为时间（时区）
`time.Parse("格式要求","需要解析的时间")`
 地区时间：
    `time.ParseInLocaton()`   
####日志库
time
文件操作
`runtime.Caller()`
###反射
变量底层是两个部分：动态类型和动态值

反射的应用： `json` / ORM等工具 / 

反射的两个方法：
   * reflect.TypeOf(); 
   * reflect.ValueOf()
###今日内容

并发
gorutine
`goroutine` 什么时候结束？

对应的函数结束的时候，goroutine就结束了

`main` 函数执行完了，由`main` 函数创建的那些 `goroutine` 也结束了

`goroutine` 调度模型


#### channel
并发模型是 CSP，提倡通过通信共享内存而不是共享内存实现通信

var 变量名 chan 类型
##### 通道的操作
`<-`
箭头表示数据流动的方向

1 发送 `ch <- 1`

2 接收 `ret := <- ch`

3 关闭 `close()`

单向通道，多用于在函数中


sync

##作业：
1 日志库中 channel 怎么使用

2 什么时候去起后台的goroutine去写日志到文件中