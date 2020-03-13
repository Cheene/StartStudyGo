package main

import (
	"fmt"
	"time"
)

//时间相关的包

func main() {
	now := time.Now()
	fmt.Println(time.Now())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp1)
	fmt.Println(timestamp2)
	//time.Unix()
	ret := time.Unix(1564803667, 0)

	fmt.Println(ret)
	fmt.Println(ret.Date())
	fmt.Println(time.Second) //时间间隔
	//一个小时之后
	fmt.Println(now)
	fmt.Println(now.Add(time.Hour))
	////定时器
	//timer := time.Tick(time.Second)
	//for t := range timer{
	//	fmt.Println(t)
	//}
	//格式化 即把时间对象转换成字符串类型的时间
	fmt.Println(now.Format("2006-1-2"))
	fmt.Println(now.Format("15:04"))
	fmt.Println(now.Format("15:04:01"))
	fmt.Println(now.Format("2016/1/2"))
	fmt.Println(now.Format("2016/1/2 03:04:05 PM"))
	fmt.Println(now.Format("2016/1/2 03:04:05 PM"))
	//	按照 value 对应的格式，来解析对应的格式
	//需要先设置时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("Error%v", err)
	}

	timeObj, err := time.ParseInLocation("2006/1/2", "2020/3/12", loc)
	if err != nil {
		fmt.Println("parse time failed:%v", err)
	}
	td := now.Sub(timeObj)
	fmt.Println(td)
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	//Sub 求两个时间之间的差值,相减时时间相减
	fmt.Println(now.Sub(timeObj))
}
