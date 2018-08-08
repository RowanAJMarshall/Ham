package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rowanajmarshall/basil/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Ham programming langauge!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
