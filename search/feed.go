package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed 包含需要处理的数据源信息
type Feed struct {
	// ``-标记 描述json解码的源数据
	Name string `json:"site"`
	URL  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// 当函数返回时 关闭文件
	defer file.Close()
	// 将文件解码到一个切片里 这个切片的每一项是一个指向Feed类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 这个函数不需要检查错误 谁掉用谁检查
	return feeds, err
}
