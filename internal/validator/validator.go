package validator

import (
	"strings"
	"sync"
)

type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
	mu             sync.RWMutex
}

func New() *Validator {
	return &Validator{
		FieldErrors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddNonFieldError(message string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func (v *Validator) AddFieldError(field, message string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[field]; !exists {
		v.FieldErrors[field] = message
	}
}

func (v *Validator) CheckField(ok bool, field, message string) {
	if !ok {
		v.AddFieldError(field, message)
	}
}

func (v *Validator) GetFieldErrors() map[string]string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.FieldErrors
}

func (v *Validator) GetNonFieldErrors() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.NonFieldErrors
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MinChars(value string, n int) bool {
	return len(strings.TrimSpace(value)) >= n
}

func MaxChars(value string, n int) bool {
	return len(strings.TrimSpace(value)) <= n
}

func Matches(value string, rx string) bool {
	return true
}

func PermittedValue(value string, permittedValues ...string) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}
