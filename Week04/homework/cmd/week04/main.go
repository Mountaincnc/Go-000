package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	v1 "week04.homework.szs/api/article/v1"
	"week04.homework.szs/internal/pkg/server"
)


func main() {
	g, ctx := errgroup.WithContext(context.Background())

	// 初始化
	artServ := InitializeArticle()

	// 注册server
	s := server.NewServer(":9090")
	v1.RegisterArticleServer(s.Server, artServ)

	// 监听信号
	g.Go(func() error {
		sigChan := make(chan os.Signal)

		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		for {
			select {
			case sig := <-sigChan:
				return fmt.Errorf("get signal: %v\n", sig)
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	// 启动server
	g.Go(func() error {
		return s.Run(ctx)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("process exited %v\n", err.Error())
	}
}

