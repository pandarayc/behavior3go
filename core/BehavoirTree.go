package core

import (
	"github.com/pandarayc/behavior3go/config"
)

type IBTree interface {
	config.IBaseConfig
}

type BehaviorTree struct {
	*config.NodeCfg
	root  INode
	debug interface{}
}

var _ IBTree = &BehaviorTree{}


// 官方导入的是树级别的配置
func (bt *BehaviorTree) Load(data *config.TreeCfg, baseResigter, customNodes RegisterHandlers) {

	// 初始化树信息
	for id := range data.Nodes {

	}



}

func (bt *BehaviorTree) Dump() {

}

func (bt *BehaviorTree) Tick(target interface{}, blackboard *BlackBoard) NodeStatus {
	return 0
}


// 树类型也可以执行，判定为子树类型