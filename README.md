# go-commando

<table>
<tr>
<td>
A package for building Go command-line applications. Inspired by Command.js.
</td>
</tr>
</table>

## Contents

- [Conventional Commits](#conventional-commits)
- [GitHub Actions](#github-actions)

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
