package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	// TODO make a command struct. (name, arg count, etc)
	commands = []string{"lock", "unlock", "exit"}
)

// Output a help message to stdout.
func printHelp() {
	fmt.Printf("Valid commands: %s\n", commands)
}

// Run the CLI in a loop until the user exits.
func runCLI() {
	var input, cmd string
	var tokens []string
	var err error
	var reader *bufio.Reader

	reader = bufio.NewReader(os.Stdin)

	for {
		// Read from stdin and tokenize the strings.
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		tokens = strings.Fields(input)

		// Filter out invalid commands.
		if len(tokens) < 1 {
			printHelp()
			continue
		}

		cmd = tokens[0]
		fmt.Printf("cmd: %s, args: %s\n", cmd, tokens[1:])

	}
}

func main() {
	runCLI()
}
