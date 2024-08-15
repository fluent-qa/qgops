package main

import (
	"github.com/fluent-qa/qgops/cmd/base"
	"github.com/fluent-qa/qgops/cmd/code" // Add this import
)

func init() {
	base.CmdRoot.AddCommand(code.StarterCmd)
}
func main() {
	_ = base.CmdRoot.Execute()
}
