package starter

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/fluent-qa/qgops/internal/utils/shell"
)

func TestCreateStartProject(t *testing.T) {
	var CurrentDir, _ = os.Getwd()
	var actions = shell.LoadCommands(path.Join(CurrentDir, "starters.json"))
	fmt.Println(actions)
}
