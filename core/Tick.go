package core

type Tick struct {
	Tree       *BehaviorTree
	Debug      interface{}
	Target     interface{}
	BlackBoard *BlackBoard
	openTree   []INode
	nodeCount  int32
}

// EnterNode 进入节点
func (t *Tick) EnterNode(node INode) {
	// 计算树节点列表
	t.nodeCount++
	t.openTree = append(t.openTree, node)
}

func (t *Tick) OpenNode(node INode) {
}

func (t *Tick) TickNode(node INode) {
}

func (t *Tick) CloseNode(node INode) {
	// 弹出节点
	t.nodeCount--
	t.openTree = t.openTree[:t.nodeCount]
}

func (t *Tick)ExitNode(node INode) {

}
