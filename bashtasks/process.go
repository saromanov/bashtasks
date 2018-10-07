package bashtasks

import (
	"fmt"
	"os/exec"
)

// executeCommand provides execution of the command
func executeCommand(cmd string) (string, error) {

	out, err := exec.Command(cmd).Output()
	if err != nil {
		return "", fmt.Errorf("unable to execute command")
	}

	return out, nil
}
