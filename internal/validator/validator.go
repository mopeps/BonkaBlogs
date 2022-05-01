package validator

import "regexp"

// taken from https://html.spec.whatwg.org/#valid-e-mail-address.
var (
	EmalRGEX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\. [a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]string
}

// A Validator type initializer
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError ,as it name says, it adds an error to a Validator's map
// ( As long as no entry with the key exists)
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In returns true if a specific value is in a list of strings
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all string values in a slice are unique
func isSet(values []string) bool {
	isValueUnique := make(map[string]bool)

	for _, value := range values {
		isValueUnique[value] = true
	}

	return len(values) == len(isValueUnique)
}

func isUnique[T comparable](value T, elems []T) bool {
	count := 0
	for _, elem := range elems {
		if elem == value {
			count++
		}
	}
	return count == 1
}
