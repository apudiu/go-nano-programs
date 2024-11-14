package validator

import (
	rules2 "customstructtags/validator/rules"
	"fmt"
	"reflect"
	"strings"
)

const tagName = "validate"

func Validate(v any) (errBag []error) {
	// make sure struct is passed

	vv := reflect.ValueOf(v)
	vt := vv.Type()

	valueKind := vt.Kind().String()

	if valueKind != "struct" {
		return []error{
			fmt.Errorf("expected struct, got %s", valueKind),
		}
	}

	// iterate over all fields and perform validation

	for i := 0; i < vt.NumField(); i++ {

		if err := applyValidation(vt.Field(i), vv.Field(i)); err != nil {
			errBag = append(errBag, err)
		}
	}

	return
}

func applyValidation(t reflect.StructField, v reflect.Value) error {
	// if theres no rule defined in the tag then do nothing
	rulesStr := t.Tag.Get(tagName)
	if rulesStr == "" {
		return nil
	}

	// apply specified validation
	rules := strings.Split(rulesStr, ",")
	for _, ruleSet := range rules {
		ruleName, ruleValue, _ := strings.Cut(ruleSet, "=")

		if ruleName == "" {
			// handle cases where rules has no value like: required
			return fmt.Errorf("invalid rule definition: %s", ruleSet)
		}

		// get required rule
		if err := validateByRule(t.Name, ruleName, ruleValue, v); err != nil {
			return err
		}
	}

	return nil
}

func validateByRule(fieldName, ruleName, ruleValue string, v reflect.Value) error {
	// get required rule
	validationRule, ok := rules2.RuleList[ruleName]
	if !ok {
		return fmt.Errorf("invalid rule: %s", ruleName)
	}

	return validationRule(fieldName, ruleValue, v.Interface())
}
