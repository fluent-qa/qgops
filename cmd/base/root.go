package base

import (
	"github.com/spf13/cobra"
	"os"
)

var VERSION = "0.0.2"
var WorkSpaceDir = os.Getenv("FLUENT_HOME")

var CmdRoot = &cobra.Command{
	Use:     "fluent",
	Example: "fluent <category> <action> <--options>",
	Short:   "fluent is for feeds usage",
	Version: VERSION,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
