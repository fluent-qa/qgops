package cursor

import (
	"fmt"
	"testing"
)

func TestGenerator_GenerateMacMachineID(t *testing.T) {
	g := NewGenerator()
	id, _ := g.GenerateMacMachineID()
	fmt.Println(id)
}
