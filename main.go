package main

import (
	"fmt"
	"os"

	"github.com/charlesbases/spider/internal/app"
)

func main() {
	if err := app.Spider().Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
