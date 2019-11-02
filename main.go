package main

import (
	"fmt"
	"os"
	"strings"
)

type CronExpression struct {
	Minutes    []string
	Hours      []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    string
}

func (e CronExpression) String() string {
	return fmt.Sprintf(
		`minute\t\t: %s
hour\t\t: %s
day of month\t: %s
month\t\t: %s
day of week\t: %s
command\t\t: %s`,
		strings.Join(e.Minutes, " "),
		strings.Join(e.Hours, " "),
		strings.Join(e.DayOfMonth, " "),
		strings.Join(e.Month, " "),
		strings.Join(e.DayOfWeek, " "),
		e.Command,
	)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`expr-parser [CRON_EXPRESSION]`)
		return
	}

	expr := os.Args[1]
	fmt.Println("Hi", expr)
}
