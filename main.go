package main

import "fmt"

// Sentences returns every way to split s into a space-joined sequence of
// dictionary words. Memoizing on the start index avoids re-solving shared
// suffixes across different prefixes.
func Sentences(s string, dict []string) []string {
	words := make(map[string]bool, len(dict))
	maxLen := 0
	for _, w := range dict {
		words[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}
	memo := make(map[int][]string)
	var solve func(start int) []string
	solve = func(start int) []string {
		if start == len(s) {
			return []string{""}
		}
		if got, ok := memo[start]; ok {
			return got
		}
		out := []string{}
		for end := start + 1; end <= len(s) && end-start <= maxLen; end++ {
			w := s[start:end]
			if !words[w] {
				continue
			}
			for _, rest := range solve(end) {
				if rest == "" {
					out = append(out, w)
				} else {
					out = append(out, w+" "+rest)
				}
			}
		}
		memo[start] = out
		return out
	}
	return solve(0)
}

func main() {
	fmt.Println(Sentences("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
