package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	// TODO: print the text of line, uppercased, with no whitespace around it.
	// Right now it prints the line untouched, which is wrong for every test.
	line = strings.TrimSpace(line)
	line = strings.ToUpper(line)
	fmt.Println(line)
}
