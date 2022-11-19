package autusecase

import (
	"eventool/internal/pkg/domain"
	"net/mail"
	"strings"
	"unicode"
)

const minPasswordLen = 8
const minPhoneLen = 5
const maxPhoneLen = 20

func validateEmail(address string) error {
	if _, err := mail.ParseAddress(address); err != nil {
		return domain.Err.ErrObj.InvalidEmail
	}
	return nil
}

func ValidateUsername(username string) error {
	for _, char := range username {
		if !(unicode.IsLetter(char) || unicode.Is(unicode.Cyrillic, char)) {
			return domain.Err.ErrObj.InvalidUsername
		}
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < minPasswordLen {
		return domain.Err.ErrObj.InvalidPassword
	}
	return nil
}

func validatePhoneNumber(phoneNumber string) error {
	for _, char := range phoneNumber {
		if !(unicode.IsDigit(char)) {
			return domain.Err.ErrObj.InvalidPhoneNumber
		}
	}
	if len(phoneNumber) < minPhoneLen  || len(phoneNumber) < maxPhoneLen{
		return domain.Err.ErrObj.InvalidPhoneNumber
	}
	return nil
}

func trimCredentials(email *string, username *string, password *string, repeatPassword *string, phoneNumber *string) {
	*email = strings.Trim(*email, " ")
	*username = strings.Trim(*username, " ")
	*password = strings.Trim(*password, " ")
	*repeatPassword = strings.Trim(*repeatPassword, " ")
	*phoneNumber = strings.Trim(*phoneNumber, " ")
}
