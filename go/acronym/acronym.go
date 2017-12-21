/*
This package helps to generate some jargon by writing a program that
converts a long name like Portable Network Graphics to its acronym (PNG).
*/

package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Replace all occurences of dashes in the string with space
	s = strings.Replace(s, "-", " ", -1)

	var acronym string
	stringList := strings.Split(s, " ")
	for _, word := range stringList {
		acronym += word[:1]
	}
	return strings.ToUpper(acronym)
}
