package commando

import (
	"flag"
	"fmt"
)

type Program struct {
	name          string
	description   string
	version       string
	options       []option
	parsedOptions map[string]*string
	action        action
}

type option struct {
	shortOption  string
	longOption   string
	mapKey       string
	description  string
	defaultValue string
}

type action func(arguments []string, options map[string]string)

func NewProgram() *Program {
	return &Program{}
}

func (p *Program) Name(name string) *Program {
	p.name = name
	return p
}

func (p *Program) Description(description string) *Program {
	p.description = description
	return p
}

func (p *Program) Version(version string) *Program {
	p.version = version
	return p
}

func (p *Program) Option(shortOption, longOption, mapKey, description, defaultValue string) *Program {
	o := option{shortOption, longOption, mapKey, description, defaultValue}
	p.options = append(p.options, o)

	return p
}

func (p *Program) Action(action action) *Program {
	p.action = action
	return p
}

func (p *Program) Parse() {
	p.parseOptions()

	// Using an anonymous function to help with the testing of helpText
	flag.Usage = func() {
		fmt.Print(p.helpText())
	}

	flag.Parse()

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
