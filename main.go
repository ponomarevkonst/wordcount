package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	args := os.Args[1:]
	s := strings.Join(args, " ")
	words := strings.Fields(s)
	fmt.Println(len(words))
}
