package behavior3go

import (
	"github.com/pandarayc/behavior3go/config"
)

type IBTree interface {
	config.IBaseConfig
	INode
	SetRoot(INode)
	GetRoot() INode
}

type BehaviorTree struct {
	BaseNode
	root  INode
	debug interface{}
}

var _ IBTree = &BehaviorTree{}

func NewBehaviorTree() *BehaviorTree {
	return &BehaviorTree{}
}

// 伪装用
func (bt *BehaviorTree) Initialize(cfg *config.NodeCfg) {
	bt.BaseNode.Initialize(cfg)
}

func (bt *BehaviorTree) Execute(tick *Tick) NodeStatus {
	return bt.root.Execute(tick)
}

func (bt *BehaviorTree) Ctor() {
	bt.category = CATEGORY_TREE
}

func (bt *BehaviorTree) Dump() {

}

func (bt *BehaviorTree) Tick(target interface{}, blackboard *BlackBoard) NodeStatus {
	return 0
}

func (bt *BehaviorTree) SetRoot(root INode) {
	bt.root = root
}

func (bt *BehaviorTree) GetRoot() INode {
	return bt.root
}
