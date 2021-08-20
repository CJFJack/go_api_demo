package services

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"io"
	"os/exec"
	"strings"
)

func asyncLog(reader io.ReadCloser) error {
	cache := ""
	buf := make([]byte, 1024, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "closed") {
				return nil
			}
			return err
		}
		if num > 0 {
			oByte := buf[:num]
			oSlice := strings.Split(string(oByte), "\n")
			line := strings.Join(oSlice[:len(oSlice)-1], "\n")
			logs.Info(fmt.Sprintf("%s%s\n", cache, line))
			cache = oSlice[len(oSlice)-1]
		}
	}
}

func RunShellPrintRealtime(scriptPath string) error {
	logs.Info("run shell start")
	cmd := exec.Command("sh", "-c", scriptPath)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		logs.Info(fmt.Sprintf("Error starting command: %s......", err.Error()))
		return err
	}

	go asyncLog(stdout)
	go asyncLog(stderr)

	if err := cmd.Wait(); err != nil {
		logs.Info(fmt.Sprintf("Error waiting for command execution: %s......", err.Error()))
		return err
	}

	logs.Info("run shell finish")
	return nil
}

func RunShell(scriptPath string) error {
	logs.Info("start run shell")

	cmd := exec.Command("/bin/bash", "-c", scriptPath)

	err := cmd.Run()
	if err != nil {
		logs.Error(fmt.Sprintf("Execute Command %s failed: %s", scriptPath, err.Error()))
		return err
	}
	logs.Info(fmt.Sprintf("Execute Command %s finished.", scriptPath))

	logs.Info("run shell finish")
	return nil
}
