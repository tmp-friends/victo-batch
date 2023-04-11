package service

import (
	"fmt"

	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/tmp-friends/victo-batch/functions/insert_fanart_tweets/dao"
	"github.com/tmp-friends/victo-batch/functions/lib"
	"github.com/tmp-friends/victo-batch/functions/models"
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
	keyword := fmt.Sprintf(
		"'%s' -filter:retweet filter:images since:%s until:%s",
		hashtagName,
		startTime,
		endTime,
	)

	tweets := fts.lib.SearchTweets(keyword)

	return tweets
}

func (fts *FanartTweetsService) Insert(hashtagId int, tweets []*twitterscraper.Tweet) {
	for _, t := range tweets {
		fts.dao.InsertTweetObject(t, hashtagId)
		fts.dao.InsertMediaObject(t)
	}
}
