package code

import (
	"fmt"
	"github.com/fluent-qa/qfluent-ops/internal/utils/shell"
	"os"
	"path"
	"testing"
)

func TestCreateStartProject(t *testing.T) {
	var CurrentDir, _ = os.Getwd()
	var actions = shell.LoadCommands(path.Join(CurrentDir, "starters.json"))
	fmt.Println(actions)
}
