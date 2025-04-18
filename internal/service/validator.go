package service

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	NonFieldErrors []error
	FieldErrors    map[string]error
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddFieldError(key string, err error) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]error)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = err
	}
}

func (v *Validator) AddNonFieldError(err error) {
	v.NonFieldErrors = append(v.NonFieldErrors, err)
}

func (v *Validator) CheckField(ok bool, key string, err error) {
	if !ok {
		v.AddFieldError(key, err)
	}
}

func (v *Validator) NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func (v *Validator) MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func (v *Validator) MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func (v *Validator) MinValue(value, n int) bool {
	return value >= n
}

func (v *Validator) MaxValue(value, n int) bool {
	return value <= n
}

func (v *Validator) Matches(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}
