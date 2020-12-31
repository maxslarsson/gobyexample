package main

import (
	"fmt"
	"strings"

	"github.com/maxslarsson/gobyexample/constants"
	_for "github.com/maxslarsson/gobyexample/for"
	"github.com/maxslarsson/gobyexample/helloworld"
	"github.com/maxslarsson/gobyexample/ifelse"
	_switch "github.com/maxslarsson/gobyexample/switch"
	"github.com/maxslarsson/gobyexample/values"
	"github.com/maxslarsson/gobyexample/variables"

	"github.com/fatih/color"
)

func main() {
	functions := []fn{
		fn{"Hello World", helloworld.HelloWorld},
		fn{"Values", values.Values},
		fn{"Variables", variables.Variables},
		fn{"Constants", constants.Constants},
		fn{"For", _for.For},
		fn{"If/Else", ifelse.IfElse},
		fn{"Switch", _switch.Switch},
	}

	for i, function := range functions {
		color.Blue(function.name)
		fmt.Println(strings.Repeat("-", len(function.name)))
		fmt.Print("\n")
		function.fn()
		if i < len(functions)-1 {
			fmt.Print("\n\n\n")
		}
	}
}

type fn struct {
	name string
	fn   func()
}
