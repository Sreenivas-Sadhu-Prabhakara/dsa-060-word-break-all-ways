package main

import (
	"reflect"
	"testing"
)

func TestSentences(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
		want []string
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"},
			[]string{"cat sand dog", "cats and dog"}},
		{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"},
			[]string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, []string{}},
		{"a", []string{"a"}, []string{"a"}},
		{"", []string{"x"}, []string{""}},
	}
	for _, c := range cases {
		if got := Sentences(c.s, c.dict); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Sentences(%q) = %v, want %v", c.s, got, c.want)
		}
	}
}
