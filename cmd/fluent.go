package main

import (
	"github.com/fluent-qa/qgops/cmd/base"
	"github.com/fluent-qa/qgops/cmd/starter" // Add this import
)

func init() {
	base.CmdRoot.AddCommand(starter.StarterCmd)
}
func main() {
	_ = base.CmdRoot.Execute()
}
