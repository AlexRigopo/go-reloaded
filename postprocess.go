package main

import (
	"regexp"
	"strings"
	"unicode"
)

func fixPunctuationSpacing(s string) string {
	s = regexp.MustCompile(`\s+([.,!?:;])`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`([.,!?:;])\s*`).ReplaceAllString(s, `$1 `)
	s = regexp.MustCompile(`([.,!?:;]) \n`).ReplaceAllString(s, `$1\n`)
	return strings.TrimRight(s, " \t")
}

func fixQuoteSpacing(s string) string {
	s = regexp.MustCompile(`(['"])\s+`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`\s+(['"])`).ReplaceAllString(s, `$1`)
	return s
}

func fixArticles(s string) string {
	// a → an (before vowel)
	reA := regexp.MustCompile(`\b([Aa])\s+([AEIOUaeiou]\w*)`)
	s = reA.ReplaceAllStringFunc(s, func(match string) string {
		parts := reA.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}
		a := parts[1]
		word := parts[2]

		if a == "A" {
			return "An " + word
		}
		return "an " + word
	})

	// an → a (before consonant)
	reAn := regexp.MustCompile(`\b([Aa])n\s+([^AEIOUaeiou]\w*)`)
	s = reAn.ReplaceAllStringFunc(s, func(match string) string {
		parts := reAn.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}
		an := parts[1]
		word := parts[2]

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
		if strings.ContainsRune(".!?", r) {
			capNext = true
		}
	}
	return string(runes)
}

func fixSpacing(s string) string {
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	s = regexp.MustCompile(`\s+([.,!?;:])`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`\s+(['"])`).ReplaceAllString(s, `$1`)
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
