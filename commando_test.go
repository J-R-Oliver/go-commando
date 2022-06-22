package commando

import (
	"os"
	"reflect"
	"testing"
)

func TestNewProgram(t *testing.T) {
	expected := &Program{}

	if p := NewProgram(); !reflect.DeepEqual(p, expected) {
		t.Errorf("NewProgram() = %v, want %v", p, expected)
	}
}

func TestProgram_Name(t *testing.T) {
	p := &Program{}
	expected := &Program{name: "Test name"}

	if p = p.Name("Test name"); !reflect.DeepEqual(p, expected) {
		t.Errorf("Name() = %v, want %v", p, expected)
	}
}

func TestProgram_Description(t *testing.T) {
	p := &Program{}
	expected := &Program{description: "Test description"}

	if p = p.Description("Test description"); !reflect.DeepEqual(p, expected) {
		t.Errorf("Description() = %v, want %v", p, expected)
	}
}

func TestProgram_Version(t *testing.T) {
	p := &Program{}
	expected := &Program{version: "Test version"}

	if p = p.Version("Test version"); !reflect.DeepEqual(p, expected) {
		t.Errorf("Version() = %v, want %v", p, expected)
	}
}

func TestProgram_Option(t *testing.T) {
	p := &Program{}
	expected := &Program{options: []option{{
		shortOption:  "t",
		longOption:   "test-option",
		mapKey:       "test",
		description:  "Test option",
		defaultValue: "Test",
	}}}

	if p = p.Option("t", "test-option", "test", "Test option", "Test"); !reflect.DeepEqual(p, expected) {
		t.Errorf("Option() = %v, want %v", p, expected)
	}
}

func TestProgram_Action(t *testing.T) {
	f := func(arguments []string, options map[string]string) {}

	p := &Program{}
	expected := &Program{action: f}

	if p = p.Action(f); reflect.ValueOf(p.action).Pointer() != reflect.ValueOf(expected.action).Pointer() {
		t.Errorf("Action() = %v, want %v", p, expected)
	}
}

func TestProgram_Parse(t *testing.T) {
	testAction := func(arguments []string, options map[string]string) {
		if arguments[0] != "argument1" {
			t.Errorf("Parse() = %s, want %s", arguments[0], "argument1")
		}

		if arguments[1] != "argument2" {
			t.Errorf("Parse() = %s, want %s", arguments[0], "argument2")
		}

		if options["shortOption"] != "shortOption" {
			t.Errorf("Parse() = %s, want %s", options["shortOption"], "shortOption")
		}

		if options["longOption"] != "longOption" {
			t.Errorf("Parse() = %s, want %s", options["longOption"], "longOption")
		}

		if options["defaultOption"] != "default" {
			t.Errorf("Parse() = %s, want %s", options["defaultOption"], "default")
		}

		if options["shortOnlyOption"] != "shortOnlyOption" {
			t.Errorf("Parse() = %s, want %s", options["shortOnlyOption"], "shortOnlyOption")
		}

		if options["longOnlyOption"] != "longOnlyOption" {
			t.Errorf("Parse() = %s, want %s", options["longOnlyOption"], "longOnlyOption")
		}
	}

	os.Args = []string{"./commando_test.go", "-o", "shortOption", "--option-two", "longOption", "-s", "shortOnlyOption", "--long-option", "longOnlyOption", "argument1", "argument2"}

	NewProgram().
		Option("o", "option-one", "shortOption", "Test option one", "default").
		Option("p", "option-two", "longOption", "Test option two", "default").
		Option("n", "option-three", "defaultOption", "Test option three", "default").
		Option("s", "", "shortOnlyOption", "Test option four", "default").
		Option("", "long-option", "longOnlyOption", "Test option five", "default").
		Action(testAction).
		Parse()
}

func TestProgram_helpText(t *testing.T) {
	p1 := NewProgram()
	e1 := "Usage:  [options] [arguments]\n\nOptions:\n  -h, --help                              display help for command\n"

	p2 := NewProgram().Name("Test command")
	e2 := "Usage: Test command [options] [arguments]\n\nOptions:\n  -h, --help                              display help for command\n"

	p3 := NewProgram().Description("Test description")
	e3 := "Usage:  [options] [arguments]\n\nTest description\n\nOptions:\n  -h, --help                              display help for command\n"

	p4 := NewProgram().Option("o", "option", "option", "Test option", "")
	e4 := "Usage:  [options] [arguments]\n\nOptions:\n  -o, --option <option>                   Test option\n  -h, --help                              display help for command\n"

	p5 := NewProgram().Option("o", "option", "option", "Test option", "default")
	e5 := "Usage:  [options] [arguments]\n\nOptions:\n  -o, --option <option>                   Test option (default: \"default\")\n  -h, --help                              display help for command\n"

	p6 := NewProgram().Option("o", "", "option", "Test option", "default")
	e6 := "Usage:  [options] [arguments]\n\nOptions:\n  -o <option>                             Test option (default: \"default\")\n  -h, --help                              display help for command\n"

	p7 := NewProgram().Option("", "option", "option", "Test option", "default")
	e7 := "Usage:  [options] [arguments]\n\nOptions:\n  --option <option>                       Test option (default: \"default\")\n  -h, --help                              display help for command\n"

	p8 := NewProgram().Version("1.0.0")
	e8 := "Usage:  [options] [arguments]\n\nOptions:\n  -v, --version                           output the version number\n  -h, --help                              display help for command\n"

	tests := []struct {
		name     string
		program  *Program
		expected string
	}{
		{"Returns minimal help", p1, e1},
		{"Returns help with command name", p2, e2},
		{"Returns help with description", p3, e3},
		{"Returns help with option", p4, e4},
		{"Returns help with option and default value", p5, e5},
		{"Returns help with short option only", p6, e6},
		{"Returns help with long option only", p7, e7},
		{"Returns help with version", p8, e8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.program.helpText()

			if h != tt.expected {
				t.Errorf("help():\nGot:\n%sWant:\n%s", h, tt.expected)
			}
		})
	}
}
