package validators

import (
	"errors"
	"fmt"
)

type RulesOfField struct {
	MinSize int
	MaxSize int
}

func (r RulesOfField) Check(field, value string) error {
	cleanedValue := CleanInput(value)

	if len(cleanedValue) > r.MaxSize {
		return fmt.Errorf("*You exceed the maximum character size allowed for the %s field!", field)
	} else if len(cleanedValue) < r.MinSize {
		return fmt.Errorf("*Please enter a valid %s!", field)
	}

	return nil
}

func (r RulesOfField) EmailValidator(field, value string) error {
	if !IsValidEmail(value) {
		return errors.New("*Please enter a valid email address!")
	}
	return nil
}
