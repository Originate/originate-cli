package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatNameForREADME(name string) string {
	parts := strings.Split(name, "-")

	lastPart := strings.ToUpper(parts[len(parts)-1])
	firstPart := cases.Title(
		language.English,
		cases.NoLower,
	).String(strings.Join(parts[:len(parts)-1], " "))

	return strings.Join([]string{firstPart, lastPart}, " ")
}
