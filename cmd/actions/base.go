package actions

import (
	"github.com/fluent-qa/qgops/internal/utils/shell"
	"github.com/spf13/cobra"
	"path"
	"strings"
)

func CreateActions(actionCategory, actionConfigFile string) ([]string, *cobra.Command) {
	usedActions := shell.LoadCommands(path.Join(WorkSpaceDir, actionConfigFile))

	availableActions := usedActions.GetAvailableActionNames()
	availableCmd := &cobra.Command{
		Use:       actionCategory,
		Short:     "fluent  " + actionCategory + " " + strings.Join(availableActions, ","),
		Long:      "fluent  " + actionCategory + " " + strings.Join(availableActions, ","),
		Example:   "fluent  " + actionCategory + " " + availableActions[0],
		ValidArgs: availableActions,
		//ValidArgsFunction: ;
		//RunE
		Run: func(cmd *cobra.Command, args []string) {
			category := args[0]
			usedActions.ExecuteActionByName(category)
		},
	}
	return availableActions, availableCmd
}
