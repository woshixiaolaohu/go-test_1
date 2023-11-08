package main

import (
	"log"
	"os"
	_ "test_1/matchers"
	"test_1/search"
)

// init在main之前调用
func init() {
	// 日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main程序入口
func main() {
	search.Run("president")
}
