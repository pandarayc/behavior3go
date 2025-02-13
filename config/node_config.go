package config

import (
	"encoding/json"
	"fmt"
)

type IBaseConfig interface {
	GetId() string
	GetTitle() string
	GetDescription() string
	GetProperties() map[string]interface{}
	GetProperty(name string) interface{}
	GetPropertyAsString(name string) string
	GetPropertyAsBool(name string) bool
	GetPropertyAsInt32(name string) int32
	GetPropertyAsInt64(name string) int64
	GetPropertyAsFloat32(name string) float32
	GetPropertyAsFloat64(name string) float64
}

// 节点配置
type NodeCfg struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
	// Display     map[string]int32       `json:"display"`
}

// 方便打印
func (t *NodeCfg) String() string {
	str, _ := json.Marshal(t)
	return string(str)
}

func (config *NodeCfg) GetId() string {
	return config.Id
}

func (node *NodeCfg) GetName() string {
	return node.Name
}

func (node *NodeCfg) GetTitle() string {
	return node.Title
}

func (node *NodeCfg) GetDescription() string {
	return node.Description
}

func (node *NodeCfg) GetProperties() map[string]interface{} {
	return node.Properties
}

func (t *NodeCfg) GetProperty(name string) interface{} {
	v, ok := t.GetProperties()[name]
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v not exist", t.GetName(), name))
	}
	return v
}

func (t *NodeCfg) GetPropertyAsInt(name string) int {
	v := t.GetProperty(name)
	val, ok := v.(int)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate", t.GetName(), name))
	}
	return val
}

func (t *NodeCfg) GetPropertyAsInt32(name string) int32 {
	v := t.GetProperty(name)
	val, ok := v.(int32)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate", t.GetName(), name))
	}
	return val
}

func (t *NodeCfg) GetPropertyAsInt64(name string) int64 {
	v := t.GetProperty(name)
	val, ok := v.(int64)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate", t.GetName(), name))
	}
	return val
}

func (t *NodeCfg) GetPropertyAsFloat32(name string) float32 {
	v := t.GetProperty(name)
	val, ok := v.(float32)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate", t.GetName(), name))
	}
	return val
}

func (t *NodeCfg) GetPropertyAsFloat64(name string) float64 {
	v := t.GetProperty(name)
	val, ok := v.(float64)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate", t.GetName(), name))
	}
	return val
}

func (t *NodeCfg) GetPropertyAsBool(name string) bool {
	v := t.GetProperty(name)
	val, ok := v.(bool)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate", t.GetName(), name))
	}
	return val
}

func (t *NodeCfg) GetPropertyAsString(name string) string {
	v := t.GetProperty(name)
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate", t.GetName(), name))
	}
	return val
}
