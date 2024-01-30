package shell

import (
	"fmt"
	"testing"
)

var s = LoadCommands("starters.json")

func TestExecShellCommands(t *testing.T) {
	commands := &ActionSet{}
	ExecShellCommands("starters.json", commands)
}

func TestLoadConfig(t *testing.T) {
	var result = LoadCommands("starters.json")
	fmt.Println(result)
}

func TestActionSet_GetAvailableActionNames(t *testing.T) {
	var names = s.GetAvailableActionNames()
	fmt.Println(names)
}

func TestActionSet_ExecuteActionByName(t *testing.T) {
	s.ExecuteActionByName("java")
}
