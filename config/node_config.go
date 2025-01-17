package config

import (
	"encoding/json"
	"fmt"
)

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

func (t *NodeCfg) getProperties(name string) interface{} {
	v, ok := t.Properties[name]
	if !ok {
		panic(fmt.Sprintf("nodeCfg:%+v no prop:%+v", t, name))
	}
	return v
}

func (t *NodeCfg) GetPropertiesAsInt32(name string) int32 {
	v := t.getProperties(name)
	val, ok := v.(int32)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail", t.Name, name))
	}
	return val
}

func (t *NodeCfg) GetPropertiesAsInt64(name string) int64 {
	v := t.getProperties(name)
	val, ok := v.(int64)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail", t.Name, name))
	}
	return val
}

func (t *NodeCfg) GetPropertiesAsFloat32(name string) float32 {
	v := t.getProperties(name)
	val, ok := v.(float32)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail", t.Name, name))
	}
	return val
}

func (t *NodeCfg) GetPropertiesAsFloat64(name string) float64 {
	v := t.getProperties(name)
	val, ok := v.(float64)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail", t.Name, name))
	}
	return val
}

func (t *NodeCfg) GetPropertiesAsBool(name string) bool {
	v := t.getProperties(name)
	val, ok := v.(bool)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail", t.Name, name))
	}
	return val
}

func (t *NodeCfg) GetPropertiesAsString(name string) string {
	v := t.getProperties(name)
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail", t.Name, name))
	}
	return val
}
