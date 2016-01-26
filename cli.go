package revealgo

import (
	_ "fmt"
)

type CLI struct {
}

func (cli *CLI) Run() {
	server := Server{port: 3000}
	server.Serve()
}
