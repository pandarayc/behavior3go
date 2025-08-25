package behavior3go

import (
	"time"

	"github.com/pandarayc/behavior3go/config"
)

type Wait struct {
	Action
	endTime int64 // milliSecond
}

func (node *Wait) Initialize(cfg *config.NodeCfg) {
	node.Action.Initialize(cfg)
	// 解析配置中的时间
	node.endTime = node.GetPropertyAsInt64(PROPERTY_KEY_WAIT_MS)
}

func (node *Wait) OnOpen(tick *Tick) {
	var startTime int64 = time.Now().UnixMilli()
	tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_START_TS, startTime)
}

func (node *Wait) OnTick(tick *Tick) NodeStatus {
	now := time.Now().UnixMilli()
	startTime := tick.GetMemory(node.GetId()).GetInt64(BLACKBOARD_KEY_START_TS)
	if now-startTime >= node.endTime {
		return STATUS_SUCCESS
	}

	return STATUS_RUNNING
}
