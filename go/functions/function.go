package function

import (
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/joho/godotenv"

	"github.com/tmp-friends/victo-batch/functions/register_vtubers/logic"
)

// Targetの定義
func init() {
	loadEnv()

	functions.HTTP("RegisterVtubers", registerVtubers)
}

func loadEnv() {
	// SOURCE_DIR -> dev: "", prod: "serverless_function_source_code/"
	err := godotenv.Load(os.Getenv("SOURCE_DIR") + ".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func registerVtubers(w http.ResponseWriter, r *http.Request) {
	rvl := logic.NewRegisterVtubersLogic()
	rvl.DoExecute()
}
