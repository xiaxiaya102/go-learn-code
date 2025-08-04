package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

// 定义平台代码和平台名称为常量
const (
	Baidu         = "baidu"
	Shaoshupai    = "shaoshupai"
	Weibo         = "weibo"
	Zhihu         = "zhihu"
	Pojie         = "52pojie"
	Bilibili      = "bilibili"
	Douban        = "douban"
	Hupu          = "hupu"
	Tieba         = "tieba"
	Juejin        = "juejin"
	Douyin        = "douyin"
	V2ex          = "v2ex"
	JRTT          = "jinritoutiao"
	StackOverflow = "stackoverflow"
	GitHub        = "github"
	HackerNews    = "hackernews"
)

// 定义热点新闻url地址
const newsUrl = "https://orz.ai/api/v1/dailynews/?platform="

// 飞书消息发送webHook地址
const feishuUrl = "https://open.feishu.cn/open-apis/bot/v2/hook/8836666e-b0d3-43bf-beb8-82b774958d48"

// 定义新闻结构体
type News struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Content     string `json:"content"`
	Source      string `json:"source"`
	PublishTime string `json:"publish_time"`
}

// 定义响应结构体
type Response struct {
	Status string `json:"status"`
	Data   []News `json:"data"`
}

// getLastLyNewsByPlatformCode 根据平台编码获取平台热点新闻
func getLastLyNewsByPlatformCode(code string) []News {
	client := resty.New()

	resp, err := client.R().Get(newsUrl + code)
	if err != nil {
		log.Fatalln("请求失败")
	}
	//使用 map 来解析 JSON 数据
	var response Response
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		log.Fatalln("Error unmarshaling JSON:", err)
	}

	status := response.Status
	if status == "200" {
		news := response.Data
		log.Printf("热点新闻获取成功,供获取到:%d条", len(news))
		return news
	}
	return make([]News, 0)
}

// formatNewsMessage 格式化单条新闻消息
func formatNewsMessage(index int, news News) string {
	return fmt.Sprintf(
		"序号: %d\n"+
			"标题: %s\n"+
			"来源: %s\n"+
			"发布时间: %s\n"+
			"详情: (%s)\n",
		index, news.Title, news.Source, news.PublishTime, news.URL,
	)
}

// formatMultipleNewsMessages 格式化多条新闻消息
func formatMultipleNewsMessages(newsList []News) string {
	var message string
	for i, news := range newsList {
		message += formatNewsMessage(i, news) + "\n"
	}
	return message
}

// sendFeishuMessage 发送消息到飞书群组
func sendFeishuMessage(message string) {
	client := resty.New()

	// 构造发送的 JSON 数据
	payload := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": "每日热点新闻\n\n" + message,
		},
	}

	result, err := client.R().SetBody(payload).Post(feishuUrl)
	if err != nil {
		return
	}
	log.Printf("消息发送结果:%s", result)
}

// 第一个 init 函数
func init() {
	log.Println("每日新闻小助手启动成功")
}

func main() {

	// 创建一个定时器，每隔 1 分钟触发一次
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("执行定时任务:", time.Now())
			//获取热点新闻
			newList := getLastLyNewsByPlatformCode(Baidu)
			//处理热点新闻数据
			message := formatMultipleNewsMessages(newList)
			//发送热点新闻数据到飞书群组
			sendFeishuMessage(message)
		}
	}

}
