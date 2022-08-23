package main

import (
	"strings"
	"unicode"
)

func spellchecker(wordlist []string, queries []string) (ans []string) {
	set := map[string]struct{}{}
	ignorelower := map[string]string{}
	ignoreV := map[string]string{}
	for _, w := range wordlist {
		set[w] = struct{}{}
		if _, ok := ignorelower[strings.ToLower(w)]; !ok {
			ignorelower[strings.ToLower(w)] = w
		}
		if _, ok := ignoreV[ignore(w)]; !ok {
			ignoreV[ignore(w)] = w
		}
	}
	for _, q := range queries {
		if _, ok := set[q]; ok {
			ans = append(ans, q)
		} else if v, ok := ignorelower[strings.ToLower(q)]; ok {
			ans = append(ans, v)
		} else if v, ok := ignoreV[ignore(q)]; ok {
			ans = append(ans, v)
		} else {
			ans = append(ans, "")
		}
	}
	return
}

func ignore(word string) string {
	b := []byte(word)
	for i, c := range b {
		c1 := byte(unicode.ToLower(rune(c)))
		if c1 == 'a' || c1 == 'e' || c1 == 'i' || c1 == 'o' || c1 == 'u' {
			c1 = '.'
		}
		b[i] = c1
	}
	return string(b)
}
