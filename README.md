# mygithelper

This is a toy project to learn golang. The idea of the project is based on this course: [Write Professional Command-line Programs in Go](https://www.educative.io/courses/prof-command-line-programs-go)

## Description

This is a wrapper aroung git. The tool provides a set of commands that replaicate or extend the existing git functionality. Also, this tool allows to apply this command recursively in several repositories.

In general this repository is going trough the same steps as the course, so the first commit starts from the simpliest possible version and every next one gradually improve the repo.

**Usage:**
```
mgh stands for MyGitHelper - wrapper around git.
It replicates and extends several git commands.
Also it allows to recursively apply git commands
to several repositories.

Usage:
  mgh [flags]
  mgh [command]

Available Commands:
  addall      Add all files in the repo
  clean       Clean repository from something described by given parameter
  commit      Commit all files in the repo
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  pull        Pull all files in the repo
  push        Push all files in the repo

Flags:
  -h, --help        help for mgh
  -m, --multi-run   mgh will go through all repositories in the current directory
                    and apply the command there.
  -v, --version     version for mgh

Use "mgh [command] --help" for more information about a command.
```

**Thigs to add later:**
- `mgh init <folders>`
- `mgh add all`
- `mgh add py`
- `mgh add go`
- `mgh clean <file type>`


