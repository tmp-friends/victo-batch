package logic

import (
	"log"
	"time"

	"github.com/tmp-friends/victo-batch/functions/pkg/insert_fanart_tweets/service"
)

// Goのtime->string変換の仕様
var LAYOUT = "2006-01-02_15:04:05_JST"

type FanartTweetsLogic struct {
	service *service.FanartTweetsService
}

func NewFanartTweetsLogic() *FanartTweetsLogic {
	fts := service.NewFanartTweetsService()

	return &FanartTweetsLogic{
		service: fts,
	}
}

func (ftl *FanartTweetsLogic) DoExecute() {
	hashtags := ftl.service.GetHashtags()

	startTime, endTime := ftl.setWithinTime()
	log.Printf("Target within time: %s ~ %s\n", startTime, endTime)

	for _, v := range hashtags {
		log.Print(v.Name)

		tweets := ftl.service.FetchTweets(v.Name, startTime, endTime)
		if len(tweets) == 0 {
			log.Print("- Result count: 0\n")
			continue
		}

		authors := ftl.service.GetAuthors(tweets)

		log.Printf(
			"- Result count: %d\n- Oldest tweet id: %s\n- Newest tweet id: %s",
			len(tweets),
			// 新しい順に取得していく
			tweets[len(tweets)-1].ID,
			tweets[0].ID,
		)

		ftl.service.Insert(v.ID, tweets)
		ftl.service.InsertAuthors(authors)
	}
}

// (日本標準時間で)Twitter検索を実施する開始と終了の時刻を取得
func (ftl *FanartTweetsLogic) setWithinTime() (string, string) {
	// @see: https://tutuz-tech.hatenablog.com/entry/2021/01/30/192956
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	// 現在時刻の取得
	nowJST := time.Now().In(jst)

	startTime := time.Date(
		nowJST.Year(),
		nowJST.Month(),
		nowJST.Day()-1,
		4,
		0,
		0,
		0,
		jst,
	).Format(LAYOUT)
	endTime := time.Date(
		nowJST.Year(),
		nowJST.Month(),
		nowJST.Day(),
		4,
		0,
		0,
		0,
		jst,
	).Format(LAYOUT)

	return startTime, endTime
}
