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

func (bt *BehaviorTree) Load() {

}

func (bt *BehaviorTree) Dump() {

}

func (bt *BehaviorTree) Tick(target interface{}, blackboard *BlackBoard) NodeStatus {
	return 0
}
