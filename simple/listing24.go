package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGorountines = 4
	taskLoad = 10
)

var wg sync.WaitGroup

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGorountines)

	for gr := 1; gr <= numberGorountines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task(%d)", post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		fmt.Printf("Worker: %d : Complated %s\n", worker, task)
	}
}