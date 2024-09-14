package base

import (
	"github.com/spf13/cobra"
)

var VERSION = "0.0.2"

var CmdRoot = &cobra.Command{
	Use:     "fluent",
	Example: "fluent <category> <action> <--options>",
	Short:   "fluent cli is a collection of daily scripts",
	Version: VERSION,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
