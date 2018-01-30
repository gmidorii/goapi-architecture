package form

import "strings"

const dot = "."

func dotParse(s string) []string {
	return strings.Split(s, dot)
}
