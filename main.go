package main

import (
	"fmt"
	"strings"

	"github.com/maxslarsson/gobyexample/constants"

	"github.com/fatih/color"

	"github.com/maxslarsson/gobyexample/variables"

	"github.com/maxslarsson/gobyexample/helloworld"
	"github.com/maxslarsson/gobyexample/values"
)

func main() {
	functions := []fn{
		fn{"Hello World", helloworld.HelloWorld},
		fn{"Values", values.Values},
		fn{"Variables", variables.Variables},
		fn{"Constants", constants.Constants},
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
