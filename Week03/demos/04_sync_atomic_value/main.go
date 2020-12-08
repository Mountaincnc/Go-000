package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct {
	a []int
}

// atomic value 用于读多写少的场景

// 无锁无原子操作的错误例子
/*
func main() {
	cfg := &Config{}

	go func() {
		i := 0
		for {
			i++
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
*/

// 用了RWLock的例子
/*
func main() {
	cfg := &Config{}

	var rw sync.RWMutex
	go func() {
		i := 0
		for {
			i++
			rw.Lock()
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			rw.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				rw.RLock()
				fmt.Printf("%v\n", cfg)
				rw.RUnlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

 */

// 用了atomic value的例子
func main() {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			cfg := &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load().(*Config)
				fmt.Printf("cfg: %v\n", cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}


