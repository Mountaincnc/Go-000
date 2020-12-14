package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello World")
	})

	g, ctx := errgroup.WithContext(context.Background())
	// 启动http server
	g.Go(func() error {
		return startServer(ctx, ":8080", mux)
	})

	// 监听signal
	g.Go(func() error {
		return sysSig(ctx)
	})

	// 等待错误
	if err := g.Wait(); err != nil {
		fmt.Printf("service exited %s\n", err.Error())
	}

}

func startServer(ctx context.Context, addr string, handler http.Handler) error {
	s := &http.Server{
		Addr:              addr,
		Handler:           handler,
	}

	// 监听context
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered: %v\n", r)
			}
		}()
		<- ctx.Done()
		fmt.Println("get system signal, shutdown http server")

		// 优雅关闭http server
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func sysSig(ctx context.Context) error {
	// 创建信号chan
	sigCh := make(chan os.Signal)

	// 监听signal
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	for {
		select {
		case <- ctx.Done():
			fmt.Println("http server exited")
			return ctx.Err()
		case s := <-sigCh:
			return fmt.Errorf("get signal: %v\n", s)
		}
	}
}