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

func (fts *FanartTweetsService) InsertTweets(hashtagId int, tweets []*gotwtr.Tweet) {
	for _, v := range tweets {
		m := fts.extractTweetUrl(v.Text)

		// HACK: Tweetに付随したMediaKeysを取得する際、取得できずにヌルポになる事象がある
		if v.Attachments == nil {
			continue
		}

		// 1つのツイートに対してメディアが配列でついているためforで回す
		for _, mediaKey := range v.Attachments.MediaKeys {
			fts.dao.InsertTweetObject(hashtagId, v, m, mediaKey)
		}
	}
}

func (fts *FanartTweetsService) InsertMedia(media []*gotwtr.Media) {
	for _, v := range media {
		fts.dao.InsertMediaObject(v)
	}
}

// ツイート本文に含まれているtweet_urlを抽出する処理
func (fts *FanartTweetsService) extractTweetUrl(text string) map[string]string {
	// https://t.co/<空白以外の一文字以上>
	re := regexp.MustCompile(`https://t.co/\S*$`)
	extractedUrl := re.FindString(text)
	tweetText := strings.Replace(text, extractedUrl, "", -1)

	m := map[string]string{"url": extractedUrl, "text": tweetText}

	return m
}
