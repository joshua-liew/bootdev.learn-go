package main

func reformat(message string, formatter func(string) string) string {
	formatted_string := formatter(formatter(formatter(message)))
	return "TEXTIO: " + formatted_string
}
