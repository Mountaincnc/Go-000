package main

import (
	"fmt"
	xerrors "github.com/pkg/errors"
	"os"
	"week02/internal/week02/biz"
)

func main() {
	user, err := biz.User(45424)
	if err != nil {
		fmt.Sprintf("original error: %T %v\n", xerrors.Cause(err), xerrors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		os.Exit(1)
	}
	fmt.Println(user)
}