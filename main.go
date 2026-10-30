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
	words := strings.Fields(strings.TrimSpace(line))
	counts := make(map[string]int)

	// TODO: put every word of words into counts, then set these two numbers.
	different := 0
	once := 0
	_ = words
	_ = counts
	for _, val := range words {
		counts[val]++
	}
	for key := range counts {
		if counts[key] == 1 {
			once++
		}
	}
	different = len(counts)
	fmt.Println(different)
	fmt.Println(once)
}
