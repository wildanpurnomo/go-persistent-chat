package validationLibs

import "regexp"

func IsValidUsername(username string) bool {
	return len(username) > 8
}

func IsValidPassword(password string) bool {
	uppercaseMatcher := regexp.MustCompile(`[A-Z]`)
	lowercaseMatcher := regexp.MustCompile(`[a-z]`)
	numberMatcher := regexp.MustCompile(`[0-9]`)

	return uppercaseMatcher.MatchString(password) &&
		lowercaseMatcher.MatchString(password) &&
		numberMatcher.MatchString(password) &&
		len(password) > 8
}
