package service

import (
	"fmt"
	"regexp"
	"strings"

	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/tmp-friends/victo-batch/functions/models"
	"github.com/tmp-friends/victo-batch/functions/pkg/insert_fanart_tweets/dao"
	"github.com/tmp-friends/victo-batch/functions/pkg/lib"
)

type FanartTweetsService struct {
	lib *lib.TwitterScraper
	dao *dao.FanartTweetsDao
}

func NewFanartTweetsService() *FanartTweetsService {
	ts := lib.NewTwitterScraper()
	ftd := dao.NewFanartTweetsDao()

	return &FanartTweetsService{
		lib: ts,
		dao: ftd,
	}
}

func (fts *FanartTweetsService) GetHashtags() models.HashtagSlice {
	return fts.dao.GetHashtags()
}

func (fts *FanartTweetsService) FetchTweets(
	hashtagName string,
	startTime string,
	endTime string,
) []*twitterscraper.Tweet {
	// HACK: 数字が含まれていると、ハッシュタグをつけての検索が効かなくなる
	re := regexp.MustCompile(`\d`)
	isContainsNumber := re.Match([]byte(hashtagName))

	var keyword string
	if isContainsNumber {
		keyword = fmt.Sprintf(
			"\"%s\" -filter:retweet filter:images since:%s until:%s",
			hashtagName,
			startTime,
			endTime,
		)
	} else {
		keyword = fmt.Sprintf(
			"\"#%s\" -filter:retweet filter:images since:%s until:%s",
			hashtagName,
			startTime,
			endTime,
		)
	}

	tweets := fts.lib.SearchTweets(keyword)

	return tweets
}

func (fts *FanartTweetsService) GetAuthors(tweets []*twitterscraper.Tweet) []*twitterscraper.Profile {
	var authors []*twitterscraper.Profile

	// authorId配列に対して一括で取得できないため、一件ずつ取得
	for _, t := range tweets {
		profile := fts.lib.GetProfile(t.Username)
		authors = append(authors, profile)
	}

	return authors
}

func (fts *FanartTweetsService) Insert(hashtagId int, tweets []*twitterscraper.Tweet) {
	for _, t := range tweets {
		t.Text = fts.extractTweetUrl(t.Text)

		fts.dao.InsertTweetObject(t, hashtagId)
		fts.dao.InsertMediaObject(t)
	}
}

// ツイート本文に含まれているtweet_urlを抽出する処理
func (fts *FanartTweetsService) extractTweetUrl(text string) string {
	// https://t.co/<空白以外の一文字以上>
	re := regexp.MustCompile(`https://t.co/\S*$`)
	url := re.FindString(text)

	extractedText := strings.Replace(text, url, "", -1)

	return extractedText
}

func (fts *FanartTweetsService) InsertAuthors(authors []*twitterscraper.Profile) {
	for _, a := range authors {
		fts.dao.InsertAuthor(a)
	}
}
