/*
InsertFanartTweets Service Unittest
*/
package service

import (
	"fmt"
	"testing"
	"time"
)

var LAYOUT = "2006-01-02_15:04:05_JST"

func TestFetchTweets(t *testing.T) {
	type Parameters struct {
		hashtagName string
		startTime   time.Time
		endTime     time.Time
	}
	type Expected struct {
		hashtagName string
		startTime   time.Time
		endTime     time.Time
	}

	providers := []struct {
		pattern    string
		parameters Parameters
	}{
		{
			pattern: "default",
			parameters: Parameters{
				hashtagName: "みとあーと",
				startTime:   time.Now(),
				endTime:     time.Now(),
			},
		},
	}

	for _, v := range providers {
		t.Run(v.pattern, func(t *testing.T) {
			// Arrange
			fts := NewFanartTweetsService()
			// Act
			tweets := fts.FetchTweets(
				v.parameters.hashtagName,
				v.parameters.startTime.Format(LAYOUT),
				v.parameters.startTime.Format(LAYOUT),
			)

			fmt.Print(tweets)
			// Assert
		})
	}
}
