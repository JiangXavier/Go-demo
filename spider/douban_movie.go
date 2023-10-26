package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	Spider()
}

func Spider() {
	//1.发送请求
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/top250", nil)
	if err != nil {
		fmt.Println("req err", err)
	}
	///加请求头,防止浏览器检测爬虫访问，伪造浏览器访问
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://movie.douban.com/chart")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	///发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败", err)
	}
	//2.解析网页
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析失败", err)
	}
	//3.获取节点信息
	// #content > div > div.article > ol > li:nth-child(1) > div > div.info > div.hd > a > span:nth-child(1)
	// #content > div > div.article > ol > li:nth-child(1)
	docDetail.Find("#content > div > div.article > ol > li"). //列表
									Each(func(i int, s *goquery.Selection) { //列表下继续找
			title := s.Find("div > div.info > div.hd > a > span:nth-child(1)").Text()
			img := s.Find("div > div.pic > a > img")
			imgTmp, ok := img.Attr("src")
			info := s.Find("div > div.info > div.bd > p:nth-child(1)").Text()
			score := s.Find("div > div.info > div.bd > div > span.rating_num").Text()
			quote := s.Find("div > div.info > div.bd > p.quote").Text()
			if ok {
				fmt.Println("title", title)
				fmt.Println("imgTmp", imgTmp)
				fmt.Println("info", info)
				fmt.Println("score", score)
				fmt.Println("quote", quote)
			}
		})
	//4.保存信息
}
