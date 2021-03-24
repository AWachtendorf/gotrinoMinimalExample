package main

import "github.com/golangee/i18n"

func main() {
	//processes the module to perform internationalization
	if err := i18n.Bundle(); err != nil {
		panic(err)
	}
}
