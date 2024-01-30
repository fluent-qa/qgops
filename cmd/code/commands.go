package code

import (
	"github.com/fluent-qa/qfluent-ops/internal/utils/shell"
	"github.com/spf13/cobra"
	_ "log/slog"
	"os"
	"path"
	"strings"
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
		Run: func(cmd *cobra.Command, args []string) {
			category := args[0]
			actions.ExecuteActionByName(category)
		},
	}
)
