package bashtasks

import (
	"fmt"
	"os/exec"
)

// RowTasks defines response
// for execute row tasks
type RowTasks struct {
	NumberOfTasks int
	CompleteTasks int
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
