# logfmt2json
[![GoDoc](https://godoc.org/github.com/ntindall/logfmt2json?status.svg)](https://godoc.org/github.com/ntindall/logfmt2json)

reads [logfmt](https://brandur.org/logfmt) log messages from `stdin` and prints json to `stdout`

## Installation
```sh
  go get -u github.com/ntindall/logfmt2json
```

## Usage

```sh
echo "foo=bar baz=bak" | logfmt2json
{"baz":"bak","foo":"bar"}
```
