package actions

import (
	_ "log/slog"
	"path"
	"strings"

	"github.com/fluent-qa/qgops/internal/utils/shell"
	"github.com/spf13/cobra"
)

var starterActions = shell.LoadCommands(path.Join(WorkSpaceDir, "starters.json"))

var (
	AvailableStarter = starterActions.GetAvailableActionNames()
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
			starterActions.ExecuteActionByName(category)
		},
	}
)
