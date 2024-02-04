package revealgo

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/bcrypt"

	flags "github.com/jessevdk/go-flags"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

type CLI struct {
	OutStream, ErrStream io.Writer
	options              *CLIOptions
	parser               *flags.Parser
}

type CLIOptions struct {
	Port              int    `short:"p" long:"port" description:"TCP port number of this server" default:"3000"`
	Theme             string `long:"theme" description:"Slide theme or original css file name. default themes: beige, black, blood, league, moon, night, serif, simple, sky, solarized, and white" default:"black.css"`
	Template		  string `long:"template" description:"Custom HTML template file name. default template: slide"`
	DisableAutoOpen   bool   `long:"disable-auto-open" description:"Disable automatic opening of the browser"`
	Transition        string `long:"transition" description:"Transition effect for slides: default, cube, page, concave, zoom, linear, fade, none" default:"default"`
	Separator         string `long:"separator" description:"Horizontal slide separator characters" default:"^---"`
	VerticalSeparator string `long:"vertical-separator" description:"Vertical slide separator characters" default:"^___"`
	Multiplex         bool   `long:"multiplex" description:"Enable slide multiplexing"`
	Version           bool   `short:"v" long:"version" description:"Show the version"`
}

func (cli *CLI) Run(args []string) int {
	cli.init()

	args, err := cli.parser.ParseArgs(args)

	if err != nil {
		flagError := err.(*flags.Error)
		if flagError.Type == flags.ErrHelp {
			cli.showHelp()
			return ExitCodeOK
		}
		if flagError.Type == flags.ErrUnknownFlag {
			fmt.Fprintf(cli.OutStream, "Use --help to view all available options.\n")
		} else {
			fmt.Fprintf(cli.OutStream, "Error parsing flags: %s\n", err)
		}
		return ExitCodeError
	}

	if cli.options.Version {
		cli.showVersion()
		return ExitCodeOK
	}

	if len(args) == 0 {
		cli.showHelp()
		return ExitCodeOK
	}

	cli.serve(args)

	return ExitCodeOK
}

func (cli *CLI) init() {
	cli.options = &CLIOptions{}
	cli.parser = cli.newParser()
}

func (cli *CLI) serve(args []string) {
	opts := cli.options
	_, err := os.Stat(opts.Theme)
	originalTheme := false
	if err == nil {
		originalTheme = true
	}

	_, err = os.Stat(opts.Template)
	originalTemplate := false
	if err == nil {
		originalTemplate = true
	}

	server := Server{
		port: opts.Port,
	}
	param := ServerParam{
		Path:              args[0],
		Theme:             addExtention(opts.Theme, "css"),
		Template:          addExtention(opts.Template, "html"),
		Transition:        opts.Transition,
		OriginalTheme:     originalTheme,
		OriginalTemplate:  originalTemplate,
		DisableAutoOpen:   opts.DisableAutoOpen,
		Separator:         opts.Separator,
		VerticalSeparator: opts.VerticalSeparator,
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

func (cli *CLI) showHelp() {
	cli.parser.WriteHelp(cli.OutStream)
}

func (cli *CLI) showVersion() {
	fmt.Fprintf(cli.OutStream, "revealgo version %s\n", Version)
}

func (cli *CLI) newParser() *flags.Parser {
	parser := flags.NewParser(cli.options, flags.HelpFlag|flags.PassDoubleDash)
	parser.Name = "revealgo"
	parser.Usage = "[options] [MARKDOWN FILE]"
	return parser
}
