package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) Increment(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++

}

func main() {
	var wg sync.WaitGroup
	containers := &Container{
		counters: map[string]int{"sajan": 0, "Jaiswal": 0},
	}

	count := func(name string, n int) {

		for i := 1; i <= n; i++ {

			containers.Increment(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go count("sajan", 100)
	go count("sajan", 300)
	go count("Jaiswal", 300)
	wg.Wait()
	fmt.Println(containers.counters)

}
