package main

import (
	"context"
	"fmt"
)

const ENV = "env"

func main() {
	ctx := context.WithValue(context.Background(), ENV, "test")
	fmt.Println(ctxValue(ctx))
}

func ctxValue(ctx context.Context) interface{} {
	return ctx.Value(ENV)
}