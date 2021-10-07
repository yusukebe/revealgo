package revealgo

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/crypto/bcrypt"

	flags "github.com/jessevdk/go-flags"
)

type CLI struct {
}

type CLIOptions struct {
	Help       bool   `short:"h" long:"help" description:"show help message"`
	Port       int    `short:"p" long:"port" description:"tcp port number of this server. default is 3000."`
	Theme      string `long:"theme" description:"slide theme or original css file name. default themes: beige, black, blood, league, moon, night, serif, simple, sky, solarized, and white" default:"black.css"`
	Transition string `long:"transition" description:"transition effect for slides: default, cube, page, concave, zoom, linear, fade, none" default:"default"`
	Multiplex  bool   `long:"multiplex" description:"enable slide multiplexing"`
	Version    bool   `short:"v" long:"version" description:"show the version"`
}

func (cli *CLI) Run() {
	opts, args, err := parseOptions()
	if err != nil {
		fmt.Printf("error:%v\n", err)
		os.Exit(1)
	}

	if opts.Version {
		showVersion()
		os.Exit(0)
	}

	if len(args) < 1 || opts.Help {
		showHelp()
		os.Exit(0)
	}

	_, err = os.Stat(opts.Theme)
	originalTheme := false
	if err == nil {
		originalTheme = true
	}

	server := Server{
		port: opts.Port,
	}
	param := ServerParam{
		Path:          args[0],
		Theme:         addExtention(opts.Theme, "css"),
		Transition:    opts.Transition,
		OriginalTheme: originalTheme,
	}
	if opts.Multiplex {
		password := "this-is-not-a-secret"
		bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		param.Multiplex = MultiplexParam{
			Secret:     password,
			Identifier: string(bytes),
		}
	}
	server.Serve(param)
}

func showHelp() {
	fmt.Fprint(os.Stderr, `Usage: revealgo [options] [MARKDOWN FILE]

Options:
`)
	t := reflect.TypeOf(CLIOptions{})
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		var o string
		if s := tag.Get("short"); s != "" {
			o = fmt.Sprintf("-%s, --%s", tag.Get("short"), tag.Get("long"))
		} else {
			o = fmt.Sprintf("--%s", tag.Get("long"))
		}
		fmt.Fprintf(os.Stderr, "  %-21s %s\n", o, tag.Get("description"))
	}
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "revealgo version %s\n", Version)
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

func addExtention(path string, ext string) string {
	if strings.HasSuffix(path, fmt.Sprintf(".%s", ext)) {
		return path
	}
	path = fmt.Sprintf("%s.%s", path, ext)
	return path
}
