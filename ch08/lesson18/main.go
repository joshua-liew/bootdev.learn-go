package main

func isValidPassword(password string) bool {
	passwordLen := 0
	passwordHasUppercaseLetter := false
	passwordHasDigit := false
	for _, char := range password {
		// ASCII Uppercase range: A(65) <-> Z(90)
		if int(char) >= 65 && int(char) <= 90 {
			passwordHasUppercaseLetter = true
		}
		// ASCII Digit range: 0(48) <-> 9(57)
		if int(char) >= 48 && int(char) <= 57 {
			passwordHasDigit = true
		}
		passwordLen++
	}

	if passwordLen < 5 || passwordLen > 12 {
		return false
	}
	if !passwordHasUppercaseLetter || !passwordHasDigit {
		return false
	}
	return true
}
