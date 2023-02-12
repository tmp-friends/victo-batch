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
	nextToken string,
) *gotwtr.SearchTweetsResponse {
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

	// 100件/req しか取得できないので再帰処理
	if res.Meta.NextToken != "" {
		resNext := tc.SearchRecent(keyword, startTime, endTime, res.Meta.NextToken)

		res.Tweets = append(res.Tweets, resNext.Tweets...)
		res.Includes.Media = append(res.Includes.Media, resNext.Includes.Media...)
		res.Meta.ResultCount += resNext.Meta.ResultCount
	}

	return res
}
