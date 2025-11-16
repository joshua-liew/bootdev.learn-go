package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		isPrime := true
		if n == 2 {
			fmt.Printf("%v\n", n)
			continue
		}
		if n % 2 == 0 {
			isPrime = false
			continue
		}
		for i := 3; i * i <= n; i += 2 {
			//fmt.Printf("res: '%v' %% '%v' is '%v'\n", n, i, (n % i))
			if n % i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%v\n", n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}

