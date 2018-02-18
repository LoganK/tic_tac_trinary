package tic_tac_trinary

import (
	"fmt"
	"strings"
)

var (
	letterMap map[rune]string
)

func init() {
	letterMap = make(map[rune]string)

	fromRune := func(c rune) string {
		trit := map[int32]rune{
			// 0 is 0 lines, 1 is 1 line, 2 is 2 lines
			0: ' ',
			1: 'O',
			2: 'X',
		}
		res := ""
		ci := int32(c) - int32('A') + 1
		for i := 0; i < 3; i++ {
			res = string(trit[ci%3]) + res
			ci /= 3
		}
		return res
	}

	// We use 1-based letters and use space as the 0/null character.
	letterMap[' '] = "   "
	for c := 'A'; c <= 'Z'; c++ {
		letterMap[c] = fromRune(c)
	}
}

// Convert the given message into a series of trit triples.
func TranslateToTrits(msg string) ([]string, error) {
	msg = strings.ToUpper(msg)
	var res []string
	for _, r := range msg {
		trits, ok := letterMap[r]
		if !ok {
			return nil, fmt.Errorf("invalid letter: %c", r)
		}
		res = append(res, trits)
	}

	return res, nil
}
