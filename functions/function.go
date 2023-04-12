package function

import (
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/joho/godotenv"

	fanartTweetsLogic "github.com/tmp-friends/victo-batch/functions/pkg/insert_fanart_tweets/logic"
	hashtagLogic "github.com/tmp-friends/victo-batch/functions/pkg/register_hashtags/logic"
	vtuberLogic "github.com/tmp-friends/victo-batch/functions/pkg/register_vtubers/logic"
)

// Targetの定義
func init() {
	loadEnv()

	// 環境変数でTargetを指定する
	functions.HTTP("RegisterVtubers", registerVtubers)
	functions.HTTP("RegisterHashtags", registerHashtags)
	functions.HTTP("InsertFanartTweets", insertFanartTweets)
}

func loadEnv() {
	// SOURCE_DIR -> 開発: "", 本番: "serverless_function_source_code/"
	err := godotenv.Load(os.Getenv("SOURCE_DIR") + ".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func registerVtubers(w http.ResponseWriter, r *http.Request) {
	rvl := vtuberLogic.NewRegisterVtubersLogic()
	rvl.DoExecute()
}

func registerHashtags(w http.ResponseWriter, r *http.Request) {
	rhl := hashtagLogic.NewRegisterHashtagsLogic()
	rhl.DoExecute()
}

func insertFanartTweets(w http.ResponseWriter, r *http.Request) {
	ftl := fanartTweetsLogic.NewFanartTweetsLogic()
	ftl.DoExecute()
}
