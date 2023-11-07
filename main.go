package main

import (
	"fmt"
	"os"
	"os/user"
	"waiig/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}
