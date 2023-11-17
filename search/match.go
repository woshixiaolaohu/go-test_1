package search

import (
	"fmt"
	"log"
)

// Result 保存搜索结果
type Result struct {
	Field   string
	Content string
}

// Matcher 定义要实现的新搜索类型行为 接口类型
// interface声明一个接口
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 为每个数据源启动一个goroutine来执行这个函数 并发执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 对特定的匹配器执行搜索
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	// 将结果写入通道
	for _, result := range searchResult {
		results <- result
	}
}

// Display 从每个单独的goroutine接收到结果之后 在终端窗口输出
func Display(results chan *Result) {
	// 通道会被一直阻塞 直到有结果写入 一旦通道关闭 for循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
