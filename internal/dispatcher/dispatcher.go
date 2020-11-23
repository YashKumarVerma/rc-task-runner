package dispatcher

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/YashKumarVerma/rc-task-runner/internal/config"
)

// ValidQuestions contain ids of all questions
var ValidQuestions []string

// CheckInventory reads ./codes directory to check if passed question IDs are valid or not
func CheckInventory() {
	codesPath, err := filepath.Abs("./" + config.Load.CodeDirectory + "/")
	ui.CheckError(err, "Cannot generate absolute path to ./codes", true)

	files, err := ioutil.ReadDir(codesPath)
	ui.CheckError(err, "Cannot read contents of ./codes", true)

	validQuestionIDs := make([]string, 0)
	for _, element := range files {
		validQuestionIDs = append(validQuestionIDs, element.Name())
	}

	ValidQuestions = validQuestionIDs
}

// DispatchOutput accepts programID and input string, and returns output of code after execution
func DispatchOutput(programID string, inputString string) string {
	// some mojo here
	rescueStdout := os.Stdout
	read, write, _ := os.Pipe()
	os.Stdout = write

	//Just for testing, replace with your subProcess
	codePath, err := filepath.Abs("./" + config.Load.CodeDirectory + "/" + programID)
	ui.CheckError(err, "Error creating codePath", true)

	subProcess := exec.Command(codePath)
	stdin, err := subProcess.StdinPipe()
	ui.CheckError(err, "Error creating stdin pipe", true)
	defer stdin.Close()

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	subProcess.Start()
	ui.CheckError(err, "Error executing code", true)

	io.WriteString(stdin, inputString)

	// close pipes
	write.Close()
	out, _ := ioutil.ReadAll(read)
	os.Stdout = rescueStdout

	return string(out)
}
