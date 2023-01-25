package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Vtuber struct {
	Name          string `json:"name"`
	BelongsTo     string `json:"belongs_to"`
	TwitterUserId string `json:"twitter_user_id"`
	Channel       string `json:"channel"`
}

func main() {
	vtuberNames := getVtuberNames()
	getVtuberInfo(vtuberNames)
}

func getVtuberNames() []string {
	BEGIN_NOT_TARGET := "艾因"

	doc := requestHTML("https://wikiwiki.jp/nijisanji/%E5%85%AC%E5%BC%8F%E3%83%A9%E3%82%A4%E3%83%90%E3%83%BC")

	vtuberNames := []string{}
	// list1クラスがたくさんあるので、img属性を条件に追加することでVtuberに限定している
	doc.Find(".list1 img").EachWithBreak(func(i int, s *goquery.Selection) bool {
		imgTitle, _ := s.Attr("title")
		name := strings.TrimSpace(imgTitle)

		if name != "" {
			// 対象外のVtuberの名前に到達したらEach処理を終える
			if name == BEGIN_NOT_TARGET {
				return false
			}
			vtuberNames = append(vtuberNames, name)
		}

		return true
	})

	fmt.Println(len(vtuberNames))

	return vtuberNames
}

// 全自動化はできていないので、手動で値を調整する必要あり
func getVtuberInfo(vtuberNames []string) {
	vtubers := []Vtuber{}
	for _, name := range vtuberNames {
		// vtuberの名前とURLParamが違う
		if name == "ましろ爻" {
			name = "ましろ"
		}

		doc := requestHTML("https://wikiwiki.jp/nijisanji/" + name)
		hrefList := findHref(doc)

		// 構造体に追加
		vtuber := Vtuber{
			Name:          name,
			BelongsTo:     "にじさんじ",
			TwitterUserId: hrefList[0],
			Channel:       hrefList[1],
		}
		vtubers = append(vtubers, vtuber)

		displayHashtagName(name, doc)
	}

	writeJsonFile(vtubers)
}

func findHref(doc *goquery.Document) []string {
	twitterUId := ""
	channel := ""
	doc.Find("#content .h-scrollable table .ext").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		href, _ := s.Attr("href")

		// Twitter UserId
		// 最初に現れるTwitterのリンクがTwitterUId
		if twitterUId == "" {
			if strings.Contains(href, "twitter.com") {
				twitterUId = href
			}
		}

		// Youtubeチャンネル
		// 最初に現れるYoutubeのリンクがchannel
		if channel == "" {
			if strings.Contains(href, "www.youtube.com/channel") {
				channel = href
			}
		}

		return true
	})

	return []string{twitterUId, channel}
}

func displayHashtagName(vtuberName string, doc *goquery.Document) {
	keyword := "ファンアート：#"

	doc.Find("#content .h-scrollable table").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		idx := strings.Index(text, keyword) + len(keyword)
		if idx > len(keyword) {
			hashtagName := text[idx : idx+80]
			fmt.Printf("%s\n\"hashtag_name\": \"%s\n\n", vtuberName, hashtagName)
		}
	})
}

func writeJsonFile(vtubers []Vtuber) {
	file, _ := json.MarshalIndent(vtubers, "", "  ")
	_ = os.WriteFile("/var/batch/outputs/vtubers.json", file, 0755)
}
