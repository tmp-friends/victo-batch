package logic

import (
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
	vtubers := rhl.service.LoadJsonFile("./assets/vtubers.json")

	// profile_image_url を取得
	pius := rhl.service.FetchProfileImageUrls(vtubers)
	vtubers = rhl.service.AddProfileImageUrl(vtubers, pius)

	rhl.service.RegisterVtubers(vtubers)

	return 0
}
