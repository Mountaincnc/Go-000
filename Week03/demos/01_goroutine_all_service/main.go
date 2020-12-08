package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 定义两个channel
	// 一个用于接收启动时的error
	// 另一个用于发送停止信号
	errorCh := make(chan error, 2)
	stopped := make(chan struct{})

	// 定义handler 启动server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp,"hello world")
	})
	go func() {
		// 如果server启动失败, 将错误信息发送到errorCh
		errorCh <- startServer(":8080", mux, stopped)
	}()

	// 启动debug
	go func() {
		errorCh <- startDebug("127.0.0.1:8081", stopped)
	}()

	// 定义flag 避免close两次 产生panic (用sync.Once应该也行?)
	var flag bool
	// 轮询 errorCh 获取error
	for err := range errorCh {
		// 判定error
		if err != nil  {
			fmt.Printf("error: %v\n", err)

			if !flag {
				flag = true
				close(stopped)
			}
		}
	}
}

func startServer(addr string, handler http.Handler, stopped chan struct{}) error {
	s := http.Server{
		Addr:              addr,
		Handler:           handler,
	}

	go func() {
		// 监听stopped channel 如果被关闭 则执行shutdown
		<- stopped
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func startDebug(pprofAdddr string, stopped chan struct{}) error {
	debug := http.Server{
		Addr:              pprofAdddr,
		Handler:           http.DefaultServeMux,
	}
	go func() {
		// 监听stopped channel 如果被关闭 则执行shutdown
		<- stopped
		debug.Shutdown(context.Background() )
	}()
	return debug.ListenAndServe()
}