package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tmp-friends/victo-batch/functions/dto"
	"github.com/tmp-friends/victo-batch/functions/register_hashtags/dao"
)

type RegisterHashtagsService struct {
	dao *dao.RegisterHashtagsDao
}

func NewRegisterHashtagsService() *RegisterHashtagsService {
	dao := dao.NewRegisterHashtagsDao()

	return &RegisterHashtagsService{
		dao: dao,
	}
}

func (rhs *RegisterHashtagsService) LoadJsonFile(filepath string) []*dto.Hashtag {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	hashtags := []*dto.Hashtag{}
	json.Unmarshal(raw, &hashtags)

	return hashtags
}

func (rhs *RegisterHashtagsService) GetVtuberIds(hashtags []*dto.Hashtag) []*dto.Hashtag {
	var vtuberNames []string
	for _, v := range hashtags {
		vtuberNames = append(vtuberNames, v.VtuberName)
	}

	vtuberInfo := rhs.dao.GetVtubersByNames(vtuberNames)
	fmt.Print()
	for _, hashtag := range hashtags {
		for _, v := range vtuberInfo {
			if hashtag.VtuberName != v.Name {
				continue
			}
			hashtag.VtuberId = v.ID
		}

		if hashtag.VtuberId == 0 {
			panic("VtuberIDが空の要素があります")
		}
	}

	return hashtags
}

func (rhs *RegisterHashtagsService) RegisterHashtags(hashtags []*dto.Hashtag) {
	is_exists := rhs.dao.IsExistsHashtags()
	rhs.dao.RegisterHashtags(hashtags, is_exists)
}
