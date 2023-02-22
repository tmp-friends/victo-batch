package service

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/sivchari/gotwtr"
	"github.com/tmp-friends/victo-batch/functions/insert_fanart_tweets/dao"
	"github.com/tmp-friends/victo-batch/functions/lib"
	"github.com/tmp-friends/victo-batch/functions/models"
)

type FanartTweetsService struct {
	lib *lib.TwitterClient
	dao *dao.FanartTweetsDao
}

func NewFanartTweetsService() *FanartTweetsService {
	tc := lib.NewTwitterClient()
	ftd := dao.NewFanartTweetsDao()

	return &FanartTweetsService{
		lib: tc,
		dao: ftd,
	}
}

func (fts *FanartTweetsService) GetHashtags() models.HashtagSlice {
	return fts.dao.GetHashtags()
}

func (fts *FanartTweetsService) FetchTweets(
	hashtagName string,
	startTime time.Time,
	endTime time.Time,
) ([]*gotwtr.Tweet, []*gotwtr.Media, int) {
	keyword := fmt.Sprintf("#%s -is:retweet has:images", hashtagName)
	tweets, media, resultCount := fts.lib.SearchRecent(keyword, startTime, endTime)

	return tweets, media, resultCount
}

func (fts *FanartTweetsService) Insert(hashtagId int, tweets []*gotwtr.Tweet, media []*gotwtr.Media) {
	for _, t := range tweets {
		// HACK: Tweetに付随したMediaKeysを取得する際、取得できずにヌルポになる事象がある
		if t.Attachments == nil {
			continue
		}

		var url string
		url, t.Text = fts.extractTweetUrl(t.Text)

		fts.dao.InsertTweetObject(t, url, hashtagId)

		// 1つのツイートに対してメディアが配列でついているためforで回す
		for _, mediaKey := range t.Attachments.MediaKeys {
			for i, m := range media {
				if m.MediaKey != mediaKey {
					continue
				}
				fts.dao.InsertMediaObject(media[i], t.ID)
			}
		}
	}
}

// ツイート本文に含まれているtweet_urlを抽出する処理
func (fts *FanartTweetsService) extractTweetUrl(text string) (string, string) {
	// https://t.co/<空白以外の一文字以上>
	re := regexp.MustCompile(`https://t.co/\S*$`)
	url := re.FindString(text)

	extractedText := strings.Replace(text, url, "", -1)

	return url, extractedText
}
