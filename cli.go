package revealgo

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type CLI struct {
}

type CLIOptions struct {
	Port int `short:"p" long:"port" description:"tcp port number of this server. defaults is 3000."`
}

func (cli *CLI) Run() {
	opts, args, err := cli.parseOptions()
	if err != nil {
		fmt.Printf("error:%v\n", err)
		os.Exit(1)
	}
	if len(args) < 1 {
		cli.showHelp()
		os.Exit(1)
	}
	server := Server{
		port:   opts.Port,
		mdpath: args[0],
	}
	server.Serve()
}

func (cli *CLI) showHelp() {
	fmt.Println("Help...")
}

func (cli *CLI) parseOptions() (*CLIOptions, []string, error) {
	opts := &CLIOptions{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.Parse()
	if err != nil {
		return nil, nil, err
	}
	return opts, args, nil
}
