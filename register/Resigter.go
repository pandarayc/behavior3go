package register

import "github.com/pandarayc/behavior3go/core"

var DefaultRegisterHandlers *RegisterHandlers

func GetDefaultRegisterHandlers() *RegisterHandlers {
	if DefaultRegisterHandlers == nil {
		DefaultRegisterHandlers = NewRegisterHandlers()
	}
	return DefaultRegisterHandlers
}

// 方法注册器
type RegisterHandlers struct {
	handlers map[string]core.INode
}

// NewRegisterHandlers 创建一个方法注册器
func NewRegisterHandlers() *RegisterHandlers {
	return &RegisterHandlers{
		handlers: make(map[string]core.INode),
	}
}

// Add 添加对应管理方法
func (h RegisterHandlers) Add(name string, node core.INode) {
	h.handlers[name] = node
}
