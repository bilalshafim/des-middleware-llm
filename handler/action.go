package handler

import (
	"fmt"
)

type ActionFunc func()

func actionOne() {
	fmt.Println("Action One executed")
}

func actionTwo() {
	fmt.Println("Action Two executed")
}

var actions = map[string]ActionFunc{
	"actionOne": actionOne,
	"actionTwo": actionTwo,
}

func PerformAction(action *string) {
	if method, exists := actions[*action]; exists {
		method()
	} else {
		fmt.Println("Action not found.")
	}
}
