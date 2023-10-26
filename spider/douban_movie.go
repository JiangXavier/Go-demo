package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	USERNAME = "root"
	PASSWORD = ""
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "douban_movie"
)

//数据库初始化
var DB *sql.DB

type MovieData struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Picture  string `json:"picture"`
	Actor    string `json:"actor"`
	Year     string `json:"year"`
	Score    string `json:"score"`
	Quote    string `json:"quote"`
}

func main() {
	InitDB()
	for i := 0; i < 10; i++ {
		fmt.Printf("正在爬取第 %d 页信息\n", i)
		Spider(strconv.Itoa(i * 25))
	}
}

func Spider(page string) {
	//1.发送请求
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/top250?start="+page, nil)
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
	docDetail.Find("#content > div > div.article > ol > li"). //列表
									Each(func(i int, s *goquery.Selection) { //列表下继续找
			var data MovieData
			title := s.Find("div > div.info > div.hd > a > span:nth-child(1)").Text()
			img := s.Find("div > div.pic > a > img")
			imgTmp, ok := img.Attr("src")
			info := s.Find("div > div.info > div.bd > p:nth-child(1)").Text()
			score := s.Find("div > div.info > div.bd > div > span.rating_num").Text()
			quote := s.Find("div > div.info > div.bd > p.quote").Text()
			if ok {
				director, actor, year := InfoSpite(info)
				data.Title = title
				data.Director = imgTmp
				data.Picture = director
				data.Actor = actor
				data.Year = year
				data.Score = score
				data.Quote = strings.TrimSpace(quote)
				if InsertData(data) {
					//fmt.Println("插入成功")
				} else {
					fmt.Println("插入失败")
					return
				}
			}
		})
	fmt.Println("插入成功")
	return
	//4.保存信息
}

func InfoSpite(info string) (director, actor, year string) {
	directorRe, _ := regexp.Compile(`导演:(.*)主演:`)
	director = string(directorRe.Find([]byte(info)))

	actorRe, _ := regexp.Compile(`主演:(.*)`)
	actor = string(actorRe.Find([]byte(info)))

	yearRe, _ := regexp.Compile(`(\d+)`)
	year = string(yearRe.Find([]byte(info)))
	return
}

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(5)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
}

func InsertData(movieData MovieData) bool {
	//新建一个事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("begin err", err)
		return false
	}
	stmt, err := tx.Prepare("INSERT INTO movie_data (`Title`,`Director`,`Picture`,`Actor`,`Year`,`Score`,`Quote`) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("prepare fail err", err)
		return false
	}
	_, err = stmt.Exec(movieData.Title, movieData.Director, movieData.Picture, movieData.Actor, movieData.Year, movieData.Score, movieData.Quote)
	if err != nil {
		fmt.Println("exec fail", err)
		return false
	}
	_ = tx.Commit()
	return true
}
