package bashtasks

import (
	"sync"
)

// ExecuteParallelTasks runs executing of bash tasks
// in parallel i.e with goroutines
func ExecuteParallelTasks(tasks []Task) error {
	if len(tasks) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, t := range tasks {
		go func() {
			executeCommand(t.Cmd)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
