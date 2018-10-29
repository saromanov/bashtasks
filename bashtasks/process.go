package bashtasks

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/fatih/color"
)

// BashTasks defines response
// for execute row tasks
type BashTasks struct {
	NumberOfTasks int
	CompleteTasks int
	Config        *Config
}

// New provides init of the bash tasks
func New(cfg *Config) *BashTasks {
	return &BashTasks{Config: cfg}
}

// ExecuteRowTasks provides executing of the
// commands step by step
func (b *BashTasks) ExecuteRowTasks() {
	root := b.Config
	tasks := root.Tasks
	if len(tasks) == 0 {
		return
	}
	b.NumberOfTasks = len(tasks)
	for _, t := range tasks {
		b.runTask(t)
	}
	return
}

// runTask provides executing of the task logic
// TODO: Add errors and stages of the task execution
func (b *BashTasks) runTask(t Task) error {
	start := time.Now()
		if t.Path != "" {
			fileName, err := downloadScript(t.Path)
			if err != nil {
				color.Red(fmt.Sprintf("unable to download file: %v", err))
				return err
			}
			t.ScriptPath = fileName
			out, err := b.executeScript(t)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			if root.ShowOutput {
				fmt.Println(string(out))
			}
			return nil
		}
		if t.ScriptPath != "" {
			out, err := b.executeScript(t)
			if err != nil {
				continue
			}
			if root.ShowOutput {
				fmt.Println(string(out))
			}
		}
		out, err := b.executeTask(t)
		if err != nil {
			if t.AbortPipeline {
				return
			}
			continue
		}
		end := time.Since(start).Seconds()
		if root.ShowOutput {
			fmt.Println(string(out))
		}
		color.Yellow(fmt.Sprintf("Task was executed on: %fs", end))
}

// executeScript provides execution of the sript
func (b *BashTasks) executeScript(t Task) ([]byte, error) {
	t.Cmd = fmt.Sprintf("%s", t.ScriptPath)
	return b.executeTask(t)
}

func (b *BashTasks) executeTask(t Task) ([]byte, error) {
	var (
		out []byte
		err error
	)

	color.Yellow(fmt.Sprintf("Executing of the task: %s", t.Title))
	if t.Path != "" || t.ScriptPath != "" {
		out, err = executeBashScript(t.Cmd)
		if err != nil {
			return nil, err
		}
	} else {
		out, err = executeCommand(t.Cmd)
		if err != nil {
			return nil, err
		}
	}
	b.CompleteTasks++
	return out, nil
}

// Response provides output message
// after execution of tasks
func (b *BashTasks) Response() {
	responseCompleteTasks(fmt.Sprintf("Complete executed tasks: %d", b.CompleteTasks))
}

// executeCommand provides execution of the command
func executeCommand(cmd string) ([]byte, error) {

	out, err := exec.Command(cmd).Output()
	if err != nil {
		return nil, fmt.Errorf("unable to execute command")
	}

	return out, nil
}

// executeBashScript provides executing of the bash script
func executeBashScript(cmd string) ([]byte, error) {

	out, err := exec.Command("/bin/sh", cmd).Output()
	if err != nil {
		return nil, fmt.Errorf("unable to execute bash script", err)
	}

	return out, nil
}

// downloadScript provides downloading of the bash script
// Its copy to the temp file
func downloadScript(url string) (string, error) {
	client := &http.Client{}
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", nil
	}
	fileName := path.Base(r.URL.Path)
	out, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
