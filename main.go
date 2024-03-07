package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/zkpranav/imonkey/repl"
)

const version = "0.0.1"

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to Monkey v%s %s!\n", version, usr.Username)
	repl.Start(os.Stdin, os.Stdout)
}