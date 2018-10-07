package bashtasks

import "github.com/fatih/color"

// StartMessage returns output before executing of tasks
func StartMessage(cfg *Config) {
	if len(cfg.Tasks) > 0 {
		color.Blue("You have %d tasks for execute", len(cfg.Tasks))
	}
	if len(cfg.ParallelTasks) > 0 {
		color.Blue("You have %d parallel tasks for execute", len(cfg.ParallelTasks))
	}
}
