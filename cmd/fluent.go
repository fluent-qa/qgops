package main

import (
	"github.com/fluent-qa/qgops/cmd/actions"
	"github.com/fluent-qa/qgops/cmd/base"
)

func init() {
	base.CmdRoot.AddCommand(actions.StarterCmd)
	base.CmdRoot.AddCommand(actions.IssueCmd)
}
func main() {
	_ = base.CmdRoot.Execute()
}
