package main

func countConnections(groupSize int) int {
	var conn int
	for i := 0; i < groupSize; i++ {
		conn += i
	}
	return conn
}
