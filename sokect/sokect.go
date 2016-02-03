package main

import (
	"./tcp"
	"time"
)

func main(){
	runChan := make(chan []byte, 1)
	go tcp.RunService()

	time.Sleep(1 * time.Millisecond)

	go tcp.RunClient()
	<- runChan
}
