package actions

import "github.com/pandarayc/behavior3go/core"

type Runner struct {
	core.Action
}

func (node *Runner) OnTick(tick core.Tick) core.NodeStatus {
	return core.STATUS_RUNNING
}
