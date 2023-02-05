package service

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sivchari/gotwtr"
	"github.com/tmp-friends/victo-batch/functions/dto"
	"github.com/tmp-friends/victo-batch/functions/register_hashtags/dao"
	"github.com/tmp-friends/victo-batch/functions/register_hashtags/lib"
)

const REQ_LIMIT = 100

type RegisterHashtagsService struct {
	lib *lib.TwitterClient
	dao *dao.RegisterHashtagsDao
}

func NewRegisterHashtagsService() *RegisterHashtagsService {
	tc := lib.NewTwitterClient()
	return &RegisterHashtagsService{
		lib: tc,
	}
}

func (rhs *RegisterHashtagsService) LoadJsonFile(filepath string) []dto.Vtuber {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	vtubers := []dto.Vtuber{}
	json.Unmarshal(raw, &vtubers)

	return vtubers
}

func (rhs *RegisterHashtagsService) FetchProfileImageUrls(vtubers []dto.Vtuber) []string {
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

func (rhs *RegisterHashtagsService) extractTwitterUserNames(vtubers []dto.Vtuber) []string {
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
	vtubers []dto.Vtuber,
	profileImageUrls []string,
) []dto.Vtuber {
	emptyIdx := 0
	for i, url := range profileImageUrls {
		if vtubers[i].TwitterUserName == "" {
			emptyIdx++
		}
		vtubers[i+emptyIdx].ProfileImageURL = url
	}

	return vtubers
}
