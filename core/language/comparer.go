package language

import "strings"

func Equal(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}
