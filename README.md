# Makeit

Makeit is a simple tool to help you view and run your Makefile targets from the command line.

![makeit](./assets/default.png)
![makeit](./assets/verbose.png)

## Installation

```shell
go install github.com/nycruz/makeit@latest
```

## Usage

Default: use the Makefile in the current directory

```shell
makeit
```

Specify a custom Makefile

```shell
makeit -f Makefile.dev
```

Verbose output. Show the command

```shell
makeit -v
```
