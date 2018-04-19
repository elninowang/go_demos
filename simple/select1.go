package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	ch := make (chan int)
	select {
	case <-ch:
	case x := <-timeout:
		fmt.Println("timeout is ", x)
	}
}