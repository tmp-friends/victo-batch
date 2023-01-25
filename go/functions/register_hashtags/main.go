package register_hashtags

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("RegisterHashtags", registerHahstags)
}

func registerHahstags(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
