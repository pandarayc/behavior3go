package config

type CustomNodeCfg struct {
	NodeCfg
}

func NewCustomNodeCfg() *CustomNodeCfg {
	return &CustomNodeCfg{
		NodeCfg: NodeCfg{},
	}
}

// Ensure CustomNodeCfg implements IBaseConfig
var _ IBaseConfig = &CustomNodeCfg{}
