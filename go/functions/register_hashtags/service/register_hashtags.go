package service

import (
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

func (rhs *RegisterHashtagsService) FetchProfileImage() []string {
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
