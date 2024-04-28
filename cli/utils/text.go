package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.AmericanEnglish)

func TitleCase(s string) string {
	return titleCaser.String(s)
}

func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if i >= 0 {
			parts[i] = titleCaser.String(parts[i])
		}
	}
	return strings.Join(parts, "")
}
