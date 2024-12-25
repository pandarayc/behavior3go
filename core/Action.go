package core

import "github.com/genet9496/behavior3go/config"

type IAction interface {
	INode
}

type Action struct {
	BaseNode
	BaseWorker
}

var _ IAction = &Action{}

func (node *Action) Ctor() {
	node.category = CATEGORY_ACTION
}

func (node *Action) Initialize(cfg *config.NodeCfg) {
	node.BaseNode.Initialize(cfg)
}
