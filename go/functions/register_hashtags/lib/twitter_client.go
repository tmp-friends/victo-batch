package lib

import (
	"context"

	"github.com/sivchari/gotwtr"
	"github.com/tmp-friends/victo-batch/functions/config"
)

type TwitterClient struct {
	client *gotwtr.Client
}

func NewTwitterClient() *TwitterClient {
	env := config.LoadEnv()

	client := gotwtr.New(env["BEARER_TOKEN"])

	return &TwitterClient{
		client: client,
	}
}

func (tc *TwitterClient) GetUsersBy() *gotwtr.UsersResponse {
	// @see: https://pkg.go.dev/github.com/sivchari/gotwtr#Client.RetrieveMultipleUsersWithUserNames
	res, err := tc.client.RetrieveMultipleUsersWithUserNames(
		context.Background(),
		[]string{"temple_circle"},
		&gotwtr.RetrieveUserOption{UserFields: []gotwtr.UserField{"profile_image_url"}},
	)
	if err != nil {
		panic(err)
	}

	return res
}
