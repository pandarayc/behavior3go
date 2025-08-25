package behavior3go

import (
	"fmt"
	"reflect"
)

// 方法注册器
type RegisterHandlers struct {
	handlers map[string]reflect.Type
}

// NewRegisterHandlers 创建一个方法注册器
func NewRegisterHandlers() *RegisterHandlers {
	return &RegisterHandlers{
		handlers: make(map[string]reflect.Type),
	}
}

// Add 添加对应管理方法
func (h RegisterHandlers) Add(name string, node INode) error {
	if _, ok := h.handlers[name]; ok {
		return fmt.Errorf("node %s already registered", name)
	}
	h.handlers[name] = reflect.TypeOf(node).Elem()
	return nil
}

func (h RegisterHandlers) New(name string) (INode, error) {
	if node, ok := h.handlers[name]; ok {
		c := reflect.New(node).Interface()
		return c.(INode), nil
	}
	return nil, fmt.Errorf("node not found")
}
