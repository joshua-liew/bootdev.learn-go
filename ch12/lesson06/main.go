package main

func concurrentFib(n int) []int {
	res := make([]int, 0)
	ch := make(chan int)
	go fibonacci(n, ch)
	for x := range ch {
		res = append(res, x)
	}
	return res
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
