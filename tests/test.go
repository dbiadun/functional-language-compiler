package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/albenik/go-serial"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	tester := new(Tester)
	tester.test()
}

type Tester struct {
	testsType    int
	testsDir     string
	compilerPath string
	passedTests  int
	allTests     int
}

const (
	functionalTests = iota
	performanceTests
)

func (t *Tester) test() {
	t.init()
	t.printStart()
	t.performTests()
	t.printEnd()
}

func (t *Tester) init() {
	_, codePath, _, _ := runtime.Caller(0)
	t.testsDir = filepath.Dir(codePath)

	compilerDir := filepath.Dir(t.testsDir)
	t.compilerPath = filepath.Join(compilerDir, "build", "compiler")

	if len(os.Args) < 2 {
		fmt.Printf("Expected '%s' or '%s' subcommands", functionalCmdName, performanceCmdName)
		os.Exit(1)
	}

	switch os.Args[1] {
	case functionalCmdName:
		t.testsType = functionalTests
	case performanceCmdName:
		t.testsType = performanceTests
	}
}

const (
	functionalCmdName  = "functional"
	performanceCmdName = "performance"
)

func (t *Tester) performTests() {
	switch t.testsType {
	case functionalTests:
		t.performFunctionalTests()
	case performanceTests:
		t.performPerformanceTests()
	default:
		return
	}
}

func (t *Tester) printStart() {
	switch t.testsType {
	case functionalTests:
		fmt.Println("Functional tests:")
	case performanceTests:
		fmt.Println("Performance tests:")
	}
}

func (t *Tester) printEnd() {
	fmt.Printf("%d/%d tests passed\n", t.passedTests, t.allTests)
}

func (t *Tester) performFunctionalTests() {
	functionalDir := filepath.Join(t.testsDir, "functionalTests")
	validDir := filepath.Join(functionalDir, "valid")
	invalidDir := filepath.Join(functionalDir, "invalid")

	t.checkTestsDir(validDir, true)
	t.checkTestsDir(invalidDir, false)
}

func (t *Tester) checkTestsDir(testsDir string, expectedValid bool) {
	dirName := filepath.Base(testsDir)
	fmt.Printf("Tests from the '%s' directory:\n", dirName)

	codeDir := filepath.Join(testsDir, "code")

	filepath.WalkDir(codeDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			if path != codeDir {
				return filepath.SkipDir
			}
		} else {
			t.checkTest(path, testsDir, expectedValid)
		}

		return nil
	})
}

func (t *Tester) checkTest(path string, testsDir string, expectedValid bool) {
	filename := filepath.Base(path)
	testName := filename[:len(filename)-3]
	inputName := testName + "Input"
	outputName := testName + "Output"

	inputFile := filepath.Join(testsDir, "inputs", inputName)
	outputFile := filepath.Join(testsDir, "outputs", outputName)
	tmpFile := filepath.Join(t.testsDir, "tmp")

	if len(testName) > 8 && testName[:8] == "feather_" {
		t.checkTestOnFeather(path, testName, outputFile, expectedValid)
	} else {
		t.checkTestOnLinux(path, testName, inputFile, outputFile, tmpFile, expectedValid)
	}
}

func (t *Tester) checkTestOnLinux(path string, testName string, inputFile string, outputFile string, tmpFile string, expectedValid bool) {
	err := t.checkCompilation(testName, path, tmpFile, expectedValid)
	if err != nil {
		return
	}

	runStdout, err := t.checkExecution(testName, tmpFile, inputFile, expectedValid)
	if err != nil {
		return
	}

	err = t.checkOutput(testName, outputFile, runStdout, expectedValid)
	if err != nil {
		return
	}
}

func (t *Tester) checkCompilation(testName string, path string, tmpFile string, expectedValid bool) error {
	compileCmd := exec.Command(t.compilerPath, "build", "-o", tmpFile, path)
	_, compileErr := compileCmd.CombinedOutput()

	valid := compileErr == nil
	return t.reportResult(testName, valid, expectedValid, false, compilationError)
}

func (t *Tester) checkExecution(testName string, tmpFile string, inputFile string, expectedValid bool) ([]byte, error) {
	runCmd := exec.Command(tmpFile)
	if input, err := os.ReadFile(inputFile); err == nil {
		runCmd.Stdin = bytes.NewBuffer(input)
	}

	runStdout, runErr := runCmd.CombinedOutput()

	valid := runErr == nil
	err := t.reportResult(testName, valid, expectedValid, false, runtimeError)

	return runStdout, err
}

func (t *Tester) checkOutput(testName string, outputFile string, runStdout []byte, expectedValid bool) error {
	validOutput := t.readValidOutputFile(outputFile, string(runStdout))
	valid := string(runStdout) == validOutput

	return t.reportResult(testName, valid, expectedValid, true, outputError)
}

func (t *Tester) checkTestOnFeather(path string, testName string, outputFile string, expectedValid bool) {
	validOutput := t.readValidOutputFile(outputFile, "")

	err := t.checkFlashToFeather(testName, path, expectedValid)
	if err != nil {
		return
	}

	t.checkFeatherOutput(testName, validOutput, expectedValid)
}

func (t *Tester) checkFlashToFeather(testName string, path string, expectedValid bool) error {
	flashCmd := exec.Command(t.compilerPath, "flash", "-target", "feather", path)
	_, flashErr := flashCmd.CombinedOutput()
	valid := flashErr == nil

	return t.reportResult(testName, valid, expectedValid, false, compilationError)
}

func (t *Tester) readValidOutputFile(outputFile string, defaultOutput string) string {
	validOutput := ""
	if outputContent, err := os.ReadFile(outputFile); err == nil {
		validOutput = string(outputContent)
	} else {
		validOutput = defaultOutput
	}

	return validOutput
}

func (t *Tester) checkFeatherOutput(testName string, validOutput string, expectedValid bool) {
	timeout := time.Second * 5
	startTime := time.Now()

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	time.Sleep(time.Millisecond * 500) // To allow microcontroller to renew the connection after reset
	port, err := serial.Open("/dev/ttyACM0", mode)
	if err != nil {
		t.reportResult(testName, !expectedValid, expectedValid, true, deviceError)
		return
	}

	buffer := make([]byte, len(validOutput))

	wholeLen := len(validOutput)
	curLen := 0

	for time.Since(startTime) < timeout && curLen < wholeLen {

		localBuffer := make([]byte, wholeLen-curLen)
		n, _ := port.Read(localBuffer)
		copy(buffer[curLen:curLen+n], localBuffer[:n])
		curLen += n
	}

	port.Close()

	valid := string(buffer) == validOutput

	if curLen != wholeLen {
		valid = false
	}

	t.reportResult(testName, valid, expectedValid, true, outputError)
}

func (t *Tester) performPerformanceTests() {

}

func (t *Tester) reportResult(testName string, valid bool, expectedValid bool, finalCheck bool, err string) error {
	passed := valid == expectedValid

	errToReport := ""
	if !valid && expectedValid {
		errToReport = err
	}

	// Only in this case the test stops and we need to report its result
	if !valid || finalCheck {
		if passed {
			t.reportPassed(testName)
		} else {
			t.reportFailed(testName, errToReport)
		}

		return errors.New("end of execution")
	}

	return nil
}

func (t *Tester) reportPassed(testName string) {
	t.allTests++
	t.passedTests++
	fmt.Printf("%s %s\n", greenStr("[PASSED]"), testName)
}

func (t *Tester) reportFailed(testName string, err string) {
	t.allTests++

	if err == "" {
		fmt.Printf("%s %s\n", redStr("[FAILED]"), testName)
	} else {
		fmt.Printf("%s %s: %s\n", redStr("[FAILED]"), testName, err)
	}
}

func redStr(str string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", str)
}

func greenStr(str string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", str)
}

const (
	compilationError = "Compilation error"
	runtimeError     = "Runtime error"
	outputError      = "Invalid output"
	deviceError      = "Device not available"
)
