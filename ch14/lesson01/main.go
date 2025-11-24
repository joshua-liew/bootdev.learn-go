package main

func getLast[T any](s []T) T {
	length := len(s)
	if length == 0 {
		var zeroVal T
		return zeroVal
	}
	return s[length-1]
}
