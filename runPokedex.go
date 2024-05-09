package main

import (
	"bufio"
	"fmt"
	"os"
)

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	scanner.Scan()
	command := scanner.Text()
	return command
}

func runPokedex(config *config) {
	userCommands := userCommands()
	for {
		command := userInput()
		c, ok := userCommands[command]
		if ok {
			err := c.callback(config)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("The command %v does not exist. Type 'help' for a list of commands.\n", command)
		}
	}
}
