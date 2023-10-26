package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Code int64 `json:"code"`
	Data struct {
		Replies []struct {
			Content struct {
				Emote struct {
				} `json:"jump_url"`
				MaxLine int64         `json:"max_line"`
				Members []interface{} `json:"members"`
				Message string        `json:"message"`
			} `json:"content"`
			Count  int64 `json:"count"`
			Folder struct {
				HasFolded bool   `json:"has_folded"`
				IsFolded  bool   `json:"is_folded"`
				Rule      string `json:"rule"`
			} `json:"folder"`
			Invisible bool  `json:"invisible"`
			Like      int64 `json:"like"`
			Replies   []struct {
				Action  int64 `json:"action"`
				Assist  int64 `json:"assist"`
				Attr    int64 `json:"attr"`
				Content struct {
					JumpURL struct{} `json:"jump_url"`
					MaxLine int64    `json:"max_line"`
					Message string   `json:"message"`
				} `json:"content"`
				Count        int64  `json:"count"`
				Ctime        int64  `json:"ctime"`
				Dialog       int64  `json:"dialog"`
				DynamicIDStr string `json:"dynamic_id_str"`
				Fansgrade    int64  `json:"fansgrade"`
				Folder       struct {
					HasFolded bool   `json:"has_folded"`
					IsFolded  bool   `json:"is_folded"`
					Rule      string `json:"rule"`
				} `json:"folder"`
				Invisible bool        `json:"invisible"`
				Like      int64       `json:"like"`
				Mid       int64       `json:"mid"`
				Oid       int64       `json:"oid"`
				Parent    int64       `json:"parent"`
				ParentStr string      `json:"parent_str"`
				Rcount    int64       `json:"rcount"`
				Replies   interface{} `json:"replies"`
				UpAction  struct {
					Like  bool `json:"like"`
					Reply bool `json:"reply"`
				} `json:"up_action"`
			} `json:"replies"`
		} `json:"replies"`
		Message string `json:"message"`
	}
}

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/v2/reply/main?next=0&type=1&oid=721375702", nil)
	if err != nil {
		fmt.Println("err", err)
	}
	req.Header.Set("authority", "www.bilibili.com")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.69")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Referer", "https://search.bilibili.com/all?keyword=%E5%A5%A2%E9%A6%99%E5%A4%AB%E4%BA%BA&from_source=webtop_search&spm_id_from=333.1007&search_source=5")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io err", err)
	}
	var resultList Result
	_ = json.Unmarshal(bodyText, &resultList)
	for _, result := range resultList.Data.Replies {
		fmt.Println("一级评论", result.Content.Message)
		for _, reply := range result.Replies {
			fmt.Println("二级评论", reply.Content.Message)
		}
	}
}
