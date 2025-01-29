package core

type Tick struct {
	tree       *BehaviorTree
	debug      interface{}
	target     interface{}
	blackBoard *BlackBoard
	_openNodes []INode
	_nodeCount int32
}

func (t *Tick) GetTree() IBTree {
	return t.tree
}

func (t *Tick) GetBlackBoard() IBlackBoard {
	return t.blackBoard
}

// EnterNode 进入节点
func (t *Tick) _enterNode(node INode) {
	// 计算树节点列表
	t._nodeCount++
	t._openNodes = append(t._openNodes, node)
}

// 开启节点
func (t *Tick) _openNode(node INode) {
}

// tick 节点
func (t *Tick) _tickNode(node INode) {
}

// 关闭节点
func (t *Tick) _closeNode(node INode) {
	// 弹出节点
	// t._nodeCount--
	if len(t._openNodes) > 0 {
		t._openNodes = t._openNodes[:t._nodeCount-1]
	}
}

// 推出节点
func (t *Tick) _exitNode(node INode) {

}
