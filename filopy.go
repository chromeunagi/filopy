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
	cmdDestroy = Command{"destroy", 0}
	cmdExit    = Command{"exit", 0}
	cmdFiles   = Command{"files", 0}
	cmdInit    = Command{"init", 0}
	cmdLock    = Command{"lock", 1}
	cmdUnlock  = Command{"unlock", 1}

	commands = []Command{cmdDestroy, cmdExit, cmdFiles, cmdInit, cmdLock, cmdUnlock}
)

const (
	nodes        = 7
	defaultPerms = 0777
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

// Check if filopy has been initialized.
// TODO: implement the check. easy, just check if the .filopy file exists.
func isInitialized() bool {
	_, err := os.Stat(".filopy")
	return err == nil
}

// TODO: make the user setup a passphrase.
func executeInit() error {
	// Create the .filopy directory.
	if err := os.Mkdir(".filopy", defaultPerms); err != nil {
		return err
	}

	// Create the nodes.
	var name string
	for i := 0; i < nodes; i++ {
		name = fmt.Sprintf(".filopy/%d", i)
		if err := os.Mkdir(name, defaultPerms); err != nil {
			return err
		}
	}

	// TODO: set up passphrase
	// ...

	return nil
}

func executeExit() {
	os.Exit(0)
}

// TODO: get pass, 3 confirmations, etc.
func executeDestroy() {
	return
}

func executeFiles() error {
	// TODO
	return nil
}

// Run the CLI in a loop until the user exits.
func runCLI() {
	var input, cmd string
	var tokens []string
	var err error
	var reader *bufio.Reader

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")

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

		cmd = tokens[0]
		if !isInitialized() && cmd != "init" {
			fmt.Println("Use the \"init\" command to initialize filopy.")
			continue
		}

		switch cmd {
		case cmdInit.name:
			if err = executeInit(); err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				dir, err := os.Getwd()
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				fmt.Printf("Successfully initialized filopy to %s\n", dir)
			}
		case cmdDestroy.name:
			executeDestroy()
		case cmdExit.name:
			executeExit()
		case cmdFiles.name:
			// TODO
		case cmdLock.name:
			fmt.Println("execute lock")
		case cmdUnlock.name:
			fmt.Println("execute unlock")
		default:
			printHelp()
			continue
		}
	}
}

func main() {
	// TODO: change working directory to the right location
	// TODO: setup in /usr/var or something.
	runCLI()
}
