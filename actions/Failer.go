package actions

import "github.com/genet9496/behavior3go/core"

type Failure struct {
	core.Action
}

func (node *Failure) OnTick(tick core.Tick) core.NodeStatus {
	return core.STATUS_FAILURE
}
