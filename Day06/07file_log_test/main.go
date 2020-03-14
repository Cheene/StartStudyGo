package main

import (
	"../mylogger"
	"fmt"
	"time"
)

func main() {
	log := mylogger.NewFileLogger("DEBUG", "./", "chenen.log", 10*1024*1024)
	i := 1
	for {
		log.Error("This is Error %d", i)
		log.Info("This is Error %d", i)
		log.Debug("This is Error %d", i)
		log.Trace("This is Error %d", i)
		log.Fatal("This is Error %d", i)
		i++
		time.Sleep(2 * time.Second)
		fmt.Println(i)

	}

}
