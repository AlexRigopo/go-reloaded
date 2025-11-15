package main

import (
	"regexp"
	"strings"
	"unicode"
)

// Join spaced ellipses like ". . ." or ".  .  . ." into "..." (keeping dot count)
func fixEllipsis(s string) string {
	re := regexp.MustCompile(`(\.\s+){2,}\.`)
	return re.ReplaceAllStringFunc(s, func(m string) string {
		count := strings.Count(m, ".")
		return strings.Repeat(".", count)
	})
}

// Fix spaces around punctuation (inside and outside quotes)
func fixPunctuationSpacing(s string) string {
	// Remove spaces BEFORE punctuation
	s = regexp.MustCompile(`\s+([.,!?;:])`).ReplaceAllString(s, `$1`)
	// Ensure a space AFTER punctuation if next is not space or punctuation
	s = regexp.MustCompile(`([.,!?;:])([^\s.,!?;:])`).ReplaceAllString(s, `$1 $2`)
	return s
}

// Trim spaces just INSIDE quotes: " hello " -> "hello", ' foo ' -> 'foo'
func fixQuoteSpacing(s string) string {
	// Opening quote + spaces
	s = regexp.MustCompile(`(['"])\s+`).ReplaceAllString(s, `$1`)
	// Spaces + closing quote
	s = regexp.MustCompile(`\s+(['"])`).ReplaceAllString(s, `$1`)
	return s
}

// Add spaces OUTSIDE double quotes, without touching apostrophes in words
func fixQuoteOutsideSpacing(s string) string {
	// Space BEFORE opening "
	s = regexp.MustCompile(`([^ \n])(")`).ReplaceAllString(s, `$1 $2`)
	// Space AFTER closing " if followed by non-space, non-punctuation
	s = regexp.MustCompile(`(")([^ \n.,!?;:])`).ReplaceAllString(s, `$1 $2`)
	return s
}

// Decide if the word should use "an" instead of "a" (approximate pronunciation)
func needsAn(word string) bool {
	lower := strings.ToLower(word)
	if lower == "" {
		return false
	}

	// Vowel letters but consonant-sound prefixes → use "a", not "an"
	// e.g. unicorn (you-), user, eulogy, euro, one, hour, etc.
	consonantSoundPrefixes := []string{
		"uni", "use", "usu", "eur", "euro", "eul", "one", "once", "ubiq", "hou",
	}
	for _, pre := range consonantSoundPrefixes {
		if strings.HasPrefix(lower, pre) {
			return false
		}
	}

	first := rune(lower[0])
	return strings.ContainsRune("aeiou", first)
}

func fixArticles(s string) string {
	// Handle "a X"
	reA := regexp.MustCompile(`\b([Aa])\s+([A-Za-z]\w*)`)
	s = reA.ReplaceAllStringFunc(s, func(m string) string {
		parts := reA.FindStringSubmatch(m)
		if len(parts) != 3 {
			return m
		}
		a := parts[1]    // "A" or "a"
		word := parts[2] // next word

		if needsAn(word) {
			if a == "A" {
				return "An " + word
			}
			return "an " + word
		}
		if a == "A" {
			return "A " + word
		}
		return "a " + word
	})

	// Handle "an X"
	reAn := regexp.MustCompile(`\b([Aa])n\s+([A-Za-z]\w*)`)
	s = reAn.ReplaceAllStringFunc(s, func(m string) string {
		parts := reAn.FindStringSubmatch(m)
		if len(parts) != 3 {
			return m
		}
		an := parts[1]   // "A" or "a" from "An"/"an"
		word := parts[2] // next word

		if needsAn(word) {
			if an == "A" {
				return "An " + word
			}
			return "an " + word
		}
		if an == "A" {
			return "A " + word
		}
		return "a " + word
	})

	return s
}

func fixSentenceCapitalization(s string) string {
	runes := []rune(s)
	n := len(runes)
	capNext := true

	for i := 0; i < n; i++ {
		r := runes[i]

		if capNext && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			capNext = false
			continue
		}

		if r == '.' {
			// If it's part of "..." skip triggering capitalization
			if i+2 < n && runes[i+1] == '.' && runes[i+2] == '.' {
				continue
			}
			capNext = true
		} else if r == '!' || r == '?' {
			capNext = true
		}
	}
	return string(runes)
}

func fixSpacing(s string) string {
	// Collapse any run of whitespace to a single space
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	// Remove spaces before punctuation
	s = regexp.MustCompile(`\s+([.,!?;:])`).ReplaceAllString(s, `$1`)
	return strings.TrimSpace(s)
}

func fixCommonGrammar(s string) string {
	replacements := map[*regexp.Regexp]string{
		regexp.MustCompile(`\b[Dd]ont\b`):           "don't",
		regexp.MustCompile(`\b[Cc]ant\b`):           "can't",
		regexp.MustCompile(`\b[Ww]ont\b`):           "won't",
		regexp.MustCompile(`\b[Ii]\s*'?\s*[Mm]\b`):  "I'm",
		regexp.MustCompile(`\b[Ii]\s*'?\s*[Vv]e\b`): "I've",
		regexp.MustCompile(`\b[Dd]oesnt\b`):         "doesn't",
		regexp.MustCompile(`\b[Ii]snt\b`):           "isn't",
		regexp.MustCompile(`\b[Cc]ouldnt\b`):        "couldn't",
		regexp.MustCompile(`\b[Ss]houldnt\b`):       "shouldn't",
		regexp.MustCompile(`\b[Ww]ouldnt\b`):        "wouldn't",
		regexp.MustCompile(`\b[Dd]idnt\b`):          "didn't",
	}
	for re, rep := range replacements {
		s = re.ReplaceAllString(s, rep)
	}
	return s
}
