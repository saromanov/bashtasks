package bashtasks

import (
	"fmt"
	"os/exec"
)

// ExecuteRowTasks provides executing of the
// commands step by step
func ExecuteRowTasks(tasks []Task) {
	if len(tasks) == 0 {
		return
	}
	for _, t := range tasks {
		executeCommand(t.Cmd)
	}
}

// executeCommand provides execution of the command
func executeCommand(cmd string) ([]byte, error) {

	out, err := exec.Command(cmd).Output()
	if err != nil {
		return nil, fmt.Errorf("unable to execute command")
	}

	return out, nil
}
