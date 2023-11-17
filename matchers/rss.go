package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"test_1/search"
)

type (
	// item根据item字段标签 将定义的字段和rss文档字段关联起来
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"geoRssPoint"`
	}
	// image根据image字段标签 将定义的字段和rss文档字段关联起来
	image struct {
		XMLName xml.Name `xml:"img"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"Link"`
	}
	// channel根据channel字段标签将定义的字段和rss文档字段关联起来
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}
	// rssDocument 定义了与rss文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher 实现Matcher接口
type rssMatcher struct {
}

// init将匹配器注册到程序
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search 在文档中查找特定搜索项
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result
	log.Printf("Search Feed Type[%s] Site[%s] For Url[%s]\n", feed.Type, feed.Name, feed.URL)

	//获取要搜索的数据
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}
	for _, channelItem := range document.Channel.Item {
		// 检查标题是否包含搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}
		// 如果找到匹配项 将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}
		// 检查描述部分是否含有搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}
		// 如果找到匹配项 将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}

func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URL == "" {
		return nil, errors.New("no rss feed URL Provided")
	}
	// 从网络获取rss数据源文件
	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, err
	}

	// 一旦函数返回 关闭返回的相应链接
	defer resp.Body.Close()
	// 检查状态码是否200 就可以知道是否收到了正确的响应
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// 将rss数据源文件解码到我们定义的机构类型里 不需要检查错误 调用者会检查
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
