package core

type TreeMemory struct {
	nodeMemory     memory // 节点数据
	openNodes      []INode // 节点数据
	traversalDepth int32
	traversalCycle int32
}

// 数据黑板
type BlackBoard struct {
	baseMemory memory
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

func (b *BlackBoard) GetMemory(treeId, nodeId string) memory {
	mem := b.baseMemory
	if treeId != "" {
		treeMem := b.getTreeMemory(treeId)
		mem = treeMem.nodeMemory
	}

}

func (b *BlackBoard) getNodeMemory(mem memory, nodeId string) memory {

	return nil
}

// 获取树节点
func (b *BlackBoard) getTreeMemory(treeId string) *TreeMemory {
	if b.treeMemory[treeId] == nil {
		b.treeMemory[treeId] = newTreeMemory()
	}

	return b.treeMemory[treeId]
}

// 插入数据
func (b *BlackBoard) Set(key string, value interface{}, treeScope, nodeScope string) {
	return
}

// 获取数据
func (b *BlackBoard) Get(key, treeScope, nodeScope string) memory {
	mem := b.getMemory(treeScope, nodeScope)
	return mem[key]
}

// 树状数据
func newTreeMemory() *TreeMemory {
	return &TreeMemory{}
}

// ---------------------------
//  Memory 取值
// ---------------------------

type memory map[string]interface{}

func newMemory() memory {
	mem := make(map[string]interface{})
	return mem
}

func (m *memory) GetString() string {
	return ""
}

func (m *memory) GetBool() bool {
	return false
}

func (m *memory) GetInt64() int64 {
	return 0
}

func (m *memory) GetInt32() int32 {
	return 0
}
