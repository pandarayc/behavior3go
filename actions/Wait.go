package actions

import (
	"time"

	"github.com/pandarayc/behavior3go/config"
	"github.com/pandarayc/behavior3go/core"
)

type Wait struct {
	core.Action
	endTime int64 // milliSecond
}

func (node *Wait) Initialize(cfg *config.NodeCfg) {
	node.Action.Initialize(cfg)
	// 解析配置中的时间
	node.endTime = cfg.GetPropertiesAsInt64(core.PROPERTY_KEY_WAIT_MS)
}

func (node *Wait) OnOpen(tick core.Tick) {
	var startTime int64 = time.Now().UnixMilli()
	tick.BlackBoard.Set(core.BLACKBOARD_KEY_WAIT_START, startTime, tick.GetTree().GetId(), node.GetId())
}

func (node *Wait) OnTick(tick core.Tick) core.NodeStatus {
	now := time.Now().UnixMilli()
	startTime := tick.GetBlackBoard()

}
