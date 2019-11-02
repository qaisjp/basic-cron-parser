# basic-cron-parser

[![GoDoc](https://godoc.org/github.com/qaisjp/basic-cron-parser?status.svg)](https://godoc.org/github.com/qaisjp/basic-cron-parser)

`basic-cron-parser` is a package for Go that allows you parse cron expressions.

Supported features:
- expressions must consist of five "parts" followed by a command.
- asterisk (`*`)
- basic slashes, for some time interval `x` (e.g. `*/15)
    - note that this includes the smallest value, and you are expected to provide intervals that evenly divide
- basic commas (e.g. `1,2`)
- hyphens (e.g. `1-7`)

Unsupported features:
- other forms of slashes
- [special time strings](https://en.wikipedia.org/wiki/Cron#Nonstandard_predefined_scheduling_definitions)
- words in place of Month or Day of week
- question marks

This repository is an experiment. In practice, you should use [robfig's cron library](https://github.com/robfig/cron).

## Installation

Install using `go get github.com/qaisjp/basic-cron-parser/cmd/expr-parser`.

If you have `$GOPATH/bin` in your `$PATH`, you should be able to run `expr-parser`.

Otherwise, you can do something like `go run github.com/qaisjp/basic-cron-parser/cmd/expr-parser "*/15 0 1,15 * 1-5 /usr/bin/find arg1 arg2"`.

## Usage

```
expr-parser [CRON_EXPRESSION]
```

`expr-parser` takes a single argument, `CRON_EXPRESSION`, which is a cron string like `"*/15 0 1,15 * 1-5 /usr/bin/find -L . -print0"`.
