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

	var b strings.Builder
	for _, tk := range out {
		b.WriteString(tk.text)
	}
	res := b.String()

	// Post-processing pipeline (in this exact order)
	res = fixEllipsis(res)
	res = fixCommonGrammar(res)
	res = fixPunctuationSpacing(res)
	res = fixSpacing(res)
	res = normalizeQuotes(res)
	res = fixArticles(res)
	res = fixSentenceCapitalization(res)

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
