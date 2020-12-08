package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled")
		}
	}(ctx)

}
