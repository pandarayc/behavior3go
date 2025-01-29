package core

type IBlackBoard interface {
	GetMemory(treeScope string, nodeScope string) IMemory
}

var _ IBlackBoard = &BlackBoard{}

// 数据黑板
type BlackBoard struct {
	baseMemory *memory
	treeMemory map[string]*TreeMemory
}

// 新建黑板数据
func NewBlackBoard() *BlackBoard {
	blackBoard := &BlackBoard{
		baseMemory: newMemory(),
		treeMemory: make(map[string]*TreeMemory),
	}

	return blackBoard
}

// 获取树节点
func (b *BlackBoard) GetTreeMemory(treeId string) *TreeMemory {
	if b.treeMemory[treeId] == nil {
		b.treeMemory[treeId] = newTreeMemory()
	}

	return b.treeMemory[treeId]
}

// GetMemory 获取记录
func (b *BlackBoard) GetMemory(treeScope, nodeScope string) IMemory {
	mem := b.baseMemory
	if len(treeScope) == 0 {
		treeMem := b.GetTreeMemory(treeScope)
		mem = treeMem.memory
		if nodeScope != "" {
			mem = treeMem.GetNodeMemory(nodeScope)
		}
	}
	return mem
}

// 作为共享信息层处理
type TreeMemory struct {
	nodeMemory map[string]*memory // 节点数据
	*memory                       // 树节点数据
}

// 树状数据
func newTreeMemory() *TreeMemory {
	return &TreeMemory{
		nodeMemory: make(map[string]*memory),
		memory:     newMemory(),
	}
}
func (tm *TreeMemory) GetNodeMemory(key string) *memory {
	if tm == nil {
		return nil
	}
	return tm.nodeMemory[key]
}
