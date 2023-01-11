package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		return
	}

	fields := strings.Fields(args[1])

	fmt.Println(len(fields))
}
