package main

import (
	"fmt"
	"hoku/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n",
		user.Username)
	fmt.Printf("Type in commands to get started\n")
	repl.Start(os.Stdin, os.Stdout)
}
