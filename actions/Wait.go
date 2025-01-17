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
	node.endTime = cfg.

}

func (node *Wait) OnOpen(tick core.Tick) {
	var startTime int64 = time.Now().Unix()
	tick.BlackBoard.Set()
}
