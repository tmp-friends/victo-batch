package logic

import (
	"time"

	"github.com/tmp-friends/victo-batch/functions/insert_fanart_tweets/service"
)

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

	tm := ftl.setWithinTime()
	for _, v := range hashtags {
		tweetsRes := ftl.service.FetchTweets(v.Name, tm)

		if tweetsRes.Tweets == nil {
			continue
		}

		tweets := tweetsRes.Tweets
		media := tweetsRes.Includes.Media

		ftl.service.InsertTweets(v.ID, tweets)
		ftl.service.InsertMedia(media)
	}

}

// (日本標準時間で)Twitter検索を実施する開始と終了の時刻を取得
func (ftl *FanartTweetsLogic) setWithinTime() map[string]time.Time {
	// @see: https://tutuz-tech.hatenablog.com/entry/2021/01/30/192956
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	// 現在時刻の取得
	nowJST := time.Now().In(jst)

	tm := map[string]time.Time{}
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

	tm["startTime"] = yesterdayMidnight
	tm["endTime"] = todayMidnight

	return tm
}
