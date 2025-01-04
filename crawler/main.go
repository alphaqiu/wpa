package main

import (
	"fmt"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/gocolly/colly/v2"
)

const (
	link = "https://mp.weixin.qq.com/s/C_1c6CoSeH8bEZ2qVCO4fg"
	// link = "https://mp.weixin.qq.com/s/DkpTbkUy-ULBNMk0DIYscg"
)

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	getArticle(c)
	err := c.Visit(link)
	if err != nil {
		fmt.Println(err)
	}

}

func getArticle(c *colly.Collector) {
	// 获取公众号文章标题
	c.OnHTML("#img-content", func(e *colly.HTMLElement) {
		// 在 #img-content 内部查找 .rich_media_title
		titleElement := e.ChildText(".rich_media_title")
		// 打印标题内容
		fmt.Println("Title:", titleElement)
		fmt.Println("--------------------------------")
		// fmt.Println(e.DOM.Html())

	})

	// 获取公众号文章正文
	c.OnHTML(`#js_content`, func(e *colly.HTMLElement) {
		html, err := e.DOM.Html()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 将正文html转换为markdown
		markdown, err := htmltomarkdown.ConvertString(html)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(markdown)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}
