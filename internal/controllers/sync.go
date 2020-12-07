package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/YashKumarVerma/rc-task-runner/internal/config"
	"github.com/YashKumarVerma/rc-task-runner/internal/dispatcher"

	"github.com/gin-gonic/gin"
)

// type declarations
type entryStruct struct {
	ID     string `json:"id"`
	Object string `json:"object"`
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
		fixPermissions()
		convertObjectFilesToArchCode(response.Payload)
	}
}

// function to download the files from given links and place at correct place
func downloadCodeBinaries(binaryDetails []entryStruct) {
	for _, item := range binaryDetails {
		targetLocation, _ := filepath.Abs("./" + config.Load.CodeDirectory + "/" + item.ID + ".o")
		err := downloadFile(item.Object, targetLocation)
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
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func fixPermissions() {
	_, err := exec.Command("chmod", "--recursive", "+x", config.Load.CodeDirectory).CombinedOutput()
	ui.CheckError(err, "Unable to set codes as executable", true)
}

func convertObjectFilesToArchCode(payload []entryStruct) {
	for _, object := range payload {
		oFile, err := filepath.Abs("./" + config.Load.CodeDirectory + "/" + object.ID + ".o")
		outFile, err := filepath.Abs("./" + config.Load.CodeDirectory + "/" + object.ID + ".out")
		_, err = exec.Command("/usr/bin/g++", "-o", outFile, oFile).CombinedOutput()
		ui.CheckError(err, "Unable to compile "+object.ID, true)
		ui.ContextPrint("package", "Compiling "+object.ID)
	}
}
