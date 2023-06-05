package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jcbbb/monkey/repl"
	"github.com/jcbbb/monkey/www"
)

func main() {
	replCmd := flag.NewFlagSet("repl", flag.ExitOnError)

	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	addr := serverCmd.String("addr", ":3000", "port to listen")

	if len(os.Args) < 2 {
		fmt.Println("expected either `repl` or `server` subcommand")
	}

	switch os.Args[1] {
	case "server":
		serverCmd.Parse(os.Args[2:])
		www.Start(*addr)
	case "repl":
		replCmd.Parse(os.Args[2:])
		repl.Start(os.Stdin, os.Stdout)
	default:
		fmt.Println("expected either `repl` or `server` subcommand")
		os.Exit(1)
	}
}
