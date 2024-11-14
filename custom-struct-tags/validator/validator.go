package validator

import (
	"fmt"
	"reflect"
	"strings"
)

const tagName = "validate"

func Validate(v any) (errBag []error) {
	// make sure struct is passed

	t := reflect.TypeOf(v)

	valueKind := t.Kind().String()

	if valueKind != "struct" {
		return []error{
			fmt.Errorf("expected struct, got %s", valueKind),
		}
	}

	// iterate over all fields and perform validation

	for i := 0; i < t.NumField(); i++ {
		if err := applyValidation(t.Field(i)); err != nil {
			errBag = append(errBag, err)
		}
	}

	return
}

func applyValidation(f reflect.StructField) error {
	// find validation tag

	rulesStr := f.Tag.Get(tagName)
	if rulesStr == "" {
		return nil
	}

	// apply specified validation

	for _, ruleSet := range strings.Split(rulesStr, ",") {
		fmt.Printf("%s - %#v \n", f.Name, ruleSet)
	}

	return nil
}
