package behavior3go

import (
	"time"

	"github.com/pandarayc/behavior3go/config"
)

type MaxTime struct {
	Decorator
	maxTime int64
}

func (node *MaxTime) Initialize(cfg *config.NodeCfg) {
	node.Decorator.Initialize(cfg)
	node.maxTime = cfg.GetPropertyAsInt64(PROPERTY_KEY_MAX_TIME)
	if node.maxTime < 1 {
		panic("maxTime parameter in MaxTime decorator is an obligatory parameter")
	}
}

func (node *MaxTime) OnOpen(tick *Tick) {
	var startTime int64 = time.Now().UnixMilli()
	tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_START_TS, startTime)
}

func (node *MaxTime) OnTick(tick *Tick) NodeStatus {
	if node.GetChild() == nil {
		return STATUS_ERROR
	}

	var currentTime int64 = time.Now().UnixMilli()
	var startTime int64 = tick.GetMemory(node.GetId()).GetInt64(BLACKBOARD_KEY_START_TS)

	status := node.GetChild().Execute(tick)
	if currentTime-startTime > node.maxTime {
		return STATUS_FAILURE
	}

	return status
}
