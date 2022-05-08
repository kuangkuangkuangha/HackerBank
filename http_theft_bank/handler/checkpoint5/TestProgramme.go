package checkpoint5

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type process struct {
	Cmd    *exec.Cmd
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
	File   *os.File // 用于标志打开过的文件，方便关闭
}

// TODO: 测试用例个数 n 从配置文件中导入
var TestResult = [][]string{
	{"1234", "12345", "12345"},
	{"2345", "12345", "12344"},
	{"54321", "2134"},
}

// TODO: 测试文件名从配置文件中导入
var TestFiles = []string{
	"./file/testSample/test1.txt",
	"./file/testSample/test2.txt",
	"./file/testSample/test3.txt",
}

func testProgramme(fileName, fileNameOnly string) error {
	// go build
	cmd := exec.Command("go", "build", "-o", "./file/bin", "./file/"+fileName)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// go run
	var processSet []*process
	for i := 0; i < len(TestFiles); i++ {
		p, err := initProcess(fileNameOnly, TestFiles[i])
		if err != nil {
			killAllProcess(processSet)
			return err
		}
		err = p.Cmd.Start()
		if err != nil {
			// kill all process and return
			killAllProcess(processSet)
			return err
		}
		processSet = append(processSet, p)
	}

	// init channel
	doneChannel := make(chan int)
	errChannel := make(chan error)
	defer close(doneChannel)
	defer close(errChannel)
	doneCount := 0

	// goroutine start
	for i := 0; i < len(processSet); i++ {
		go processMonitor(processSet[i], errChannel, doneChannel)
	}

	// for - select
	for doneCount != len(processSet) {
		select {
		case <-doneChannel:
			doneCount++
			fmt.Println("one done")
		case err := <-errChannel:
			closeAllFile(processSet)
			killAllProcess(processSet)
			return err
		case <-time.After(time.Second * 2):
			// fmt.Println("over time and kill all process")
			closeAllFile(processSet)
			killAllProcess(processSet)
			return errors.New("over time")
		}
	}

	// 关掉所有文件
	closeAllFile(processSet)

	// TODO: check stdout
	for i := 0; i < len(processSet); i++ {
		fmt.Println(processSet[i].Stdout.String())
	}

	return nil
}

// initProcess ... 初始化进程
func initProcess(fileNameOnly, testFileName string) (*process, error) {
	p := new(process)
	p.Cmd = exec.Command("./file/bin/" + fileNameOnly)
	p.Cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // 每个进程单独一个组

	// 打开测试文件
	file, err := os.Open(testFileName)
	if err != nil {
		return nil, err
	}

	p.Cmd.Stdin = file
	p.File = file

	// 获取标准输出
	p.Stdout = new(bytes.Buffer)
	p.Stderr = new(bytes.Buffer)
	p.Cmd.Stdout = p.Stdout
	p.Cmd.Stderr = p.Stderr

	return p, nil
}

// processMonitor ... 用于 goroutine 等待
func processMonitor(p *process, errChannel chan<- error, doneChannel chan<- int) {
	err := p.Cmd.Wait()
	if err != nil {
		errChannel <- err
		return
	}

	doneChannel <- 1
	return
}

func killAllProcess(processSet []*process) {
	for i := 0; i < len(processSet); i++ {
		err := syscall.Kill(-processSet[i].Cmd.Process.Pid, syscall.SIGKILL)
		if err != nil {
			// log
			fmt.Println("kill all process failed: " + err.Error())
		}
	}
}

// closeAllFile ... 关掉所有文件
func closeAllFile(processSet []*process) {
	for i := 0; i < len(processSet); i++ {
		err := processSet[i].File.Close()
		if err != nil {
			// log
			fmt.Println("kill all process failed: " + err.Error())
		}
	}
}
