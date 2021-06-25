package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	var s string

	var x int
	var y int
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s)

		x, _ = strconv.Atoi(string(s[0]))
		y, _ = strconv.Atoi(string(s[2]))

		if x == y {
			fmt.Println(x * y)
		} else if unicode.IsUpper(rune(s[1])) {
			fmt.Println(y - x)
		} else {
			fmt.Println(y + x)
		}

	}
}
