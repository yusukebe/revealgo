package main

import (
	"os"

	"github.com/yusukebe/revealgo"
)

func main() {
	cli := revealgo.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	exitCode := cli.Run(os.Args[1:])
	os.Exit(exitCode)
}
