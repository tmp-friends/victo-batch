// Deprecated
// @see: lib/twitter_scraper.go
package lib

import (
	"context"
	"os"
	"time"

	"github.com/sivchari/gotwtr"
)

type TwitterClient struct {
	client *gotwtr.Client
}

func NewTwitterClient() *TwitterClient {
	client := gotwtr.New(os.Getenv("BEARER_TOKEN"))

	return &TwitterClient{
		client: client,
	}
}

func (tc *TwitterClient) GetUsersBy(userNames []string) *gotwtr.UsersResponse {
	// @see: https://pkg.go.dev/github.com/sivchari/gotwtr#Client.RetrieveMultipleUsersWithUserNames
	res, err := tc.client.RetrieveMultipleUsersWithUserNames(
		context.Background(),
		userNames,
		&gotwtr.RetrieveUserOption{UserFields: []gotwtr.UserField{"profile_image_url"}},
	)
	if err != nil {
		panic(err)
	}

	return res
}

func (tc *TwitterClient) SearchRecent(
	keyword string,
	startTime time.Time,
	endTime time.Time,
) ([]*gotwtr.Tweet, []*gotwtr.Media, int) {
	var tweetsRes []*gotwtr.Tweet
	var mediaRes []*gotwtr.Media
	var resultCount int

	var nextToken string
	// 無限ループを避けるため連続10回(=1000件)までのrequestとする
	for i := 0; i < 10; i++ {
		// @see: https://pkg.go.dev/github.com/sivchari/gotwtr#Client.SearchRecentTweets
		res, err := tc.client.SearchRecentTweets(
			context.Background(),
			keyword,
			&gotwtr.SearchTweetsOption{
				StartTime:  startTime,
				EndTime:    endTime,
				MaxResults: 100,
				NextToken:  nextToken,
				Expansions: []gotwtr.Expansion{
					gotwtr.ExpansionAuthorID,
					gotwtr.ExpansionAttachmentsMediaKeys,
				},
				TweetFields: []gotwtr.TweetField{
					gotwtr.TweetFieldCreatedAt,
					gotwtr.TweetFieldPublicMetrics,
				},
				MediaFields: []gotwtr.MediaField{
					gotwtr.MediaFieldPreviewImageURL,
					gotwtr.MediaFieldURL,
				},
			},
		)
		if err != nil {
			panic(err)
		}

		// res.Includes.Mediaのappend時にヌルポになるため早期break
		if res.Meta.ResultCount == 0 {
			break
		}

		tweetsRes = append(tweetsRes, res.Tweets...)
		mediaRes = append(mediaRes, res.Includes.Media...)
		resultCount += res.Meta.ResultCount

		// NextTokenが空になる or 前回のreqとTokenが同じ になるまで回す
		if res.Meta.NextToken == "" || res.Meta.NextToken == nextToken {
			break
		}

		nextToken = res.Meta.NextToken
	}

	return tweetsRes, mediaRes, resultCount
}
