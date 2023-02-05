package lib

import (
	"context"
	"os"

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