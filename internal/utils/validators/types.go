package validators

import "go/types"

type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

type NamedRule struct {
	Name         string
	ValidateFunc types.Func
}
