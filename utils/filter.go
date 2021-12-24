package utils

import "strings"

func FilterOrderBy(field, direction, filter string) bool {
	filter = strings.ToLower(filter)
	if strings.Contains(filter, field) && strings.Contains(filter, direction) {
		return true
	}
	return false
}
