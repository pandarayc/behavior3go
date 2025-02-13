package core

// Tick 对象是用于每个Update用于传递上下文的参数对象
// 并记录当前节点的运行状态
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

// GetMemory 获取当前树对应节点的黑板信息
func (t *Tick) GetMemory(nodeId string) IMemory {
	return t.blackBoard.GetMemory(t.tree.Id, nodeId)
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
