package internals

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"learn.zone01dakar.sn/forum-rest-api/lib"
)

var Rules = map[string]RulesOfField{
	"email":      {MinSize: 0, MaxSize: 100},
	"firstname":  {MinSize: 3, MaxSize: 20},
	"title":      {MinSize: 3, MaxSize: 100},
	"username":   {MinSize: 3, MaxSize: 16},
	"password":   {MinSize: 8, MaxSize: 100},
	"lastname":   {MinSize: 3, MaxSize: 20},
	"bio":        {MinSize: 0, MaxSize: 255},
	"content":    {MinSize: 3, MaxSize: 255},
	"author_id":  {MinSize: 1, MaxSize: 16},
	"post_id":    {MinSize: 1, MaxSize: 16},
	"entries_id": {MinSize: 1, MaxSize: 16},
	"action":     {MinSize: 1, MaxSize: 16},
}

const (
	ErrMaxSizeExceeded = "You have exceeded the maximum character size allowed for the %s field!"
	ErrMinSizeRequired = "Please enter a valid %s!"
	ErrInvalidEmail    = "Please enter a valid email address!"
	ErrInvalidAction   = "Invalid action. Please provide a valid action."
)

type RulesOfField struct {
	MinSize int
	MaxSize int
}

func (r RulesOfField) Check(field, value string) error {
	cleanedValue := r.CleanInput(value)

	if len(cleanedValue) > r.MaxSize {
		return fmt.Errorf(ErrMaxSizeExceeded, field)
	} else if len(cleanedValue) < r.MinSize {
		return fmt.Errorf(ErrMinSizeRequired, field)
	}

	return nil
}

func (r RulesOfField) EmailValidator(field, value string) error {
	if !r.IsValidEmail(value) {
		return errors.New(ErrInvalidEmail)
	}
	return nil
}

func (r RulesOfField) CleanInput(input string) string {
	pattern := regexp.MustCompile(`\s+`)
	cleaned := strings.TrimSpace(input)
	cleaned = pattern.ReplaceAllString(cleaned, " ")
	return cleaned
}

// Check if the provided string is a valid email address
func (r RulesOfField) IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func (r RulesOfField) ActionValidator(reaction string) error {
	reactions := []string{"postlikes", "postdislikes", "commentlikes", "commentdislikes"}

	if !slices.Contains(reactions, reaction) {
		return fmt.Errorf(ErrInvalidAction)
	}
	return nil
}

// Validator
type Validators struct{}

func (v Validators) ValidatorService(submittedData interface{}) error {
	var errorsValidator []string
	valueOfData := reflect.ValueOf(submittedData)

	// Ensure the submittedData is a struct
	if valueOfData.Kind() != reflect.Struct {
		// Check if it's a pointer to struct
		if valueOfData.Kind() == reflect.Ptr && valueOfData.Elem().Kind() == reflect.Struct {
			valueOfData = valueOfData.Elem()
		} else {
			return fmt.Errorf("Something went wrong during the validator!!!")
		}
	}

	for i := 0; i < valueOfData.NumField(); i++ {
		field := strings.ToLower(valueOfData.Type().Field(i).Name)
		fieldValue := valueOfData.Field(i).Interface()

		err := v.Validate(field, lib.ConvertToString(fieldValue))
		if err != nil {
			errorsValidator = append(errorsValidator, err.Error())
		}
	}

	if len(errorsValidator) > 0 {
		return fmt.Errorf(strings.Join(errorsValidator, "\n"))
	}

	return nil
}

func (v Validators) Validate(field, value string) error {
	if field == "identifiers" {
		field = v.Identifiers(field, value)
	}

	rule, ok := Rules[field]
	if !ok {
		return nil
	}

	switch field {
	case "email":
		return rule.EmailValidator(field, value)
	case "action":
		return rule.ActionValidator(value)
	default:
		return rule.Check(field, value)
	}
}

func (v Validators) Identifiers(field, value string) string {
	if strings.Contains(value, "@") {
		return "email"
	}
	return "username"
}
