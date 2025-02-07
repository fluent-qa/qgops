package main

import (
	"github.com/fluent-qa/qgops/app/fluentcli/actions"
	"github.com/fluent-qa/qgops/app/fluentcli/base"
)

func init() {
	base.CmdRoot.AddCommand(actions.StarterCmd)
	base.CmdRoot.AddCommand(actions.IssueCmd)
	base.CmdRoot.AddCommand(actions.UtilCmd)
	base.CmdRoot.AddCommand(actions.FetchAllCmd)
	base.CmdRoot.AddCommand(actions.AddURLCmd)
	base.CmdRoot.AddCommand(actions.FetchByCategoryCmd)
}
func main() {
	_ = base.CmdRoot.Execute()
}
