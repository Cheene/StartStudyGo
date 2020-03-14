package main

import (
	"../mylogger"
	"time"
)

func main() {
	log := mylogger.NewLog("Debug")
	for {
		log.Debug("这是一条Debug日志")
		log.Trace("This is Trace")
		log.Info("This is info")
		id := 1000
		name := "chenene"
		log.Error("This is id:%d,name:%s", id, name)
		time.Sleep(time.Second * 3)
	}
}
