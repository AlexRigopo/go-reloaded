package main

import (
	"fmt"
	"strconv"
	"strings"
)

func applyMarker(out *[]Token, marker string) {
	m := reMarker.FindStringSubmatch(marker)
	if m == nil {
		return
	}
	if m[1] != "" {
		mode := strings.ToLower(m[1])
		N := 1
		if m[2] != "" {
			if v, err := strconv.Atoi(m[2]); err == nil && v > 0 {
				N = v
			}
		}
		applyCaseToPreviousWords(out, mode, N)
		return
	}
	if m[3] != "" {
		baseWord := strings.ToLower(m[3])
		convertPrevNumber(out, baseWord)
	}
}

func applyCaseToPreviousWords(out *[]Token, mode string, N int) {
	if N <= 0 {
		return
	}
	count := 0
	affected := []string{}
	for i := len(*out) - 1; i >= 0 && count < N; i-- {
		tk := (*out)[i]
		if tk.kind == "word" {
			switch mode {
			case "up":
				(*out)[i].text = strings.ToUpper(tk.text)
			case "low":
				(*out)[i].text = strings.ToLower(tk.text)
			case "cap":
				(*out)[i].text = capitalize(tk.text)
			}
			affected = append(affected, (*out)[i].text)
			count++
		}
	}
	if debug {
		fmt.Printf("[DEBUG] Marker (%s,%d) affected: %v\n", mode, N, reverseStrings(affected))
	}
}

func convertPrevNumber(out *[]Token, baseWord string) {
	for i := len(*out) - 1; i >= 0; i-- {
		if (*out)[i].kind == "number" {
			lit := (*out)[i].text
			base := 10
			if baseWord == "hex" {
				base = 16
			} else if baseWord == "bin" {
				base = 2
			}
			if v, err := strconv.ParseInt(lit, base, 64); err == nil {
				(*out)[i].text = strconv.FormatInt(v, 10)
				if debug {
					fmt.Printf("[DEBUG] Converted %s (%s) → %s\n", lit, baseWord, (*out)[i].text)
				}
			}
			break
		}
	}
}
