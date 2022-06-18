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

- [Getting Started](#getting-started)
- [Testing](#testing)
- [Conventional Commits](#conventional-commits)
- [GitHub Actions](#github-actions)

## Getting Started

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
