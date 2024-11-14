package rules

import (
	"fmt"
	"strconv"
)

func Min(fieldName, fieldValue string, v any) error {

	expLen, err := strconv.Atoi(fieldValue)
	if err != nil {
		return fmt.Errorf("min rules value parse failed. It should be like: min=4")
	}

	errMsg := fmt.Errorf("%s is required to be at-least %d", fieldName, expLen)

	switch v.(type) {
	case string:
		if len(v.(string)) < expLen {
			return errMsg
		}
	case int:
		if v.(int) < expLen {
			return err
		}
	}

	return nil
}
