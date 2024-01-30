package yamlutil

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"testing"
)

type Person struct {
	Name   string   `mapstructure:"name" json:"name" yaml:"name"`
	Age    int      `mapstructure:"age" json:"age" yaml:"age"`
	Emails []string `mapstructure:"email" json:"email" yaml:"email"`
}

type ClassMate struct {
	Persons []Person `mapstructure:"persons" json:"persons" yaml:"persons"`
}

func TestWriteStructToFile(t *testing.T) {
	p1 := &Person{
		Name:   "name",
		Age:    10,
		Emails: []string{"123@chaqin.com", "324@qq.coms"},
	}
	p2 := &Person{
		//Name:   "name2",
		//Age:    12,
		//Emails: []string{"123@chaqin.com", "324@qq.coms"},
	}

	//classMate := &ClassMate{
	//	Persons: []Person{*p1, *p2},
	//}
	// Marshal the struct to YAML
	WriteStructToFile(&p1, "output.yaml")
	readIt(&p2)

	fmt.Println(p2)
}

func TestReadFileToStruct(t *testing.T) {
	classMate := &ClassMate{}
	ToStructFromFile("output.yaml", classMate)
	fmt.Println(classMate)
}

func TestToStruct(t *testing.T) {
	yamlString, _ := os.ReadFile("output.yaml")
	fmt.Println(string(yamlString))
	classMate := &ClassMate{}
	readIt(&classMate) //why it failed?
	fmt.Println(classMate)
	err := yaml.Unmarshal(yamlString, &classMate)
	if err != nil {
		slog.Error("convert yaml str failed")
	}
	fmt.Println(classMate.Persons[0].Name)
}

func readIt(data any) {
	yamlStr, err := os.ReadFile("output.yaml")
	if err != nil {
		slog.Error("read yaml str failed")
	}
	err = yaml.Unmarshal(yamlStr, &data)
	if err != nil {
		slog.Error("convert yaml str failed")
	}
}
