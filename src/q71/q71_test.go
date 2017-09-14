package q71

import "testing"

func TestIsStopWords(t *testing.T) {
	s := "a"
	if !IsStopWords(s) {
		t.Errorf("your input: `%s`", s)
	}

}
