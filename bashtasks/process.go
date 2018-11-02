package bashtasks

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/saromanov/bashtasks/util"
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

// Run provides entry point at bashtasks
func (b *BashTasks) Run() {
	root := b.Config
	tasks := root.Tasks
	parallelTasks := root.ParallelTasks
	startMessage(root)
	if len(tasks) == 0 && len(parallelTasks) == 0 {
		return
	}
	b.NumberOfTasks = len(tasks) + len(parallelTasks)
	for _, t := range tasks {
		b.runTask(root, t)
	}
	return
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
		b.runTask(root, t)
	}
	return
}

// runTask provides executing of the task logic
// TODO: Add errors and stages of the task execution
func (b *BashTasks) runTask(root *Config, t Task) error {
	start := time.Now()
	if t.Path != "" {
		fileName, err := util.DownloadFile(t.Path)
		if err != nil {
			color.Red(fmt.Sprintf("unable to download file: %v", err))
			return err
		}
		defer util.RemoveFile(fileName)
		t.ScriptPath = fileName
		out, err := b.executeScript(t)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		end := time.Since(start).Seconds()
		if root.ShowOutput {
			fmt.Println(string(out))
		}
		color.Yellow(fmt.Sprintf("Task was executed on: %fs", end))
		return nil
	}
	if t.ScriptPath != "" {
		out, err := b.executeScript(t)
		if err != nil {
			return err
		}
		if root.ShowOutput {
			fmt.Println(string(out))
		}
	}
	out, err := b.executeTask(t)
	if err != nil {
		if t.AbortPipeline {
			return nil
		}
		return err
	}
	end := time.Since(start).Seconds()
	if root.ShowOutput {
		fmt.Println(string(out))
	}
	color.Yellow(fmt.Sprintf("Task was executed on: %fs", end))
	return nil
}

// GetRule returns rule of current task contains
// tag attached to the rule
func (b *BashTasks) GetRule(t Task) (*Rule, bool) {
	if len(t.Tags) == 0 {
		return nil, false
	}
	if len(b.Config.Rules) == 0 {
		return nil, false
	}
	tags := t.Tags
	for _, tag := range tags {
		for _, r := range b.Config.Rules {
			if tag == r.Tag {
				return &r, true
			}
		}
	}
	return nil, false
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
