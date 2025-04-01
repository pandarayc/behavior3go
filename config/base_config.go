package config

import (
	"fmt"

	"github.com/pandarayc/behavior3go/utils"
	"github.com/spf13/cast"
)

type BaseConfig struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Properties  map[string]any `json:"properties"`
}

func (config *BaseConfig) GetId() string {
	if config.Id == "" {
		config.Id = utils.GenUuid() // 其实project级别已经全部关联好了，保底
	}
	return config.Id
}

func (node *BaseConfig) GetName() string {
	return node.Name
}

func (node *BaseConfig) GetTitle() string {
	return node.Title
}

func (node *BaseConfig) GetDescription() string {
	return node.Description
}

func (node *BaseConfig) GetProperties() map[string]any {
	return node.Properties
}

func (t *BaseConfig) GetProperty(name string) interface{} {
	v, ok := t.GetProperties()[name]
	if !ok {
		panic(fmt.Sprintf("nodeId:%+v prop:%v not exist", t.GetName(), name))
	}
	return v
}

func (t *BaseConfig) GetPropertyAsInt(name string) int {
	v := t.GetProperty(name)
	val, err := cast.ToIntE(v)
	if err != nil {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate error:%v", t.GetName(), name, err))
	}
	return val
}

func (t *BaseConfig) GetPropertyAsInt32(name string) int32 {
	v := t.GetProperty(name)
	val, err := cast.ToInt32E(v)
	if err != nil {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate error:%v", t.GetName(), name, err))
	}
	return val
}

func (t *BaseConfig) GetPropertyAsInt64(name string) int64 {
	v := t.GetProperty(name)
	val, err := cast.ToInt64E(v)
	if err != nil {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate error:%v", t.GetName(), name, err))
	}
	return val
}

func (t *BaseConfig) GetPropertyAsFloat32(name string) float32 {
	v := t.GetProperty(name)
	val, err := cast.ToFloat32E(v)
	if err != nil {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate error:%v", t.GetName(), name, err))
	}
	return val
}

func (t *BaseConfig) GetPropertyAsFloat64(name string) float64 {
	v := t.GetProperty(name)
	val, err := cast.ToFloat64E(v)
	if err != nil {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate error:%v", t.GetName(), name, err))
	}
	return val
}

func (t *BaseConfig) GetPropertyAsBool(name string) bool {
	v := t.GetProperty(name)
	val, err := cast.ToBoolE(v)
	if err != nil {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate error:%v", t.GetName(), name, err))
	}
	return val
}

func (t *BaseConfig) GetPropertyAsString(name string) string {
	v := t.GetProperty(name)
	val, err := cast.ToStringE(v)
	if err != nil {
		panic(fmt.Sprintf("nodeId:%+v prop:%v fail translate error:%v", t.GetName(), name, err))
	}
	return val
}
