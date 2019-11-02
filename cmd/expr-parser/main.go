package main

import (
	"fmt"
	"os"

	cron "github.com/qaisjp/basic-cron-parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`expr-parser [CRON_EXPRESSION]`)
		return
	}

	expr, err := cron.NewCronExpression(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	fmt.Println(*expr)
}
