package behavior3go

import "github.com/pandarayc/behavior3go/config"

type IDecorator interface {
	INode
	GetChild() INode
	SetChild(child INode)
}

type Decorator struct {
	BaseNode
	BaseWorker
	child INode
}

var _ IDecorator = &Decorator{}

func (node *Decorator) Ctor() {
	node.category = CATEGORY_DECORATOR
}

func (node *Decorator) Initialize(cfg *config.NodeCfg) {
	node.BaseNode.Initialize(cfg)
}

func (node *Decorator) GetChild() INode {
	return node.child
}

func (node *Decorator) SetChild(child INode) {
	node.child = child
}
