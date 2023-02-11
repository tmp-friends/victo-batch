package service

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sivchari/gotwtr"
	"github.com/tmp-friends/victo-batch/functions/struct"
	"github.com/tmp-friends/victo-batch/functions/lib"
	"github.com/tmp-friends/victo-batch/functions/register_vtubers/dao"
)

const REQ_LIMIT = 100

type RegisterVtubersService struct {
	lib *lib.TwitterClient
	dao *dao.RegisterVtubersDao
}

func NewRegisterVtubersService() *RegisterVtubersService {
	tc := lib.NewTwitterClient()
	dao := dao.NewRegisterVtubersDao()

	return &RegisterVtubersService{
		lib: tc,
		dao: dao,
	}
}

func (rvs *RegisterVtubersService) LoadJsonFile(filepath string) []*struct.Vtuber {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	vtubers := []*struct.Vtuber{}
	json.Unmarshal(raw, &vtubers)

	return vtubers
}

func (rvs *RegisterVtubersService) FetchProfileImageUrls(vtubers []*struct.Vtuber) []string {
	userNames := rvs.extractTwitterUserNames(vtubers)

	usersRes := rvs.requestAPI(userNames)

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

func (rvs *RegisterVtubersService) extractTwitterUserNames(vtubers []*struct.Vtuber) []string {
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

func (rvs *RegisterVtubersService) requestAPI(userNames []string) []*gotwtr.UsersResponse {
	// 小数点以下切り捨ての除算
	process := len(userNames) / REQ_LIMIT

	usersRes := []*gotwtr.UsersResponse{}
	for i := 0; i <= process; i++ {
		requns := userNames[i*REQ_LIMIT : (i+1)*REQ_LIMIT]
		if i == process {
			requns = userNames[i*REQ_LIMIT:]
		}

		res := rvs.lib.GetUsersBy(requns)
		usersRes = append(usersRes, res)
	}

	return usersRes
}

// 引数のvtubersを更新する必要があるため参照渡し
func (rvs *RegisterVtubersService) AddProfileImageUrl(
	vtubers []*struct.Vtuber,
	profileImageUrls []string,
) []*struct.Vtuber {
	var emptyIdx int
	for i, v := range vtubers {
		if v.TwitterUserName == "" {
			emptyIdx++
			continue
		}

		if (i - emptyIdx) > len(profileImageUrls) {
			break
		}
		v.ProfileImageURL = profileImageUrls[i-emptyIdx]
	}

	return vtubers
}

func (rvs *RegisterVtubersService) RegisterVtubers(vtubers []*struct.Vtuber) {
	rvs.dao.RegisterVtubers(vtubers)
}
