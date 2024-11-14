package rules

import (
	"fmt"
	"net/mail"
)

func Email(fieldName, _ string, v any) error {
	errMsg := fmt.Errorf("%s should be a valid email", fieldName)

	emailStr, ok := v.(string)
	if !ok {
		return errMsg
	}

	if _, err := mail.ParseAddress(emailStr); err != nil {
		return errMsg
	}

	return nil
}
