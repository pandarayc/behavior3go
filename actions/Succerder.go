package actions

import "github.com/genet9496/behavior3go/core"

type Succeeder struct {
	core.Action
}

func (node *Succeeder) OnTick(tick core.Tick) core.NodeStatus {
	return core.STATUS_SUCCESS
}
