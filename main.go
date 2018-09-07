package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/marahin/yaaibig/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! This is YAAIBIG.\n", user.Username)
	fmt.Printf("Feel free to type\n")

	repl.Start(os.Stdin, os.Stdout)
}