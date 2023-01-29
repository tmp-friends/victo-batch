package service

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/tmp-friends/victo-batch/functions/register_hashtags/lib"
)

type RegisterHashtagsService struct {
	lib *lib.TwitterClient
}

func NewRegisterHashtagsService() *RegisterHashtagsService {
	tc := lib.NewTwitterClient()
	return &RegisterHashtagsService{
		lib: tc,
	}
}

type Vtuber struct {
	Name          string `json:"name"`
	BelongsTo     string `json:"belongs_to"`
	TwitterUserId string `json:"twitter_user_id"`
	Channel       string `json:"channel"`
}

func (rhs *RegisterHashtagsService) LoadJsonFile(filepath string) []Vtuber {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	vtubers := []Vtuber{}
	json.Unmarshal(raw, &vtubers)

	return vtubers
}

func (rhs *RegisterHashtagsService) FetchProfileImageUrls() []string {
	res := rhs.lib.GetUsersBy()

	profileImageUrls := []string{}
	for _, u := range res.Users {
		piu := u.ProfileImageURL

		// "_normal"がつくと画像サイズが小さくなるので削除
		piu = strings.Replace(piu, "_normal", "", -1)
		profileImageUrls = append(profileImageUrls, piu)
	}

	return profileImageUrls
}
