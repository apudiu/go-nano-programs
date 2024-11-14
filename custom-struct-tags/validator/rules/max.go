package rules

import (
	"fmt"
	"strconv"
)

func Max(fieldName, fieldValue string, v any) error {

	expLen, err := strconv.Atoi(fieldValue)
	if err != nil {
		return fmt.Errorf("max rules value parse failed. It should be like: max=4")
	}

	errMsg := fmt.Errorf("%s should be maximum of %d", fieldName, expLen)

	switch v.(type) {
	case string:
		if len(v.(string)) > expLen {
			return errMsg
		}
	case int:
		if v.(int) > expLen {
			return err
		}
	}

	return nil
}
