package dispatcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	ui "github.com/YashKumarVerma/go-lib-ui"
)

// ValidQuestions contain ids of all questions
var ValidQuestions []string

// CheckInventory reads ./codes directory to check if passed question IDs are valid or not
func CheckInventory() {
	codesPath, err := filepath.Abs("./codes/")
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
	fmt.Println("Running")
	os.Stdout = write

	//Just for testing, replace with your subProcess
	codePath, err := filepath.Abs("./codes/" + programID + "/code.out")
	ui.CheckError(err, "Error creating codePath", true)

	subProcess := exec.Command(codePath)
	stdin, err := subProcess.StdinPipe()
	ui.CheckError(err, "Error creating stdin pipe", true)
	defer stdin.Close()

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	err = subProcess.Start()
	fmt.Println(err)
	ui.CheckError(err, "Error executing code", true)

	io.WriteString(stdin, inputString)

	// close pipes
	write.Close()
	out, _ := ioutil.ReadAll(read)
	os.Stdout = rescueStdout

	return string(out)
}
