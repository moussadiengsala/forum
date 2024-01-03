package validators

import (
	"regexp"
	"strings"
)

// Clean input by trimming spaces and replacing multiple spaces with a single space
func CleanInput(input string) string {
	pattern := regexp.MustCompile(`\s+`)
	cleaned := strings.TrimSpace(input)
	cleaned = pattern.ReplaceAllString(cleaned, " ")
	return cleaned
}

// Check if the provided string is a valid email address
func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func convertToString(value interface{}) string {
	var stringValue string
	switch v := value.(type) {
	case string:
		stringValue = v
	case *string:
		if v != nil {
			stringValue = *v
		}
	}
	return stringValue
}

func isContains(slices []string, value string) bool {
	for _, s := range slices {
		if strings.Contains(value, s) {
			return true
		}
	}
	return false
}

// check if the reaction action is whether: postlikes, postdislikes, commentlikes, commentdislikes
func IsAValidReaction(value string) bool {
	var reaction = []string{"postlikes", "postdislikes", "commentlikes", "commentdislikes"}
	return isContains(reaction, value)
}
