package bashtasks

import (
	"fmt"
	"os/exec"
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
func ExecuteRowTasks(tasks []Task) *RowTasks {
	if len(tasks) == 0 {
		return nil
	}
	rt := &RowTasks{NumberOfTasks: len(tasks)}
	for _, t := range tasks {
		out, err := executeCommand(t.Cmd)
		if err != nil {
			continue
		}
		rt.CompleteTasks++
		fmt.Println(string(out))
	}
	return rt
}

// executeCommand provides execution of the command
func executeCommand(cmd string) ([]byte, error) {

	out, err := exec.Command(cmd).Output()
	if err != nil {
		return nil, fmt.Errorf("unable to execute command")
	}

	return out, nil
}
