/*
This package is about Bob.
Bob is a lackadaisical teenager. In conversation, his responses are very limited.
Bob answers 'Sure.' if you ask him a question.
He answers 'Whoa, chill out!' if you yell at him.
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying anything.
He answers 'Whatever.' to anything else.
*/
package bob

import (
	"strings"
	"unicode"
)

func isAlphabetPresent(remark string) bool {
	// This function checks if an alphabet is at least present in the remark given
	for _, character := range remark {
		if unicode.IsLetter(character) {
			return true
		}
	}
	return false
}

func isStringsUpper(remark string) bool {
	// This function compares if remark is all upper case
	return remark == strings.ToUpper(remark)
}

func Hey(remark string) string {
	// This function handles Bob's conversation and returns the right response where necessary.
	remark = strings.TrimSpace(remark)

	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case strings.HasSuffix(remark, "?") && isAlphabetPresent(remark) && isStringsUpper(remark):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case isAlphabetPresent(remark) && isStringsUpper(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
