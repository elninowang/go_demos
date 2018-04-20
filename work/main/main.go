package main

import (
	"log"
	"sync"
	"time"
	".."
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println("+", m.name)
	time.Sleep(time.Second)
	log.Println("-", m.name)
}

func main() {
	p := work.New(5)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i ++ {
		for _, name := range names {
			np := namePrinter {
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}