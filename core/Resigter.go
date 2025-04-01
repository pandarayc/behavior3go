package core

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
func (h RegisterHandlers) Add(name string, node INode) {
	h.handlers[name] = reflect.TypeOf(node).Elem()
}

func (h RegisterHandlers) New(name string) (INode, error) {
	if node, ok := h.handlers[name]; ok {
		c := reflect.New(node).Interface()
		return c.(INode), nil
	}
	return nil, fmt.Errorf("node not found")
}
