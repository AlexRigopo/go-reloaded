package main

import (
	"regexp"
	"strings"
	"unicode"
)

func fixPunctuationSpacing(s string) string {
	// Remove spaces before punctuation
	s = regexp.MustCompile(`\s+([.,!?:;])`).ReplaceAllString(s, `$1`)
	// Ensure exactly one space after punctuation (unless followed by another punctuation or newline)
	s = regexp.MustCompile(`([.,!?:;])\s*`).ReplaceAllString(s, `$1 `)
	// Remove space before newline after punctuation
	s = regexp.MustCompile(`([.,!?:;]) \n`).ReplaceAllString(s, `$1\n`)
	return strings.TrimRight(s, " \t")
}

func fixQuoteSpacing(s string) string {
	// Remove spaces just INSIDE quotes, but do NOT touch spaces outside.
	// " hello " => "hello"
	s = regexp.MustCompile(`(['"])\s+`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`\s+(['"])`).ReplaceAllString(s, `$1`)
	return s
}

// needsAn decides if the word should use "an" (by rough English pronunciation).
func needsAn(word string) bool {
	lower := strings.ToLower(word)
	if lower == "" {
		return false
	}

	// Exceptions: vowel letters but consonant sounds → use "a", not "an"
	// e.g. unicorn (you-), university, user, europe, eulogy, one, once, ubiquitous...
	consonantSoundPrefixes := []string{
		"uni", "use", "usu", "eur", "euro", "eul", "one", "once", "ubiq",
	}
	for _, pre := range consonantSoundPrefixes {
		if strings.HasPrefix(lower, pre) {
			return false
		}
	}

	// Default: starts with a/e/i/o/u → use "an"
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
			// Should be "an"
			if a == "A" {
				return "An " + word
			}
			return "an " + word
		}
		// Should be "a"
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
		an := parts[1]   // "A" or "a" from "An" / "an"
		word := parts[2] // next word

		if needsAn(word) {
			// Should be "an"
			if an == "A" {
				return "An " + word
			}
			return "an " + word
		}
		// Should be "a"
		if an == "A" {
			return "A " + word
		}
		return "a " + word
	})

	return s
}

func fixEllipsis(s string) string {
	// Match patterns like ". . ." or ".   .   . ."
	re := regexp.MustCompile(`(\.\s+){2,}\.`)
	return re.ReplaceAllStringFunc(s, func(m string) string {
		// Count how many dots are in the match
		count := strings.Count(m, ".")
		// Return that many dots without spaces
		return strings.Repeat(".", count)
	})
}

func fixSentenceCapitalization(s string) string {
	runes := []rune(s)
	n := len(runes)
	capNext := true

	for i := 0; i < n; i++ {
		r := runes[i]

		// Apply capitalization when needed
		if capNext && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			capNext = false
			continue
		}

		// Handle sentence boundaries:
		// - '.' ends a sentence, EXCEPT when part of "..."
		// - '?' also ends a sentence
		// - we IGNORE '!' as a sentence end to avoid unwanted caps after quotes,
		//   like: ... "SPARTA!" but no one ...
		if r == '.' {
			// If we see "...", do NOT treat this as sentence end
			if i+2 < n && runes[i+1] == '.' && runes[i+2] == '.' {
				continue
			}
			capNext = true
		} else if r == '?' {
			capNext = true
		}
	}

	return string(runes)
}

func fixSpacing(s string) string {
	// Collapse multiple whitespace into a single space
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	// Remove spaces before punctuation
	s = regexp.MustCompile(`\s+([.,!?;:])`).ReplaceAllString(s, `$1`)
	// NOTE: we NO LONGER remove spaces before quotes here.
	// Quote spacing is handled only inside fixQuoteSpacing.
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
