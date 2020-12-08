package main

import (
	"context"
	"fmt"
	"time"
)

type Tracker struct {
	eventCh chan string
	stop    chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		eventCh: make(chan string, 10),
		stop:    make(chan struct{}),
	}
}

func (t *Tracker) Event(ctx context.Context, event string) error {
	select {
	case t.eventCh <- event:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for event := range t.eventCh {
		fmt.Println(event)
	}
	// ③发送数据到stop
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	// ①关闭eventCh
	// ②然后Run方法中将eventCh内缓存的数据读取完毕之后 在Run方法中发送数据到stop
	close(t.eventCh)

	select {
	// ④捕获到stop被关闭的动作, 然后做一些操作
	case <-t.stop:
		fmt.Printf("shutdown")
	// context做超时控制
	case <-ctx.Done():
		fmt.Printf("context canceled")
	}
}

func main() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
	// cancel 必须要调用一下 否则可能出现goroutine泄露
	defer cancel()

	tr.Shutdown(ctx)
}
