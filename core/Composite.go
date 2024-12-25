package core

import "github.com/genet9496/behavior3go/config"

type IComposite interface {
	INode
	GetChildren() []INode
	AddChild(child INode)
	GetChild(index int) INode
	PopChild() INode
	GetChildrenCount() int
}

type Composite struct {
	BaseNode
	BaseWorker
	children []INode
}

var _ IComposite = &Composite{}

func (node *Composite) _ctor() {
	node.category = CATEGORY_COMPOSITE
}

func (node *Composite) Initialize(cfg *config.NodeCfg) {
	node.BaseNode.Initialize(cfg)
	node.children = make([]INode, 0)
}

func (node *Composite) GetChildren() []INode {
	return node.children
}

func (node *Composite) AddChild(child INode) {
	node.children = append(node.children, child)
}

func (node *Composite) GetChild(index int) INode {
	return node.children[index]
}

func (node *Composite) PopChild() INode {
	if len(node.children) == 0 {
		return nil
	}
	child := node.children[len(node.children)-1]
	node.children = node.children[:len(node.children)-1]

	return child
}

func (node *Composite) GetChildrenCount() int {
	return len(node.children)
}
