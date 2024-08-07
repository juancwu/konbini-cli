package main

import (
	"fmt"
	"os"

	"github.com/juancwu/mi/cmd"
	"github.com/juancwu/mi/text"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%s\n", text.Foreground(text.RED, err.Error()))
		os.Exit(1)
	}
}
