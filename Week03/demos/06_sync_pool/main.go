package main

import (
	"fmt"
	"sync"
)

func main() {
	// 注: sync pool的缓存期限是两次gc之间间隔的时间 所以不能用于做连接池
	p := &sync.Pool{New: func() interface{} {
		return 0
	}}

	i := p.Get().(int)
	fmt.Println(i)
	p.Put(1)
	j := p.Get().(int)
	fmt.Println(j)
}