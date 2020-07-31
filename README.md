# multiwriter

[![Build Status](https://travis-ci.org/alanshaw/multiwriter.svg?branch=master)](https://travis-ci.org/alanshaw/multiwriter)
[![Coverage](https://codecov.io/gh/alanshaw/multiwriter/branch/master/graph/badge.svg)](https://codecov.io/gh/alanshaw/multiwriter)
[![Standard README](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)
[![pkg.go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/alanshaw/multiwriter)
[![golang version](https://img.shields.io/badge/golang-%3E%3D1.14.0-orange.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanshaw/multiwriter)](https://goreportcard.com/report/github.com/alanshaw/multiwriter)

A writer that writes to multiple other writers _and_ the writers can be added and removed _dynamically_.

## Install

```sh
go get github.com/alanshaw/multiwriter
```

## Usage

Example:

```go
package main

import (
	"os"
	"github.com/alanshaw/multiwriter"
)

func main() {
	w := multiwriter.New(os.Stdout, os.Stderr)

	w.Write([]byte("written to stdout AND stderr\n"))

	w.Remove(os.Stderr)

	w.Write([]byte("written to ONLY stdout\n"))

	w.Remove(os.Stdout)
	w.Add(os.Stderr)

	w.Write([]byte("written to ONLY stderr\n"))
}
```

## API

[pkg.go.dev Reference](https://pkg.go.dev/github.com/alanshaw/multiwriter)

## Contribute

Feel free to dive in! [Open an issue](https://github.com/alanshaw/multiwriter/issues/new) or submit PRs.

## License

[MIT](LICENSE) © Alan Shaw
