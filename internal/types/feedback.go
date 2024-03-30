package types

import (
	"strings"
)

type FeedBack struct {
	Name        string
	Email       string
	Subject     string
	Description string
	Date        string
}

func (f *FeedBack) Validate() ([]string, bool) {
	var errors []string
	var hasError bool
	f.Name = strings.TrimRight(f.Name, " ")
	f.Name = strings.TrimLeft(f.Name, " ")
	if len(f.Name) < 5 || len(f.Name) > 15 {
		errors = append(errors, "invalid name length")
		hasError = true
	}
	if !EmailPattern.MatchString(f.Email) {
		errors = append(errors, "invalid email")
		hasError = true
	}
	if len(f.Description) < 15 || len(f.Description) > 280 {
		errors = append(errors, "invalid description length 15 - 280 ")
		hasError = true
	}
	return errors, hasError
}
