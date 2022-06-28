package main

import (
	"fmt"
	"github.com/tzcl/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err) // Couldn't get a user
	}
	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("This repl will tokenise any input you type in\n")
	repl.Start(os.Stdin, os.Stdout)
}
