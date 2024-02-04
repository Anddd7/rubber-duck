package main

import (
	"fmt"

	tui "github.com/Anddd7/rubber-duck/rubber-duck-tui"
)

// demo
func main() {
	// options := []string{"foo", "bar", "baz"}
	// model := NewStringSelection("Select:", options)
	// model.Start()
	// fmt.Println(model.GetSelected())

	type option struct {
		id   int
		name string
	}
	options := []option{
		{id: 10086, name: "foo"},
		{id: 20086, name: "bar"},
		{id: 30086, name: "baz"},
	}
	model := tui.NewSelection("Select:", options, func(o option) string { return o.name })
	model.SetSingleMode()
	model.SetAutoclear()
	model.Start()
	fmt.Println(model.GetSelected())
}
