# go-memoizer

A simple thread-safe memoizer for Go using only the standard library.

Supports Go 1.18 and higher due to the use of generics.

[![Reference](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/LightningDev1/go-memoizer)
[![Linter](https://goreportcard.com/badge/github.com/LightningDev1/go-memoizer?style=flat-square)](https://goreportcard.com/report/github.com/LightningDev1/go-memoizer)
[![Build status](https://github.com/LightningDev1/go-memoizer/actions/workflows/ci.yml/badge.svg)](https://github.com/LightningDev1/go-memoizer/actions)

```go
func expensiveFunction() (*any, error) {
    // ...
}

var memoizer = memoizer.New(expensiveFunction, 10*time.Second)

for i := 0; i < 50; i++ {
    value, err := memoizer.Get()

    // ...
}
```

## Installation

```bash
go get github.com/LightningDev1/go-memoizer
```

## Usage

See [config example](./examples/config/main.go) for a complete example program.
