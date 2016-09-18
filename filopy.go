package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// TODO: possibly add a help message for each cmd?
type Command struct {
	name string
	argc int
}

var (
	cmdLock   = Command{"lock", 1}
	cmdUnlock = Command{"unlock", 1}
	cmdExit   = Command{"exit", 0}

	commands = []Command{cmdLock, cmdUnlock, cmdExit}
)

// Output a help message to stdout.
func printHelp() {
	for _, c := range commands {
		fmt.Printf("Command: %s, Args: %d\n", c.name, c.argc)
	}
}

// Returns whether the given string represents the name of a
// valid command.
func isValidCmd(name string) bool {
	for _, c := range commands {
		if name == c.name {
			return true
		}
	}
	return false
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

		// Filter out empty input.
		if len(tokens) < 1 {
			printHelp()
			continue
		}

		// Filter out invalid commands.
		// TODO: make this more specific. check for arglength, etc
		cmd = tokens[0]
		if !isValidCmd(cmd) {
			printHelp()
			continue
		}

		fmt.Printf("cmd: %s, args: %s\n", cmd, tokens[1:])

	}
}

func main() {
	runCLI()
}
