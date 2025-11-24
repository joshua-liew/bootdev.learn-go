package main

type emailStatus int
const (
	emailBounced = iota
	emailInvalid
	emailDelivered
	emailOpened
)
