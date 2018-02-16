package tic_tac_trinary

import (
	"fmt"
	"testing"
)

func TestLetterMap(t *testing.T) {
	tests := []struct {
		in string
		want []string
	}{
		{"pfy", []string{"OXO", " X ", "XXO"}},
	}

	for _, test := range tests {
		got, err := TranslateToTrits(test.in)
		if err != nil || fmt.Sprintf("%#v", got) != fmt.Sprintf("%#v", test.want) {
			t.Errorf("failed to convert %s, actual=%v, want=%v, err=%s", test.in, got, test.want, err)
		}
	}
}
