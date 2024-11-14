package rules

import (
	"fmt"
)

func Required(fieldName, _ string, value any) error {
	err := fmt.Errorf("%s is required", fieldName)

	switch value.(type) {
	case string:
		if value.(string) == "" {
			return err
		}
	case int:
		if value.(int) == 0 {
			return err
		}
	}
	// handle other types

	return nil
}
