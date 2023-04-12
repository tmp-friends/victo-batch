package lib

import (
	"context"
	"log"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

const LIMIT = 300

type TwitterScraper struct {
	scraper *twitterscraper.Scraper
}

func NewTwitterScraper() *TwitterScraper {
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	scraper.WithDelay(1)

	return &TwitterScraper{
		scraper,
	}
}

func (ts *TwitterScraper) SearchTweets(keyword string) []*twitterscraper.Tweet {
	var tweets []*twitterscraper.Tweet

	for res := range ts.scraper.SearchTweets(context.Background(), keyword, LIMIT) {
		if res.Error != nil {
			log.Fatal(res.Error)
		}
		tweets = append(tweets, &res.Tweet)
	}

	return tweets
}

func (ts *TwitterScraper) GetProfile(username string) *twitterscraper.Profile {
	profile, err := ts.scraper.GetProfile(username)
	if err != nil {
		log.Fatal(err)
	}

	return &profile
}
