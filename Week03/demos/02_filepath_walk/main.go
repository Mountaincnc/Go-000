package main

import (
	"fmt"
	"os"
	"path/filepath"
)
// 忽略Music下的"网易云音乐"目录

func main() {
	err := filepath.Walk("/Users/songzeshan/Music", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("handling path %s failed: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == "网易云音乐" {
			fmt.Printf("skip 网易云音乐 directory")
			return filepath.SkipDir
		}
		fmt.Printf("get dir: %v\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
