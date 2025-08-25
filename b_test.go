package behavior3go

import (
	"fmt"
	"testing"

)

func TestXxx(t *testing.T) {
	mgr := New()
	mgr.LoadCfg("example/project2.json")
	// mgr.EchoConfig()
	err := mgr.Load()
	if err != nil {
		t.Fatal(err)
	}
	// mgr.DumpTree()

	for i := 0; i < 10; i++ {
		fmt.Println("tick start ", i)
		mgr.Tick(nil, nil)
		fmt.Println("tick end ", i)
	}
}

func TestNewNode(t *testing.T) {
	name := "Wait"
	handlers := GetDefaultRegisterHandlers()
	node, err := handlers.New(name)
	if err != nil {
		t.Fatal(err)
	}
	node.SetWorker(node.(IWorker))
	_ = node
}
