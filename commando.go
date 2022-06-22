// Package commando provides a succinct and simple method for creating command line applications. Primarily commando
// wraps existing implementations in package flag.
package commando // import "github.com/J-R-Oliver/go-commando"

import (
	"flag"
	"fmt"
	"os"
)

// Program represents a command line application.
type Program struct {
	name          string
	description   string
	version       string
	options       []option
	parsedOptions map[string]*string
	action        Action
	showVersion   *bool
}

type option struct {
	shortOption  string
	longOption   string
	mapKey       string
	description  string
	defaultValue string
}

// Action is a type of function that receives both the arguments and options given when the application is run. Action
// is the entry point when building command line applications using commando. arguments is a slice of strings containing
// all user input when executing application after any options have been parsed. This slice maintains the order of the
// inputted arguments. options is a map of strings containing options input. Specific options can be accessed using the
// mapKey set when adding the option.
type Action func(arguments []string, options map[string]string)

// NewProgram returns a pointer to a new unconfigured program.
func NewProgram() *Program {
	return &Program{}
}

// Name sets the name of the program, used when creating the -h or --help output, and returns a pointer to the program.
func (p *Program) Name(name string) *Program {
	p.name = name
	return p
}

// Description sets the description of the program, used when creating the -h or --help output, and returns a pointer
// to the program.
func (p *Program) Description(description string) *Program {
	p.description = description
	return p
}

// Version sets the version of the program, returned when the application is called with -v or --version, and returns a
// pointer to the program.
func (p *Program) Version(version string) *Program {
	p.version = version
	return p
}

// Option adds a commandline option, and returns a pointer to the program.
//
// Option takes a shortOption and a longOption string. Commando will automatically add '-' for short options and '--'
// for long options. Only short or only long options can be configured by passing "" for the unneeded variant.
//
// mapKey is used as the key to store the option input in the map passed to the action function.
// description is used when creating the -h or --help output.
// defaultValue will be set in the options map passed to the action function and be overridden by user input.
func (p *Program) Option(shortOption, longOption, mapKey, description, defaultValue string) *Program {
	o := option{shortOption, longOption, mapKey, description, defaultValue}
	p.options = append(p.options, o)

	return p
}

// Action sets the action function of the program, and returns a pointer to the program.
func (p *Program) Action(action Action) *Program {
	p.action = action
	return p
}

// Parse initiates starting the program and should be the final function call on program. Once the desired program
// configuration has been loaded the action function will be called with the program arguments and options.
func (p *Program) Parse() {
	p.parseOptions()
	p.addVersionOption()

	// Using an anonymous function to help with the testing of helpText
	flag.Usage = func() {
		fmt.Print(p.helpText())
	}

	flag.Parse()

	if *p.showVersion {
		fmt.Println(p.version)
		os.Exit(0)
	}

	o := p.dereferenceParsedOptionsMap()

	p.action(flag.Args(), o)
}

func (p *Program) parseOptions() {
	p.parsedOptions = make(map[string]*string)

	for _, o := range p.options {
		s := new(string)

		if o.shortOption != "" {
			flag.StringVar(s, o.shortOption, o.defaultValue, o.description)
		}

		if o.longOption != "" {
			flag.StringVar(s, o.longOption, o.defaultValue, o.description)
		}

		p.parsedOptions[o.mapKey] = s
	}
}

func (p *Program) addVersionOption() {
	v := new(bool)

	flag.BoolVar(v, "v", false, "Show program version")
	flag.BoolVar(v, "version", false, "Show program version")

	p.showVersion = v
}

func (p *Program) helpText() string {
	h := fmt.Sprintf("Usage: %s [options] [arguments]\n", p.name)

	if p.description != "" {
		h += fmt.Sprintf("\n%s\n", p.description)
	}

	h += "\nOptions:\n"

	for _, o := range p.options {
		n := p.helpTextOptionName(o)

		h += fmt.Sprintf("  %-40s%s", n, o.description)

		if o.defaultValue != "" {
			h += fmt.Sprintf(" (default: \"%s\")\n", o.defaultValue)
		} else {
			h += "\n"
		}
	}

	if p.version != "" {
		h += fmt.Sprintf("  %-40s%s\n", "-v, --version", "output the version number")
	}

	h += fmt.Sprintf("  %-40s%s\n", "-h, --help", "display help for command")

	return h
}

func (p *Program) helpTextOptionName(o option) string {
	var n string
	if o.shortOption != "" {
		n += fmt.Sprintf("-%s", o.shortOption)
	}

	if o.shortOption != "" && o.longOption != "" {
		n += ", "
	}

	if o.longOption != "" {
		n += fmt.Sprintf("--%s", o.longOption)
	}

	n += fmt.Sprintf(" <%s>", o.mapKey)

	return n
}

func (p *Program) dereferenceParsedOptionsMap() map[string]string {
	parsedOptions := make(map[string]string)

	for k, v := range p.parsedOptions {
		parsedOptions[k] = *v
	}

	return parsedOptions
}
