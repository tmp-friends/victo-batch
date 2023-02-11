package lib

import (
	"context"
	"fmt"
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
) *gotwtr.SearchTweetsResponse {
	keyword = fmt.Sprintf("#%s -is:retweet has:images", keyword)
	// @see: https://pkg.go.dev/github.com/sivchari/gotwtr#Client.SearchRecentTweets
	res, err := tc.client.SearchRecentTweets(
		context.Background(),
		keyword,
		&gotwtr.SearchTweetsOption{
			StartTime:  startTime,
			EndTime:    endTime,
			MaxResults: 100,
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

	return res
}
