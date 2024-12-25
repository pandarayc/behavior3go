package core

import "github.com/genet9496/behavior3go/config"

type ICondition interface {
	INode
}

type Condition struct {
	BaseNode
	BaseWorker
}

var _ ICondition = &Condition{}

func (node *Condition) _ctor() {
	node.category = CATEGORY_CONDITION
}

func (node *Condition) Initialize(cfg *config.NodeCfg) {
	node.BaseNode.Initialize(cfg)
}
