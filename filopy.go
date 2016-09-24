package main

import (
	"bufio"
	"bytes"
	"errors"
	//"encoding/base64"
	//"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
TODO:
cmds:
	arglen check.

users:
	add users? maybe not. is this a singly-user application?

Nodes:
	-figure out a better for Nodes than just defaulting to 7.
	-make node directories permissions only allow for write by
	the user that initialized filopy.

*/

type (
	Filopy struct {
		TrackedFiles []File
		Nodes        []Node
	}

	Command struct {
		name string
		argc int
		help string
	}
)

var (
	cmdInit    = Command{"init", 0, "Initialize Filopy in this directory"}
	cmdDestroy = Command{"destroy", 0, "Destroy this Filopy space"}
	cmdAdd     = Command{"add", -1, "Begin tracking the given files"}
	cmdRm      = Command{"rm", -1, "Stop tracking the given files"}
	cmdFiles   = Command{"files", 0, "List the files tracked by Filopy"}
	cmdExit    = Command{"exit", 0, "Exit the Filopy CLI"}
	cmdLock    = Command{"lock", -1, "Lock the given files"}
	cmdUnlock  = Command{"unlock", -1, "Unlock the given files"}
	cmdHelp    = Command{"help", 1, "Show commands"}

	commands = []Command{
		cmdInit, cmdDestroy, cmdAdd, cmdRm, cmdFiles, cmdExit, cmdLock,
		cmdUnlock, cmdHelp,
	}
	preInitCommands = []Command{
		cmdInit, cmdHelp, cmdExit,
	}

	errAlreadyInit = errors.New("Filopy has already been initialized.")
	errNeedInit    = errors.New("Filopy hasn't been initialized.")
	errUnsupported = errors.New("Unsupported command. Use \"help\" to" +
		" view supported commands.")
)

const (
	numNodes     = 7
	defaultPerms = 0777
)

// Generate a UUID.
func generateUUID() (string, error) {
	var reader *bufio.Reader
	var file *os.File
	var buf []byte
	var out string
	var err error

	file, err = os.Open("/dev/urandom")
	if err != nil {
		return "", err
	}

	buf = make([]byte, 16)
	reader = bufio.NewReader(file)
	reader.Read(buf)

	out = fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8],
		buf[8:10], buf[10:16],
	)
	return out, nil
}

// Return whether the given command is supported.
func isValidCmd(name string) bool {
	for _, c := range commands {
		if name == c.name {
			return true
		}
	}
	return false
}

// Initialize Filopy in this directory.
// TODO make the user set up a passphrase
func initialize() (*Filopy, error) {
	if err := os.Mkdir(".filopy", defaultPerms); err != nil {
		return nil, err
	}

	var filopy *Filopy
	filopy = &Filopy{
		TrackedFiles: make([]File, 0),
		Nodes:        make([]Node, 0),
	}

	// Create the Nodes.
	var name string
	for i := 0; i < numNodes; i++ {
		name = fmt.Sprintf(".filopy/%d", i)
		if err := os.Mkdir(name, defaultPerms); err != nil {
			return nil, err
		}
	}

	return filopy, nil
}

// Begin tracking the given files.
func (f *Filopy) add() error {
	fmt.Println("running add")
	return nil
}

// Stop tracking the given files.
func (f *Filopy) rm() error {
	return nil
}

// TODO
func (f *Filopy) destroy() error {
	return nil
}

// Output a help message to stdout.
// TODO fix the spacing here
func (f *Filopy) help() {
	var name, help string
	var argc int

	fmt.Printf("    Command  Args  Help\n")
	for _, c := range commands {
		name = c.name
		argc = c.argc
		help = c.help
		if len(name) > 8 {
			name = name[:8]
		}

		fmt.Printf("    %s  %d  %s\n", name, argc, help)
	}
}

// Lock the given files
func (f *Filopy) lock() error {
	return nil
}

// Unlock the given files
func (f *Filopy) unlock() error {
	return nil
}

// Leave the Filopy CLI.
func (f *Filopy) exit() {
	os.Exit(0)
}

// Output the files tracked by Filopy.
// TODO make the formatting thing more generic.
func (f *Filopy) files() error {
	var name string
	var locked bool
	for _, file := range f.TrackedFiles {
		name = file.AbsolutePath
		locked = file.Locked
		if len(file.AbsolutePath) > 15 {
			name = file.AbsolutePath[:15]
		}
		fmt.Printf("File: %s  Locked: %t\n", name, locked)
	}
	return nil
}

// Returns whether the input command is allowed to run before initializing
// filopy.
func inPreInit(cmd string) bool {
	for _, c := range preInitCommands {
		if c.name == cmd {
			return true
		}
	}
	return false
}

// TODO: improve the serialization
func serialize(f *Filopy) error {
	buf := new(bytes.Buffer)
	lengths := fmt.Sprintf("%s%s", len(f.TrackedFiles), len(f.Nodes))
	if _, err := buf.WriteString(lengths); err != nil {
		return err
	}

	for _, f := range f.TrackedFiles {
		buf.WriteString(f.String() + "$")
	}

	// Write to file
	// TODO

	return nil
}

func deserialize() *Filopy {
	return nil
}

// Run the CLI in a loop until the user exits.
func run() {
	var filopy *Filopy
	var reader *bufio.Reader
	var input, cmd string
	var tokens []string
	var err error

	filopy = deserialize()
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		if input == "\n" {
			continue
		}

		tokens = strings.Fields(input)
		cmd = tokens[0]

		if !isValidCmd(cmd) {
			fmt.Println(errUnsupported)
			continue
		}

		if filopy == nil {
			if !inPreInit(cmd) {
				fmt.Println(errNeedInit)
				continue
			}
		}

		switch cmd {
		case cmdInit.name:
			filopy, err = initialize()
			if err != nil {
				fmt.Println(errAlreadyInit)
			}
		case cmdDestroy.name:
			filopy.destroy()
		case cmdAdd.name:
			filopy.add()
		case cmdRm.name:
			filopy.rm()
		case cmdExit.name:
			filopy.exit()
		case cmdFiles.name:
			filopy.files()
		case cmdLock.name:
			filopy.lock()
		case cmdUnlock.name:
			filopy.unlock()
		case cmdHelp.name:
			filopy.help()
		}

		if err = serialize(filopy); err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	// TODO: change working directory to the right location
	// TODO: setup in /usr/var or something.
	run()
}
