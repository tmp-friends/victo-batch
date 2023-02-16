package logic

import (
	"log"
	"time"

	"github.com/tmp-friends/victo-batch/functions/insert_fanart_tweets/service"
)

var layout = "2022-01-01 00:00:00"

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
	log.Printf(
		"Target within time: %s ~ %s\n",
		startTime.Format(time.RFC3339),
		endTime.Format(time.RFC3339),
	)

	for _, v := range hashtags {
		log.Print(v.Name)

		tweets, media, resultCount := ftl.service.FetchTweets(v.Name, startTime, endTime)

		if resultCount == 0 {
			continue
		}

		log.Printf(
			"- Result count: %d\n- Oldest tweet id: %s\n- Newest tweet id: %s",
			resultCount,
			// 新しい順に取得していく
			tweets[len(tweets)-1].ID,
			tweets[0].ID,
		)

		ftl.service.InsertTweets(v.ID, tweets)
		ftl.service.InsertMedia(media)
	}
}

// (日本標準時間で)Twitter検索を実施する開始と終了の時刻を取得
func (ftl *FanartTweetsLogic) setWithinTime() (time.Time, time.Time) {
	// @see: https://tutuz-tech.hatenablog.com/entry/2021/01/30/192956
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	// 現在時刻の取得
	nowJST := time.Now().In(jst)

	yesterdayMidnight := time.Date(
		nowJST.Year(),
		nowJST.Month(),
		nowJST.Day()-1,
		0,
		0,
		0,
		0,
		jst,
	)
	todayMidnight := time.Date(
		nowJST.Year(),
		nowJST.Month(),
		nowJST.Day(),
		0,
		0,
		0,
		0,
		jst,
	)

	return yesterdayMidnight, todayMidnight
}
