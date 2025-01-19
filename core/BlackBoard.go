package core

type IBlackBoard interface {
	GetMemory(key string, treeScope string, nodeScope string) IMemory
}

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

func (b *BlackBoard) GetMemory(treeScope, nodeScope string) *memory {
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

// func (b *BlackBoard) Get(key string, treeScope, nodeScope string) interface{} {
// 	mem := b.GetMemory(treeScope, nodeScope)
// 	if mem != nil {
// 		return mem.Get(key)
// 	}
// 	return nil
// }

// func (b *BlackBoard) Set(key string, value interface{}, treeScope, nodeScope string) {
// 	mem := b.GetMemory(treeScope, nodeScope)
// 	if mem != nil {
// 		mem.Set(key, value)
// 	}
// }

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
