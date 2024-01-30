package validators

import (
	"fmt"
	"testing"
)

type DemoModel struct {
	Name     string
	Page     int
	PageSize int
}

func TestVerify(t *testing.T) {
	DemoModelRules := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
	testDemoModel := DemoModel{
		Name:     "test",
		Page:     0,
		PageSize: 0,
	}
	err := Verify(testDemoModel, DemoModelRules)
	if err == nil {
		fmt.Println(err)
	}
}
