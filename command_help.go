package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	helpMessage := "Welcome to the Pokedex!\nUsage:\n"

	var usage string
	for _, value := range getCommands() {
		usage += fmt.Sprintf("\n%v: %v", value.name, value.description)
	}

	fmt.Println(helpMessage + usage)

	return nil
}
