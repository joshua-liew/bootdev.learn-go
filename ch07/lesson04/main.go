package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz() {
	accumulator := ""
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			accumulator += "fizz"
		}
		if i % 5 == 0 {
			accumulator += "buzz"
		}
		if i % 3 != 0 && i % 5 != 0 {
			accumulator += strconv.Itoa(i)
		}
		accumulator += "\n"
	}
	fmt.Printf(accumulator)
}

func main() {
	fizzbuzz()
}
