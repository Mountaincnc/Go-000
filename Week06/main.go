package main

import (
	"container/ring"
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

type slidingWindowLimit struct {
	// 单位时间内的最大请求数量
	maxRequest int32
	// 单位时间
	unitTime time.Duration
	// 窗口时间
	windowTime time.Duration
	// 窗口数量
	windowCount int
}

type window struct {
	// 窗口编号
	index int
	// 窗口时间内能处理的最大请求量
	maxRequest int32
	// 窗口时间内已处理请求
	handledRequest int32
}

// NewSwl 构建新的滑动窗口限制
func NewSwl(maxRequest int32, unitTime time.Duration, windowCount int) *slidingWindowLimit {
	if ! ( unitTime / time.Duration(windowCount) > 0 ) {
		return nil
	}
	return &slidingWindowLimit{
		maxRequest:  maxRequest,
		unitTime:    unitTime,
		windowTime:  unitTime / time.Duration(windowCount),
		windowCount: windowCount,
	}
}

// HandleConn 处理连接请求
func HandleConn(w *window ,conn *net.TCPConn) {
	defer (*conn).Close()

	n := atomic.AddInt32(&w.handledRequest, 1)

	if n > w.maxRequest {
		atomic.AddInt32(&w.handledRequest, -1)
		(*conn).Write([]byte("HTTP/1.1 429 TOO MANY REQUEST\r\n\r\nError, too many request, please try again."))
		return
	}

	(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World."))
	return

}

// UnitHandledCount 单位时间内处理的请求总数
func UnitHandledCount(r *ring.Ring) (sum int32) {
	for i := 0; i < r.Len(); i++ {
		sum += r.Value.(*window).handledRequest
		r = r.Next()
	}

	return
}

func main() {
	// 启动tcp server
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	swl := NewSwl(100, 10, 10)

	// 创建窗口(环形链表)
	r := ring.New(swl.windowCount)
	for i := 0; i < r.Len(); i++ {
		r.Value = &window{
			index:          i,
			maxRequest:     swl.maxRequest,
			handledRequest: 0,
		}
		r = r.Next()
	}
	// 获取窗口
	w := r.Value.(*window)
	// 配置定时器
	tracker := time.NewTicker(swl.windowTime * time.Second)
	defer tracker.Stop()

	go func() {
		for range tracker.C {
			// 每经过一个窗口时间, 就向后滑动
			r = r.Move(1)
			w = r.Value.(*window)
			// 计算当前窗口时间内还可以请求的数量
			w.maxRequest = swl.maxRequest - UnitHandledCount(r)
			fmt.Printf("window index: %v, max request: %v, handled request: %v\n", w.index, w.maxRequest,
				w.handledRequest)
		}
	}()

	for {
		// 获取连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		err = conn.SetKeepAlive(false)
		if err != nil {
			continue
		}

		// 处理连接
		go HandleConn(w, conn)
	}

}
