package main

import (
	"strings"
	"unicode"
)

func transform(s string) string {
	tokens := tokenize(s)

	out := make([]Token, 0, len(tokens))
	for _, tk := range tokens {
		if tk.kind == "marker" {
			applyMarker(&out, tk.text)
			continue
		}
		out = append(out, tk)
	}

	// Deduplicate consecutive identical tokens
	cleaned := make([]Token, 0, len(out))
	for i, tk := range out {
		if i > 0 && tk.text == out[i-1].text && tk.kind == out[i-1].kind {
			continue
		}
		cleaned = append(cleaned, tk)
	}
	out = cleaned

	var b strings.Builder
	for _, tk := range out {
		b.WriteString(tk.text)
	}
	res := b.String()

	res = fixPunctuationSpacing(res)
	res = fixQuoteSpacing(res)
	res = fixArticles(res)
	res = fixSentenceCapitalization(res)
	res = fixSpacing(res)
	res = fixCommonGrammar(res)

	return res
}

func capitalize(w string) string {
	if w == "" {
		return w
	}
	runes := []rune(strings.ToLower(w))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func reverseStrings(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
