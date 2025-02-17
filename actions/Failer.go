package actions

import "github.com/pandarayc/behavior3go/core"

type Failer struct {
	core.Action
}

func (node *Failer) OnTick(tick core.Tick) core.NodeStatus {
	return core.STATUS_FAILURE
}
