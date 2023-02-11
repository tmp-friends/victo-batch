package function

import (
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/joho/godotenv"

	fanartTweetsLogic "github.com/tmp-friends/victo-batch/functions/insert_fanart_tweets/logic"
	hashtagLogic "github.com/tmp-friends/victo-batch/functions/register_hashtags/logic"
	vtuberLogic "github.com/tmp-friends/victo-batch/functions/register_vtubers/logic"
)

// Targetの定義
func init() {
	loadEnv()

	functions.HTTP("RegisterVtubers", registerVtubers)
	functions.HTTP("RegisterHashtags", registerHashtags)
	functions.HTTP("InsertFanartTweets", insertFanartTweets)
}

func loadEnv() {
	// SOURCE_DIR -> dev: "", prod: "serverless_function_source_code/"
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
