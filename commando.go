package commando

import (
	"flag"
	"regexp"
)

var optionExpression = regexp.MustCompile(`\s*-{1,2}([\w\d\-]+)\s*,\s*-{1,2}([\w\d\-]+)\s*<([\w\d]+)>\s*`)

type Program struct {
	name        string
	description string
	version     string
	options     []option
	action      action
}

type option struct {
	name         string
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

func (p *Program) Option(name, description, defaultValue string) *Program {
	o := option{name, description, defaultValue}
	p.options = append(p.options, o)

	return p
}

func (p *Program) Action(action action) *Program {
	p.action = action
	return p
}

func (p *Program) Parse() {
	o := p.parseOptions()

	p.action(flag.Args(), o)
}

func (p *Program) parseOptions() map[string]string {
	flags := make(map[string]*string)

	for _, option := range p.options {
		s := new(string)

		name, longOption, value := parseOptionName(option.name)

		flag.StringVar(s, name, option.defaultValue, option.description)
		flag.StringVar(s, longOption, option.defaultValue, option.description)
		flags[value] = s
	}

	flag.Parse()

	parsedOptions := make(map[string]string)
	for k, v := range flags {
		parsedOptions[k] = *v
	}

	return parsedOptions
}

func parseOptionName(n string) (shortOption, longOption, value string) {
	s := optionExpression.FindStringSubmatch(n)
	return s[1], s[2], s[3]
}
