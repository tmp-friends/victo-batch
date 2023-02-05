package logic

import (
	"os"

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
func (rhl *RegisterHashtagsLogic) DoExecute() {
	// jsonファイルからvtuberの情報をload
	vtubers := rhl.service.LoadJsonFile(os.Getenv("SOURCE_DIR") + "assets/vtubers.json")

	// profile_image_url を取得
	pius := rhl.service.FetchProfileImageUrls(vtubers)
	vtubers = rhl.service.AddProfileImageUrl(vtubers, pius)

	// INSERT実行
	rhl.service.RegisterVtubers(vtubers)
}
