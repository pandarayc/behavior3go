package behavior3go

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	a := make(map[string]interface{})
	a["k"] = nil
	x, ok := a["k"]
	if ok {
		fmt.Println(x)
	}
	t.Fatalf("fail")

}
