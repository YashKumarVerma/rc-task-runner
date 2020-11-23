package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/YashKumarVerma/rc-task-runner/internal/config"

	"github.com/gin-gonic/gin"
)

// Sync local program storage with firebase storage via lambda endpoint
func Sync(ctx *gin.Context) {
	res, err := http.Get(config.Load.SeedSource)
	ui.CheckError(err, "Error contacting seeding source", false)
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	ui.CheckError(readErr, "Error reading request body", false)

	response := responseStruct{}
	parseErr := json.Unmarshal(body, &response)
	ui.CheckError(parseErr, "Error mapping response to struct", false)

	// now we have an array containing the urls to download
}

type entryStruct struct {
	ID  string   `json:"id"`
	URL []string `json:"url"`
}

type responseStruct struct {
	Err bool          `json:"error"`
	Key []entryStruct `json:"payload"`
}
