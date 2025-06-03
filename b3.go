package behavior3go

import (
	"fmt"

	"github.com/pandarayc/behavior3go/config"
	"github.com/pandarayc/behavior3go/core"
)

// 配置以项目级别维护

// B3 represents the behavior tree manager
type B3 struct {
	// Project configuration
	projectCfg *config.B3ProjectCfg

	// Node registries
	customNodes *core.CustomNodes
	trees       map[string]*core.BehaviorTree
	// 通用方法表
	defaultHandlers *core.RegisterHandlers
	// 自定义方法表
	customHandlers *core.RegisterHandlers
	root           core.IBTree
	target         interface{}
	blackboard     *core.BlackBoard
}

// New creates a new behavior tree manager
func New() *B3 {
	mgr := &B3{
		customNodes: &core.CustomNodes{},
		trees:       make(map[string]*core.BehaviorTree),
	}
	mgr.defaultHandlers = GetDefaultRegisterHandlers()
	mgr.customHandlers = core.NewRegisterHandlers()
	mgr.blackboard = core.NewBlackBoard()
	mgr.target = nil
	return mgr
}

func (mgr *B3) RegisterCustomFunc(name string, node core.INode) {
}

// LoadProjectCfg loads a behavior tree project configuration from a file
func (mgr *B3) LoadCfg(path string) error {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return err
	}
	mgr.projectCfg = cfg
	return nil
}

// 打印配置
func (mgr *B3) EchoConfig() {
	fmt.Println("mgr.projectCfg:", mgr.projectCfg)
}

// 构建出项目树
func (mgr *B3) Load() error {
	if mgr.projectCfg == nil {
		return fmt.Errorf("project config not loaded")
	}

	// 节点实例化
	nodes := make(map[string]core.INode)
	for _, treeCfg := range mgr.projectCfg.Trees {
		for _, nodeCfg := range treeCfg.Nodes {
			// 构建叶子节点
			var node core.INode

			name := nodeCfg.GetName()

			// uid命名通常为树节点
			if len(name) >= 36 {
				node = &core.TreeNode{}
			} else {
				// 检查自定义表
				if tnode, err := mgr.customHandlers.New(name); err == nil {
					node = tnode
				} else {
					// 检查默认表
					if tnode, err := mgr.defaultHandlers.New(name); err == nil {
						node = tnode
					}
				}
			}
			if node == nil {
				return fmt.Errorf("node not found: %s", name)
			}
			node.Initialize(nodeCfg)
			node.Ctor()
			node.SetWorker(node.(core.IWorker))
			nodes[nodeCfg.GetId()] = node
		}
		// 构建树根节点
		treeNode := core.NewBehaviorTree()
		treeNode.Initialize(&treeCfg.NodeCfg)
		treeNode.Ctor()
		nodes[treeCfg.GetId()] = treeNode
	}

	// 对每棵树连接子节点
	for _, treeCfg := range mgr.projectCfg.Trees {
		for _, nodeCfg := range treeCfg.Nodes {
			node := nodes[nodeCfg.Id]
			switch node.GetCategory() {
			case core.CATEGORY_TREE_NODE:
				treeNode := node.(core.ITreeNode)
				if nodeCfg.Name != "" {
					treeNode.SetChild(nodes[nodeCfg.Name])
				}
			case core.CATEGORY_COMPOSITE:
				compositeNode := node.(core.IComposite)
				for _, childId := range nodeCfg.Children {
					childNode := nodes[childId]
					compositeNode.AddChild(childNode)
				}
			case core.CATEGORY_DECORATOR:
				if nodeCfg.Child != "" {
					decoratorNode := node.(core.IDecorator)
					decoratorNode.SetChild(nodes[nodeCfg.Child])
				}
			}
		}
		treeNode := nodes[treeCfg.Id].(core.IBTree)
		treeNode.SetRoot(nodes[treeCfg.Root])
	}

	mgr.root = nodes[mgr.projectCfg.SelectedTree].(core.IBTree)

	return nil
}

func (mgr *B3) Tick() core.NodeStatus {
	if mgr.blackboard == nil {
		fmt.Println("blackboard not initialized")
		return core.STATUS_ERROR
	}
	if mgr.root == nil {
		return core.STATUS_ERROR
	}
	// status := mgr.root.Execute(mgr.blackboard)
	status := core.STATUS_SUCCESS
	return status
}

func (mgr *B3) DumpCfg() *config.B3ProjectCfg {
	return mgr.projectCfg
}

func (mgr *B3) DumpTree() {
	printNode(mgr.root, 0)
}

func printNode(root core.INode, blk int) {

	//fmt.Println("new node:", root.Name, " children:", len(root.Children), " child:", root.Child)
	for i := 0; i < blk; i++ {
		fmt.Print(" ") //缩进
	}

	//fmt.Println("|—<", root.Name, ">") //打印"|—<id>"形式
	fmt.Print("|—", root.GetTitle())

	if root.GetCategory() == core.CATEGORY_TREE {
		fmt.Println("")
		tree := root.(core.IBTree)
		printNode(tree.GetRoot(), blk+3)
		return
	}

	if root.GetCategory() == core.CATEGORY_DECORATOR {
		dec := root.(core.IDecorator)
		if dec.GetChild() != nil {
			//fmt.Print("=>")
			printNode(dec.GetChild(), blk+3)
			return
		}
	}

	fmt.Println("")
	if root.GetCategory() == core.CATEGORY_COMPOSITE {
		comp := root.(core.IComposite)
		if comp.GetChildrenCount() > 0 {
			for i := 0; i < comp.GetChildrenCount(); i++ {
				printNode(comp.GetChild(i), blk+3)
			}
		}
		return
	}

}
