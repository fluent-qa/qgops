package shell

import (
	"github.com/fluent-qa/qfluent-ops/internal/utils"
	"github.com/fluent-qa/qfluent-ops/internal/utils/jsonutil"
	"github.com/samber/lo"
	"io"
	"log"
	"log/slog"
	"os"
)

type ActionSet struct {
	Actions []NamedAction `json:"actions"`
}

type NamedAction struct {
	Name     string   `yaml:"name" mapstructure:"name" json:"name"`
	Desc     string   `yaml:"desc" mapstructure:"desc" json:"desc"`
	Commands []string `yaml:"commands" mapstructure:"commands" json:"commands"`
}

var NotFoundAction = &NamedAction{
	Name: "Not Found",
	Desc: "Action doesn't exist",
}

func ExecShellCommand(cmd string) (int, error) {
	return Exec(cmd).
		FilterScan(func(s string, writer io.Writer) {
			log.Println(s)
		}).Stdout()
}

func ExecShellCommands(jsonFile string, data any) {
	jsonBytes, err := os.ReadFile(jsonFile)
	if err != nil {
		slog.Error("error when read file", err)
		return
	}

	var commands = &ActionSet{}
	jsonutil.ToStruct(string(jsonBytes), &commands)
	for _, namedAction := range commands.Actions {
		lo.ForEach(namedAction.Commands, func(command string, index int) {
			output := command
			if data != nil {
				output = utils.RenderTemplate(command, data)
			}
			_, err := ExecShellCommand(output)
			if err != nil {
				slog.Error("execute command set failed,", err)
				return
			}
		})
	}
}

func LoadCommands(configPath string) *ActionSet {
	slog.Info("Current Action File is " + configPath)
	actions := &ActionSet{}
	jsonutil.ToStructFromFile(configPath, &actions)
	return actions
}

func (s *ActionSet) GetAvailableActionNames() []string {
	var result []string
	lo.ForEach(s.Actions, func(command NamedAction, index int) {
		result = append(result, command.Name)
	})
	return result
}

func (s *ActionSet) GetNamedActionByName(name string) NamedAction {
	result := lo.Filter(s.Actions, func(item NamedAction, index int) bool {
		return item.Name == name
	})
	if len(result) > 0 {
		return result[0]
	} else {
		return *NotFoundAction
	}
}

func (s *ActionSet) ExecuteActionByName(name string) {
	ExecAction(s.GetNamedActionByName(name))
}

func ExecAction(action NamedAction) {
	slog.Info("start to execute command: ", action.Name, action.Desc)
	lo.ForEach(action.Commands, func(item string, index int) {
		_, _ = ExecShellCommand(item)
	})
}
