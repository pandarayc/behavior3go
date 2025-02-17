package core

// 方法注册器
type RegisterHandlers struct {
	handlers map[string]INode
}

// NewRegisterHandlers 创建一个方法注册器
func NewRegisterHandlers() *RegisterHandlers {
	return &RegisterHandlers{
		handlers: make(map[string]INode),
	}
}

// Add 添加对应管理方法
func (h RegisterHandlers) Add(name string, node INode) {
	h.handlers[name] = node
}
