package main

type formatter interface {
	format() (formattedString string)
}

type plainText struct {
	message string
}
type bold struct {
	message string
}
type code struct {
	message string
}

func (p plainText) format() (formattedString string) {
	return p.message
}
func (b bold) format() (formattedString string) {
	return "**" + b.message + "**"
}
func (c code) format() (formattedString string) {
	return "`" + c.message + "`"
}

func sendMessage(format formatter) string {
	return format.format()
}
