# go-commando

[![Build](https://github.com/J-R-Oliver/go-commando/actions/workflows/build.yml/badge.svg)](https://github.com/J-R-Oliver/go-commando/actions/workflows/build.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/J-R-Oliver/go-commando)](https://github.com/gomods/athens)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/J-R-Oliver/go-commando)](https://goreportcard.com/report/github.com/J-R-Oliver/go-commando)

<table>
<tr>
<td>
A package for building Go command-line applications. Inspired by Command.js.
</td>
</tr>
</table>

## Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Declaring a Program](#declaring-a-program)
- [Options](#options)
- [Action](#action)
- [Parse](#action)
- [Name](#action)
- [Description](#action)
- [Version](#action)
- [Help](#action)
- [Local Development](#local-development)
- [Testing](#testing)
- [Conventional Commits](#conventional-commits)
- [GitHub Actions](#github-actions)

## Installation

To install `go-commando` using Go modules execute:

```shell
go get github.com/J-R-Oliver/go-commando
```

You'll then be able to import `go-commando` in your code:

```go
import "github.com/J-R-Oliver/go-commando"
```

## Quick Start

`go-commando` allows you to write code to configure your command line applications. `go-commando` automates the parsing of
options and arguments, and implements a help text option.

An example application can be found in [example.go](./example/example.go):

```go
package main

import (
    "github.com/J-R-Oliver/go-commando"
)

func main() {
    fileSplitter := func(arguments []string, options map[string]string) {
        // implementation removed for brevity
    }
    
    program := commando.NewProgram()
    
    program.
        Name("file-splitter").
        Description("CLI to split file written in go.").
        Version("1.0.0").
        Option("i", "input", "input", "Input file", "./input.txt").
        Option("o", "output", "output", "Output file", "./output.txt").
        Action(fileSplitter).
        Parse()
}
```

`program` contains all the methods required to configure your application. This configuration is used to parse the 
application inputs and build the help text.

## Declaring a `program`

A new `program` can be created by calling the `NewProgram` function.

```go
program := commando.NewProgram()
```

## Options

Options can be configured by calling the `Option` method. All options are parsed as string variables. An option can be 
set with a `shortOption`, `longOption`, `mapKey`, `description` and `defaultValue`.

```go
program.Option("i", "input", "input", "Input file", "./input.txt")
```

To configure either a short option only, or long option only, pass `""` to the unrequired option type.

```go
program.Option("i", "", "input", "Input file", "./input.txt") // short option
program.Option("", "input", "input", "Input file", "./input.txt") // long option
```

## Action

The action handler receives both the arguments and options given when the application is run. Action is the entry point 
when building command line applications using commando. `arguments` is a slice of strings containing all user input when 
executing application after any options have been parsed. This slice maintains the order of the inputted arguments. 
`options` is a map of strings containing options input. Specific options can be accessed using the mapKey set when adding 
the option.

For example, the following action prints out the `arguments` and `options`:

```go
func(arguments []string, options map[string]string) {
    fmt.Println("Arguments:")

    for i, a := range arguments {
        fmt.Printf("\tindex: %d, argument: %s\n", i, a)
    }

    fmt.Println("Options:")

    for k, v := range options {
        fmt.Printf("\tkey: %s, option: %s\n", k, v)
    }
}
```

## Parse

Parse initiates starting the program and should be the final function call on program. Once the desired program 
configuration has been loaded the action function will be called with the program arguments and options.

## Name

Name sets the name of the program, used when creating the `-h` or `--help` output, and returns a pointer to the program.

```go
program.Name("file-splitter")
```

## Description

Description sets the description of the program, used when creating the `-h` or `--help` output, and returns a pointer to 
the program.

```go
program.Description("CLI to split file written in go.")
```

## Version

Version sets the version of the program, returned when the application is called with `-v` or `--version`, and returns a 
pointer to the program.

```go
program.Version("1.0.0")
```

Example console output when application is called with `-v`:

```console
$ file-splitter -v
1.0.0
```

## Help

Help text is automatically created by `go-commando` and can be viewed when executing applications with either `-h` or
`--help` options. The help text is also displayed if an unconfigured option is entered.

```console
$ file-splitter -h
Usage: file-splitter [options] [arguments]

CLI to split file written in go.

Options:
  -i, --input <input>                     Input file (default: "./input.txt")
  -o, --output <output>                   Output file (default: "./output.txt")
  -v, --version                           output the version number
  -h, --help                              display help for command
```

## Local Development

### Prerequisites

To install and modify this project you will need to have:

- [Go](https://go.dev)
- [Git](https://git-scm.com)

### Installation

To start, please `fork` and `clone` the repository to your local machine.

## Testing

All tests have been written using the [testing](https://pkg.go.dev/testing) package from the
[Standard library](https://pkg.go.dev/std). To run the tests execute:

```shell
go test -v ./...
```

Code coverage is also measured by using the `testing` package. To run tests with coverage execute:

```shell
go test -coverprofile=coverage.out  ./...
```

## Conventional Commits

This project uses the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for commit
messages. The specification provides a simple rule set for creating commit messages, documenting features, fixes, and
breaking changes in commit messages.

A [pre-commit](https://pre-commit.com) [configuration file](.pre-commit-config.yaml) has been provided to automate
commit linting. Ensure that *pre-commit* has been [installed](https://www.conventionalcommits.org/en/v1.0.0/) and
execute...

```shell
pre-commit install
````

...to add a commit [Git hook](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks) to your local machine.

An automated pipeline job has been [configured](.github/workflows/build.yml) to lint commit messages on a push.

## GitHub Actions

A CI/CD pipeline has been created using [GitHub Actions](https://github.com/features/actions) to automated tasks such as
linting and testing.

### Build Workflow

The [build](./.github/workflows/build.yml) workflow handles integration tasks. This workflow consists of two jobs, `Git`
and `Go`, that run in parallel. This workflow is triggered on a push to a branch.

#### Git

This job automates tasks relating to repository linting and enforcing best practices.

#### Go

This job automates `Go` specific tasks.
