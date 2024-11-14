package rules

type RuleFunc func(fieldName, ruleValue string, value any) error

var RuleList = map[string]RuleFunc{
	"required": Required,
	"min":      Min,
	"max":      Max,
	"email":    Email,
}
