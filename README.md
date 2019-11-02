# expr-parser

Deliveroo Technical Task

## Usage

```
expr-parser [CRON_EXPRESSION]
```

`expr-parser` takes a single argument, `CRON_EXPRESSION`, which is a cron string like `"*/15 0 1,15 * 1-5 /usr/bin/find -L . -print0"`.

A cron expression must consist of five "parts" followed by a command.

Supported features:
- asterisk (`*`)
- basic slashes, for some time interval `x` (e.g. `*/15)
    - note that this includes the smallest value, and you are expected to provide intervals that evenly divide
- basic commas (e.g. `1,2`)
- hyphens (e.g. `1-7`)

Unsupported features:
- Other forms of slashes
- [Special time strings](https://en.wikipedia.org/wiki/Cron#Nonstandard_predefined_scheduling_definitions))
- Words in place of Month or Day of week
- Question marks

In practice, you might prefer to use [robfig's cron library](https://github.com/robfig/cron).
