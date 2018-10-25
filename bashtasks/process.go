package bashtasks

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
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
		start := time.Now()
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
	return
}

func (b *BashTasks) executeTask(t Task) ([]byte, error) {
	color.Yellow(fmt.Sprintf("Executing of the task: %s", t.Title))
	out, err := executeCommand(t.Cmd)
	if err != nil {
		return nil, err
	}
	b.CompleteTasks++
	return out, nil
}

// Response provides output message
// after execution of tasks
func (b *BashTasks) Response() {
	color.Green("Complete executed tasks: %d", b.CompleteTasks)
}

// executeCommand provides execution of the command
func executeCommand(cmd string) ([]byte, error) {

	out, err := exec.Command(cmd).Output()
	if err != nil {
		return nil, fmt.Errorf("unable to execute command")
	}

	return out, nil
}

// downloadScript provides downloading of the bash script
func downloadScript(path string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
