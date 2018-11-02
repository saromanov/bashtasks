package bashtasks

import (
	"fmt"
	"sync"
)

// executeParallelTasks runs executing of bash tasks
// in parallel i.e with goroutines
func (b *BashTasks) executeParallelTasks(tasks []Task) error {
	if len(tasks) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(len(tasks))
	fmt.Println(len(tasks))
	root := b.Config
	for _, t := range tasks {
		go func(ta Task) {
			b.runTask(root, ta)
			wg.Done()
		}(t)
	}
	wg.Wait()
	return nil
}
