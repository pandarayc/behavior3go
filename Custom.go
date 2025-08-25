package behavior3go

// 自定义结构
type CustomNodes struct {
	store map[string]INode
}

// 注册处理方法
func (m *CustomNodes) Register(name string, handler INode) {
	m.store[name] = handler
}

// 读取自定义方法
func (m *CustomNodes) Load(name string) INode {
	return m.store[name]
}
