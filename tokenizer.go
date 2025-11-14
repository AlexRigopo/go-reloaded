package main

import (
	"regexp"
	"strings"
	"unicode"
)

type Token struct {
	kind string
	text string
}

var (
	reMarker     = regexp.MustCompile(`^\((?i:(up|low|cap))(?:\s*,\s*(\d+))?\)$|^\((?i:(hex|bin))\)$`)
	reWord       = regexp.MustCompile(`^[A-Za-z]+$`)
	reNumber     = regexp.MustCompile(`^[0-9A-Fa-f]+$`)
	reWhitespace = regexp.MustCompile(`\s+`)
	rePunct      = regexp.MustCompile(`^[.,!?:;]$`)
)

func tokenize(s string) []Token {
	var tokens []Token
	runes := []rune(s)
	n := len(runes)
	i := 0

	for i < n {
		r := runes[i]
		switch {
		case r == '\n' || r == '\r':
			tokens = append(tokens, Token{"newline", string(r)})
			i++
		case unicode.IsSpace(r):
			j := i + 1
			for j < n && unicode.IsSpace(runes[j]) && runes[j] != '\n' && runes[j] != '\r' {
				j++
			}
			tokens = append(tokens, Token{"space", string(runes[i:j])})
			i = j
		case strings.ContainsRune(".,!?:;", r):
			tokens = append(tokens, Token{"punct", string(r)})
			i++
		case r == '(':
			j := i + 1
			for j < n && runes[j] != ')' {
				j++
			}
			if j < n && runes[j] == ')' {
				candidate := string(runes[i : j+1])
				if reMarker.MatchString(candidate) {
					tokens = append(tokens, Token{"marker", candidate})
					i = j + 1
					continue
				}
			}
			tokens = append(tokens, Token{"other", string(r)})
			i++
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			j := i + 1
			for j < n && (unicode.IsLetter(runes[j]) || unicode.IsNumber(runes[j])) {
				j++
			}
			lit := string(runes[i:j])
			tokens = append(tokens, Token{classifyLiteral(lit), lit})
			i = j
		default:
			tokens = append(tokens, Token{"other", string(r)})
			i++
		}
	}
	return tokens
}

func classifyLiteral(lit string) string {
	if reWord.MatchString(lit) {
		return "word"
	}
	if reNumber.MatchString(lit) {
		return "number"
	}
	return "other"
}
