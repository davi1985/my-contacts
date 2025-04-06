package utils

import "regexp"

var uuidRegex = regexp.MustCompile(`(?i)^[0-9a-f]{8}\b-[0-9a-f]{4}\b-[0-9a-f]{4}\b-[0-9a-f]{4}\b-[0-9a-f]{12}$`)

func IsValidUUID(s string) bool {
	return uuidRegex.MatchString(s)
}
