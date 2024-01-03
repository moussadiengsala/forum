package validators

import "errors"

var Rules = map[string]RulesOfField{
	"email":      {MinSize: 0, MaxSize: 100},
	"firstName":  {MinSize: 3, MaxSize: 20},
	"title":      {MinSize: 3, MaxSize: 100},
	"username":   {MinSize: 3, MaxSize: 16},
	"password":   {MinSize: 8, MaxSize: 100},
	"lastName":   {MinSize: 3, MaxSize: 20},
	"bio":        {MinSize: 0, MaxSize: 255},
	"content":    {MinSize: 3, MaxSize: 255},
	"author_id":  {MinSize: 1, MaxSize: 16},
	"post_id":    {MinSize: 1, MaxSize: 16},
	"entries_id": {MinSize: 1, MaxSize: 16},
	"action":     {MinSize: 1, MaxSize: 16},
}

func ValidatorService(submittedData map[string]interface{}) error {

	for field, value := range submittedData {
		if rule, ok := Rules[field]; ok {
			var err error

			if field == "email" {
				err = rule.EmailValidator(field, convertToString(value))
			} else if field == "action" {
				if !IsAValidReaction(convertToString(value)) {
					err = errors.New("*Invalid action, Please provide the action you wanna perform!")
				}
			} else {
				err = rule.Check(field, convertToString(value))
			}

			if err != nil {
				return err
			}
		}
	}

	return nil
}
