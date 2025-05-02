package helpers

import (
	"regexp"
	"unicode"

	emailVerifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailVerifier.NewVerifier()
)

func CheckPassword(password string) (bool, string) {
	if len(password) < 6 {
		return false, "Password length should be greater than 6"
	}
	containsUpper := false
	containsLower := false
	containsDigits := false
	containsSpecialCharacters := false

	for _, ch := range password {
		if unicode.IsDigit(ch) {
			containsDigits = true
		} else if unicode.IsLower(ch) {
			containsLower = true
		} else if unicode.IsUpper(ch) {
			containsUpper = true
		} else {
			containsSpecialCharacters = true
		}
	}
	if containsDigits && containsLower && containsSpecialCharacters && containsUpper {
		return true, "Password is valid"
	} else {
		return false, "Password should contains uppercase, lowercase, digits and special characters"
	}
}

func VerifyEmail(email string) (bool, string) {
	res, err := verifier.Verify(email)
	if err != nil {
		return false, "verify email address failed"
	}
	if !res.Syntax.Valid {
		return false, "Invalid email syntax"
	}
	return true, "Valid email"
}

func VerifyMobileNumber(phonenumber string) (bool, string) {
	pattern := `^\d{10}$`
	match, _ := regexp.Match(pattern, []byte(phonenumber))
	return match, "Phone number is not valid"
}
