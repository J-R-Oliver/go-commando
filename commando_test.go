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
		name:         "-t, --test-option <test>",
		description:  "Test option",
		defaultValue: "Test",
	}}}

	if p = p.Option("-t, --test-option <test>", "Test option", "Test"); !reflect.DeepEqual(p, expected) {
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
	}

	os.Args = []string{"./commando_test.go", "-o", "shortOption", "--option-two", "longOption", "argument1", "argument2"}

	NewProgram().
		Option("-o, --option-one <shortOption>", "Test option one", "default").
		Option("-p, --option-two <longOption>", "Test option two", "default").
		Option("-n, --option-three <defaultOption>", "Test option three", "default").
		Action(testAction).
		Parse()
}
