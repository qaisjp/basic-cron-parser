package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// CronExpression represents a parsed cron spec
type CronExpression struct {
	Minutes    []string
	Hours      []string
	DayOfMonth []string
	Months     []string
	DayOfWeek  []string
	Command    string
}

func (e CronExpression) String() string {
	return fmt.Sprintf(
		"minute\t\t%s"+
			"\nhour\t\t%s"+
			"\nday of month\t%s"+
			"\nmonth\t\t%s"+
			"\nday of week\t%s"+
			"\ncommand\t\t%s",
		strings.Join(e.Minutes, " "),
		strings.Join(e.Hours, " "),
		strings.Join(e.DayOfMonth, " "),
		strings.Join(e.Months, " "),
		strings.Join(e.DayOfWeek, " "),
		e.Command,
	)
}

func parse(str string, min int, max int) ([]string, error) {
	if n, err := strconv.Atoi(str); err == nil {
		if n < min || n > max {
			return nil, errors.Errorf("number (%d) is out of range (%d-%d)", n, min, max)
		}
		return []string{str}, nil
	} else if str == "*" {
		results := []string{}
		for n := min; n <= max; n++ {
			results = append(results, strconv.Itoa(n))
		}
		return results, nil
	} else if part := strings.TrimPrefix(str, "*/"); part != str {
		results := []string{}
		interval, err := strconv.Atoi(part)
		if err != nil {
			return nil, errors.Errorf("unexpected interval %#v for %s", part, str)
		}

		for n := min; n <= max; n += interval {
			results = append(results, strconv.Itoa(n))
		}
		return results, nil
	} else if strings.Contains(str, ",") {
		parts := strings.Split(str, ",")
		results := []string{}
		for _, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, errors.Errorf("unexpected input %#v", part)
			} else if n < min || n > max {
				return nil, errors.Errorf("number (%d) is out of range (%d-%d)", n, min, max)
			}
			results = append(results, part)
		}
		return results, nil
	} else if strings.Contains(str, "-") {
		parts := strings.Split(str, "-")
		if len(parts) != 2 {
			// todo: report errors
			return nil, errors.Errorf("unexpected input %#s", str)
		}

		a, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, errors.Errorf("unexpected input %#s", parts[0])
		} else if a < min {
			return nil, errors.Errorf("number (%d) must be greater than %d", a, min)
		}

		b, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, errors.Errorf("unexpected input %#s", parts[1])
		} else if b > max {
			return nil, errors.Errorf("number (%d) must be less than %d", b, max)
		}

		if a > b {
			return nil, errors.Errorf("unexpected input %#s", str)
		}

		results := []string{}
		for i := a; i <= b; i++ {
			results = append(results, strconv.Itoa(i))
		}

		return results, nil
	}

	return nil, errors.Errorf("unexpected input string")
}

func NewCronExpression(str string) (*CronExpression, error) {
	parts := strings.Fields(os.Args[1])
	if len(parts) < 6 {
		return nil, errors.Errorf("expression is too short")
	}

	var err error

	try := func(sequence []string, parseError error) []string {
		if parseError != nil {
			err = multierror.Append(err, parseError)
		}
		return sequence
	}

	expr := CronExpression{
		Minutes:    try(parse(parts[0], MinMinutes, MaxMinutes)),
		Hours:      try(parse(parts[1], MinHours, MaxHours)),
		DayOfMonth: try(parse(parts[2], MinDayOfMonth, MaxDayOfMonth)),
		Months:     try(parse(parts[3], MinMonths, MaxMonths)),
		DayOfWeek:  try(parse(parts[4], MinDayOfWeek, MaxDayOfWeek)),
		Command:    strings.Join(parts[5:], " "),
	}

	if err != nil {
		return nil, err
	}

	return &expr, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`expr-parser [CRON_EXPRESSION]`)
		return
	}

	expr, err := NewCronExpression(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	fmt.Println(*expr)
}
