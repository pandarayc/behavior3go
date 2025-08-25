package behavior3go

import (
	"fmt"

	"github.com/pandarayc/behavior3go/config"
)

// 配置以项目级别维护

// B3 represents the behavior tree manager
type B3 struct {
	// Project configuration
	projectCfg *config.B3ProjectCfg

	// Node registries
	customNodes *CustomNodes
	trees       map[string]*BehaviorTree
	// 通用方法表
	defaultHandlers *RegisterHandlers
	// 自定义方法表
	customHandlers *RegisterHandlers
	root           IBTree
}

// New creates a new behavior tree manager
func New() *B3 {
	mgr := &B3{
		customNodes: &CustomNodes{},
		trees:       make(map[string]*BehaviorTree),
	}
	mgr.defaultHandlers = GetDefaultRegisterHandlers()
	mgr.customHandlers = NewRegisterHandlers()
	return mgr
}

func (mgr *B3) RegisterCustomFunc(name string, node INode) error {
	// 检查是否已经注册
	if err := mgr.customHandlers.Add(name, node); err != nil {
		return err
	}
	return nil
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
	nodes := make(map[string]INode)
	for _, treeCfg := range mgr.projectCfg.Trees {
		for _, nodeCfg := range treeCfg.Nodes {
			// 构建叶子节点
			var node INode

			name := nodeCfg.GetName()

			// uid命名通常为树节点
			if len(name) >= 36 {
				node = &TreeNode{}
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
			node.SetWorker(node.(IWorker))
			nodes[nodeCfg.GetId()] = node
		}
		// 构建树根节点
		treeNode := NewBehaviorTree()
		treeNode.Initialize(&treeCfg.NodeCfg)
		treeNode.Ctor()
		nodes[treeCfg.GetId()] = treeNode
	}

	// 对每棵树连接子节点
	for _, treeCfg := range mgr.projectCfg.Trees {
		for _, nodeCfg := range treeCfg.Nodes {
			node := nodes[nodeCfg.Id]
			switch node.GetCategory() {
			case CATEGORY_TREE_NODE:
				treeNode := node.(ITreeNode)
				if nodeCfg.Name != "" {
					treeNode.SetChild(nodes[nodeCfg.Name])
				}
			case CATEGORY_COMPOSITE:
				compositeNode := node.(IComposite)
				for _, childId := range nodeCfg.Children {
					childNode := nodes[childId]
					compositeNode.AddChild(childNode)
				}
			case CATEGORY_DECORATOR:
				if nodeCfg.Child != "" {
					decoratorNode := node.(IDecorator)
					decoratorNode.SetChild(nodes[nodeCfg.Child])
				}
			}
		}
		treeNode := nodes[treeCfg.Id].(IBTree)
		treeNode.SetRoot(nodes[treeCfg.Root])
	}

	mgr.root = nodes[mgr.projectCfg.SelectedTree].(IBTree)

	return nil
}

// Tick 执行树的遍历
func (mgr *B3) Tick(target interface{}, blackboard *BlackBoard) NodeStatus {
	if target == nil {
		fmt.Println("target not initialized")
		return STATUS_ERROR
	}
	if blackboard == nil {
		fmt.Println("blackboard not initialized")
		return STATUS_ERROR
	}
	if mgr.root == nil {
		fmt.Println("root not initialized")
		return STATUS_ERROR
	}

	tick := NewTick()
	tick.SetBlackBoard(blackboard)
	tick.SetTree(mgr.root)
	tick.SetTarget(target)

	status := mgr.root.Execute(tick)
	// status := core.STATUS_SUCCESS

	// 记录遍历过程
	return status
}

func (mgr *B3) DumpCfg() *config.B3ProjectCfg {
	return mgr.projectCfg
}

func (mgr *B3) DumpTree() {
	printNode(mgr.root, 0)
}

func printNode(root INode, blk int) {

	//fmt.Println("new node:", root.Name, " children:", len(root.Children), " child:", root.Child)
	for i := 0; i < blk; i++ {
		fmt.Print(" ") //缩进
	}

	//fmt.Println("|—<", root.Name, ">") //打印"|—<id>"形式
	fmt.Print("|—", root.GetTitle())

	if root.GetCategory() == CATEGORY_TREE {
		fmt.Println("")
		tree := root.(IBTree)
		printNode(tree.GetRoot(), blk+3)
		return
	}

	if root.GetCategory() == CATEGORY_DECORATOR {
		dec := root.(IDecorator)
		if dec.GetChild() != nil {
			//fmt.Print("=>")
			printNode(dec.GetChild(), blk+3)
			return
		}
	}

	fmt.Println("")
	if root.GetCategory() == CATEGORY_COMPOSITE {
		comp := root.(IComposite)
		if comp.GetChildrenCount() > 0 {
			for i := 0; i < comp.GetChildrenCount(); i++ {
				printNode(comp.GetChild(i), blk+3)
			}
		}
		return
	}
}
