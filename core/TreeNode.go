package core

import "github.com/pandarayc/behavior3go/config"

type ITreeNode interface {
	INode
	GetChild() INode
	SetChild(child INode)
}

type TreeNode struct {
	BaseNode
	BaseWorker
	child INode
}

func (t *TreeNode) Ctor() {
	t.category = CATEGORY_TREE_NODE
}

func (t *TreeNode) Initialize(cfg *config.NodeCfg) {
	t.BaseNode.Initialize(cfg)
}

func (node *TreeNode) GetChild() INode {
	return node.child
}

func (node *TreeNode) SetChild(child INode) {
	node.child = child
}
