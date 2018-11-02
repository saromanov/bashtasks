package bashtasks

import "github.com/fatih/color"

// startMessage returns output before executing of tasks
func startMessage(cfg *Config) {
	if len(cfg.Tasks) > 0 {
		color.Blue("You have %d tasks for execute", len(cfg.Tasks))
	}
	if len(cfg.ParallelTasks) > 0 {
		color.Blue("You have %d parallel tasks for execute", len(cfg.ParallelTasks))
	}
}

// ResponseCompleteTasks retruns message about
// complete executed tasks
func responseCompleteTasks(msg string) {
	color.Green(msg)
}
