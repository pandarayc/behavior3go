package core

import "github.com/pandarayc/behavior3go/config"

type ICondition interface {
	INode
}

type Condition struct {
	BaseNode
	BaseWorker
}

var _ ICondition = &Condition{}

func (node *Condition) Ctor() {
	node.category = CATEGORY_CONDITION
}

func (node *Condition) Initialize(cfg *config.NodeCfg) {
	node.BaseNode.Initialize(cfg)
}
