package logic

import (
	"os"

	"github.com/tmp-friends/victo-batch/functions/pkg/register_hashtags/service"
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
	hashtags := rhl.service.LoadJsonFile(os.Getenv("SOURCE_DIR") + "assets/hashtags.json")

	// vtuberNameに該当するレコードのVtuberIDを取得する
	hashtags = rhl.service.GetVtuberIds(hashtags)

	// UPSERT実行
	rhl.service.RegisterHashtags(hashtags)
}
