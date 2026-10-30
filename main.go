package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    // TODO: accumulate the sum of 1..n with a loop, and print that instead of the placeholder.
    sum := 0
    for ran := 1; ran <= n; ran++ {
        sum = sum + ran
    }
    fmt.Println(sum)
}
