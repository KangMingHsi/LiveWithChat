package chat_subsystem

import (
	"regexp"
	"strings"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// Convert camel case to underscore case
func Underscore(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}