package logic

import (
	"os"

	"github.com/tmp-friends/victo-batch/functions/register_vtubers/service"
)

type RegisterVtubersLogic struct {
	service *service.RegisterVtubersService
}

func NewRegisterVtubersLogic() *RegisterVtubersLogic {
	rhs := service.NewRegisterVtubersService()

	return &RegisterVtubersLogic{
		service: rhs,
	}
}

// Schedulerには登録せず手動実行させるバッチ
func (rvl *RegisterVtubersLogic) DoExecute() {
	// jsonファイルからvtuberの情報をload
	vtubers := rvl.service.LoadJsonFile(os.Getenv("SOURCE_DIR") + "assets/vtubers.json")

	// profile_image_url を取得
	pius := rvl.service.FetchProfileImageUrls(vtubers)
	vtubers = rvl.service.AddProfileImageUrl(vtubers, pius)

	// INSERT実行
	rvl.service.RegisterVtubers(vtubers)
	rvl.service.RegisterVtubers(vtubers)
}
