package service

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sivchari/gotwtr"
	"github.com/tmp-friends/victo-batch/functions/register_hashtags/lib"
)

const REQ_LIMIT = 100

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
	Name            string `json:"name"`
	BelongsTo       string `json:"belongs_to"`
	ProfileImageURL string `json:"profile_image_url"`
	TwitterUserName string `json:"twitter_user_name"`
	Channel         string `json:"channel"`
	HashtagName     string `json:"hashtag_name"`
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

func (rhs *RegisterHashtagsService) FetchProfileImageUrls(vtubers []Vtuber) []string {
	userNames := rhs.extractTwitterUserNames(vtubers)

	usersRes := rhs.requestAPI(userNames)

	profileImageUrls := []string{}
	for i := 0; i < len(usersRes); i++ {
		for _, u := range usersRes[i].Users {
			piu := u.ProfileImageURL

			// "_normal"がつくと画像サイズが小さくなるので削除
			piu = strings.Replace(piu, "_normal", "", -1)
			profileImageUrls = append(profileImageUrls, piu)
		}
	}
	return profileImageUrls
}

func (rhs *RegisterHashtagsService) extractTwitterUserNames(vtubers []Vtuber) []string {
	userNames := []string{}
	for _, v := range vtubers {
		un := v.TwitterUserName

		if un == "" {
			continue
		}

		un = strings.Replace(un, "https://twitter.com/", "", -1)
		un = strings.Replace(un, "https://www.twitter.com/", "", -1)

		userNames = append(userNames, un)
	}

	return userNames
}

func (rhs *RegisterHashtagsService) requestAPI(userNames []string) []*gotwtr.UsersResponse {
	// 小数点以下切り捨ての除算
	process := len(userNames) / REQ_LIMIT

	usersRes := []*gotwtr.UsersResponse{}
	for i := 0; i <= process; i++ {
		requns := userNames[i*REQ_LIMIT : (i+1)*REQ_LIMIT]
		if i == process {
			requns = userNames[i*REQ_LIMIT:]
		}

		res := rhs.lib.GetUsersBy(requns)
		usersRes = append(usersRes, res)
	}

	return usersRes
}

func (rhs *RegisterHashtagsService) AddProfileImageUrl(
	vtubers []Vtuber,
	profileImageUrls []string,
) []Vtuber {
	emptyIdx := 0
	for i, url := range profileImageUrls {
		if vtubers[i].TwitterUserName == "" {
			emptyIdx++
		}
		vtubers[i+emptyIdx].ProfileImageURL = url
	}

	return vtubers
}
