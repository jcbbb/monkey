package main

import (
	"fmt"
	"os/user"

	"github.com/jcbbb/monkey/www"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	//repl.Start(os.Stdin, os.Stdout)
	www.Start(0)
}
