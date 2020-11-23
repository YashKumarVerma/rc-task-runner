package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/YashKumarVerma/rc-task-runner/internal/dispatcher"

	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/YashKumarVerma/rc-task-runner/internal/config"

	"github.com/gin-gonic/gin"
)

// type declarations
type entryStruct struct {
	ID  string   `json:"id"`
	URL []string `json:"url"`
}

type responseStruct struct {
	Err     bool          `json:"error"`
	Payload []entryStruct `json:"payload"`
}

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

	if response.Err == true {
		ui.ContextPrint("fire", "error on function side")
		ctx.JSON(500, gin.H{
			"error": true,
		})
	} else {
		downloadCodeBinaries(response.Payload)
		dispatcher.CheckInventory()
	}
}

// function to download the files from given links and place at correct place
func downloadCodeBinaries(binaryDetails []entryStruct) {
	for _, item := range binaryDetails {
		targetLocation, _ := filepath.Abs("./" + config.Load.CodeDirectory + "/" + item.ID)
		err := downloadFile(item.URL[0], targetLocation)
		ui.CheckError(err, "Error downloading binary :"+item.ID, false)
	}
}

// https://golangcode.com/download-a-file-from-a-url/
func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
