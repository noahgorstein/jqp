package main

import (
	"os"

	"github.com/noahgorstein/jqp/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		// error is discarded as cobra already reported it
		os.Exit(1)
	}
}
