package actions

import "github.com/genet9496/behavior3go/core"

type Error struct {
	core.Action
}

var _ core.INode = &Error{}

func NewError() core.INode {
	return core.NewAction("Error", "Error", nil)
}

func (e *Error) OnTick(tick *core.Tick) core.NodeStatus {
	return core.STATUS_ERROR
}
