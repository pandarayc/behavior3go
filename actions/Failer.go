package actions

import "github.com/pandarayc/behavior3go/core"

type Failure struct {
	core.Action
}

func (node *Failure) OnTick(tick core.Tick) core.NodeStatus {
	return core.STATUS_FAILURE
}
