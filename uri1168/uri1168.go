package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	ledsByDigit := map[int]int{
		0: 6,
		1: 2,
		2: 5,
		3: 5,
		4: 4,
		5: 5,
		6: 6,
		7: 3,
		8: 7,
		9: 6,
	}

	var x int
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanf("%s", &str)

		sum := 0
		for i := 0; i < len(str); i++ {
			x, _ = strconv.Atoi(string(str[i]))

			sum += ledsByDigit[x]
		}

		fmt.Println(sum, "leds")
	}
}
