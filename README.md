# README

## Build and Deploy to Local

```shell
sh deploy-local.sh
```

## fluent cli help

```shell
❯ fluent -h       
fluent cli is a collection of daily scripts

Usage:
  fluent [flags]
  fluent [command]

Examples:
fluent <category> <action> <--options>

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  issue       Create a new issue
  starter     fluent start java,go,ts-lib,fe-next,mono,python
  util        fluent  util echo,mac-cleanup

Flags:
  -h, --help      help for fluent
  -v, --version   version for fluent

Use "fluent [command] --help" for more information about a command.

```

## fluent starer cli

create template project: ***fluent starter <name>***

```shell
fluent starter -h 


Usage:
  fluent starter [flags]

Examples:
fluent start java

Flags:
  -h, --help   help for starter

```

```shell
fluent starter -h
fluent starter java
fluent starter python
fluent starter go
fluent starter ts-lib
fluent starter mono
```

## fluent util cli

***fluent util <cmd>***, to run ,and add more commands in utils.json file

```shell
fluent  util echo,mac-cleanup

Usage:
  fluent util [flags]

Examples:
fluent  util echo

Flags:
  -h, --help   help for util

```

## action json file

action structure:
- name: command name
- desc: command usage
- commands: a list of shell commands or something else which is runnable 

```json
{
  "actions": [
    {
      "name": "echo",
      "desc": "try cli",
      "commands": [
        "echo \"hello fluent cli\""
      ]
    },    {
      "name": "mac-cleanup",
      "desc": "cleanup mac pc to free disk space",
      "commands": [
        "sh mac-cleanup.sh"
      ]
    }
  ]
}
```




## Pocketbase Server

```shell
./pocketbase serve
```

All Url:

- http://127.0.0.1:8090 - if pb_public directory exists, serves the static content from it (html, css, images, etc.)
- http://127.0.0.1:8090/_/ - Admin dashboard UI
- http://127.0.0.1:8090/api/ - REST API



### References 

- [github-cli](https://github.com/cli/cli.git)

