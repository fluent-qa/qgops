package actions

import (
	_ "log/slog"
	"os"
	"path"
	"strings"

	"github.com/fluent-qa/qgops/internal/utils/shell"
	"github.com/spf13/cobra"
)

var CurrentDir = os.Getenv("FLUENT_HOME")

var actions = shell.LoadCommands(path.Join(CurrentDir, "starters.json"))

var (
	AvailableStarter = actions.GetAvailableActionNames()
	StarterCmd       = &cobra.Command{
		Use:       "starter",
		Short:     "fluent start " + strings.Join(AvailableStarter, ","),
		Long:      "fluent start " + strings.Join(AvailableStarter, ","),
		Example:   "fluent start " + AvailableStarter[0],
		ValidArgs: AvailableStarter,
		//ValidArgsFunction: ;
		//RunE
		Run: func(cmd *cobra.Command, args []string) {
			category := args[0]
			actions.ExecuteActionByName(category)
		},
	}
)
