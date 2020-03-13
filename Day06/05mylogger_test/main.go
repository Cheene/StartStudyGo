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
		log.Error("This is info")

		time.Sleep(time.Second * 3)
	}
}
