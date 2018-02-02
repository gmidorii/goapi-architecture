package form

import "strings"

const dot = "."

type Former interface {
	valid() (bool, error)
}

func dotParse(s string) []string {
	return strings.Split(s, dot)
}
