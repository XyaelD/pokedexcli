package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanUserInput(command string) []string {
	output := strings.ToLower(command)
	words := strings.Fields(output)
	return words
}

func runPokedex(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	userCommands := userCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		commandWords := cleanUserInput(scanner.Text())
		if len(commandWords) == 0 {
			continue
		}

		commandChoice := commandWords[0]
		var commandParam string
		if len(commandWords) > 1 {
			commandParam = commandWords[1]
		}

		c, ok := userCommands[commandChoice]
		if ok {
			err := c.callback(config, commandParam)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("The command %v does not exist. Type 'help' for a list of commands.\n", commandChoice)
		}
	}
}
