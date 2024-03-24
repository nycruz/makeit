# makeit

**makeit** is a command-line tool designed to enhance your productivity by enabling easy viewing and execution of Makefile targets.
Whether you're working on a small project or managing complex builds, **makeit** simplifies your workflow by bringing the power of Makefiles directly to your terminal.

| ![makeit](./assets/verbose.png) |
| :-----------------------------: |
|         _verbose mode_          |

## Features

**Ease of Use**: Instantly view and run Makefile targets from the command line without memorizing complex commands.

**Flexibility**: Works with any Makefile in your current directory or specify a custom Makefile to use.

**Enhanced Feedback**: Opt for verbose output to see exactly what commands are being executed, helping you understand your build process better.

## Installation

To install **makeit**, ensure you have [Go installed](https://formulae.brew.sh/formula/go) on your system, then run the following command:

```sh
go install github.com/nycruz/makeit@latest
```

This command installs the latest version of **makeit**, making it available in your system's PATH.

## Getting Started

### Viewing and Running Makefile Targets

- **Default Usage**: Simply run `makeit` in the terminal to view and execute targets from the Makefile in your current directory.

```sh
makeit
```

| ![makeit](./assets/default.png) |
| :-----------------------------: |
|         _default mode_          |

- **Using a Custom Makefile**: To use a different Makefile, use the `-f` option followed by the path to your custom Makefile.

```sh
makeit -f Makefile.dev
```

- **Verbose Output**: For a detailed view of the commands being executed, use the `-v` option to enable verbose output.

```sh
makeit -v
```

| ![makeit](./assets/verbose.png) |
| :-----------------------------: |
|         _verbose mode_          |
