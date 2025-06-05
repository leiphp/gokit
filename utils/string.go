package utils

import "strings"

func SnakeToCamel(input string) string {
	parts := strings.Split(input, "_")
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
