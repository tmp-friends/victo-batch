package function

import (
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"github.com/tmp-friends/victo-batch/functions/register_hashtags/logic"
)

// Targetの定義
func init() {
	functions.HTTP("RegisterHashtags", registerHashtags)
}

func registerHashtags(w http.ResponseWriter, r *http.Request) {
	rhl := logic.NewRegisterHashtagsLogic()
	returnParam := rhl.DoExecute()

	os.Exit(returnParam)
}
