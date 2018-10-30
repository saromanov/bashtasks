package util

import "os"

// RemoveFile provides removing of the file(script)
// after execution
func RemoveFile(path string) error {
	return os.Remove(path)
}
