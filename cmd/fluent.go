package main

import (
	"github.com/fluent-qa/qfluent-ops/cmd/base"
	"github.com/fluent-qa/qfluent-ops/cmd/code"
)

func init() {
	base.CmdRoot.AddCommand(code.StarterCmd)
}
func main() {
	_ = base.CmdRoot.Execute()
}
