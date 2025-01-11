package actions

import "github.com/pandarayc/behavior3go/core"

type Error struct {
	core.Action
}

//  通过配置表注入的正常不用重新赋值name

func (node *Error) OnTick(tick core.Tick) core.NodeStatus {
	return core.STATUS_ERROR
}
