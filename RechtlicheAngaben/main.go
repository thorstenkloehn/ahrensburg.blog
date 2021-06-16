package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
)

func main() {
	input := "## Hallo Leute"

	output := blackfriday.Run([]byte(input))

	fmt.Println(string(output))
}
