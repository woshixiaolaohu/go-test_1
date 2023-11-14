package search

// Result 保存搜索结果
type Result struct {
	Field   string
	Content string
}

// Matcher 定义要实现的新搜索类型行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
