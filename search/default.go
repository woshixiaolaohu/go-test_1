package search

// defaultMatcher 实现了默认匹配器
type defaultMatcher struct {
}

// init 讲默认匹配器注册到程序
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现默认匹配器的行为 此处声明了一个方法  m为接收者类型
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
