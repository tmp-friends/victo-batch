package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Vtuber struct {
	Name            string `json:"name"`
	BelongsTo       string `json:"belongs_to"`
	TwitterUserName string `json:"twitter_user_name"`
	Channel         string `json:"channel"`
	HashtagName     string `json:"hashtag_name"`
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

		hashtagName := findHashtagName(name, doc)

		// 構造体に追加
		vtuber := Vtuber{
			Name:            name,
			BelongsTo:       "にじさんじ",
			TwitterUserName: hrefList[0],
			Channel:         hrefList[1],
			HashtagName:     hashtagName,
		}
		vtubers = append(vtubers, vtuber)

	}

	writeJsonFile(vtubers)
}

func findHref(doc *goquery.Document) []string {
	twitterUName := ""
	channel := ""
	doc.Find("#content .h-scrollable table .ext").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		href, _ := s.Attr("href")

		// TwitterUserName
		// 最初に現れるTwitterのリンクがtwitter_user_name
		if twitterUName == "" {
			if strings.Contains(href, "twitter.com") {
				twitterUName = href
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

	return []string{twitterUName, channel}
}

func findHashtagName(vtuberName string, doc *goquery.Document) string {
	keyword := "ファンアート：#"

	hashtagName := ""
	doc.Find("#content .h-scrollable table").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		idx := strings.Index(text, keyword) + len(keyword)
		if idx > len(keyword) {
			hashtagName = text[idx : idx+80]
			fmt.Printf("%s\n%s\n\n", vtuberName, hashtagName)
		}
	})

	return hashtagName
}

func writeJsonFile(vtubers []Vtuber) {
	file, _ := json.MarshalIndent(vtubers, "", "  ")
	_ = os.WriteFile("/var/batch/scripts/scraping/outputs/vtubers.json", file, 0755)
}
