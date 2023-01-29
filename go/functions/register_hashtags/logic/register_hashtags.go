package logic

import (
	"fmt"

	"github.com/tmp-friends/victo-batch/functions/register_hashtags/service"
)

type RegisterHashtagsLogic struct {
	service *service.RegisterHashtagsService
}

func NewRegisterHashtagsLogic() *RegisterHashtagsLogic {
	rhs := service.NewRegisterHashtagsService()

	return &RegisterHashtagsLogic{
		service: rhs,
	}
}

// Schedulerには登録せず手動実行させるバッチ
func (rhl *RegisterHashtagsLogic) DoExecute() int {
	// jsonファイルからvtuberの情報をload

	// profile_image_url を取得
	pius := rhl.service.FetchProfileImage()
	fmt.Println(pius)
	// register vtubers

	// register hashtags

	return 0
}
