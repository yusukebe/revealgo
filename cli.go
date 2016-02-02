package revealgo

import (
	"fmt"
	"os"
	"reflect"

	flags "github.com/jessevdk/go-flags"
)

type CLI struct {
}

type CLIOptions struct {
	Port int `short:"p" long:"port" description:"tcp port number of this server. defaults is 3000."`
}

func (cli *CLI) Run() {
	opts, args, err := parseOptions()
	if err != nil {
		fmt.Printf("error:%v\n", err)
		os.Exit(1)
	}
	if len(args) < 1 {
		showHelp()
		os.Exit(0)
	}
	server := Server{
		port:   opts.Port,
	}
	server.Serve( args[0] )
}

func showHelp() {
	fmt.Fprint( os.Stderr, `Usage: revealgo [options] [FILE]

Options:
`)

	t := reflect.TypeOf(CLIOptions{})
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		var o string
		if s := tag.Get("short"); s != "" {
			o = fmt.Sprintf("-%s, --%s", tag.Get("short"), tag.Get("long"))
		} else {
			o = fmt.Sprintf("-%s", tag.Get("long"))
		}
		fmt.Fprintf(	os.Stderr, "  %-21s %s\n", o, tag.Get("description") )
	}
}

func parseOptions() (*CLIOptions, []string, error) {
	opts := &CLIOptions{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.Parse()
	if err != nil {
		return nil, nil, err
	}
	return opts, args, nil
}
