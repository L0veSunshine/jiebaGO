package jiebaGo

import (
	"encoding/json"
	"fmt"
	"github.com/jiebaGo/utils"
	"strconv"
	"unicode/utf8"
)

type (
	Token struct {
		Word  string
		Start int
		End   int
		Pos   int
	}
)

func genFlagSet() *utils.Set {
	s := utils.NewSet()
	s.Add("\r\n")
	s.Add("\n")
	s.Add("\r")
	s.Add("\t")
	s.Add(" ")
	s.Add("\u3000")
	return s
}

type (
	Tokens []Token
)

func (t *Tokenizer) Tokenize(article string, search, hmm bool) (res Tokens) {
	if t.specFlag == nil {
		t.specFlag = genFlagSet()
	}
	res = make([]Token, 0, 50)
	var start = 0
	if !search {
		for p, word := range t.Cut(article, hmm) {
			width := utf8.RuneCountInString(word)
			if t.specFlag.Has(word) {
				word = strconv.Quote(word)
			}
			token := &Token{
				Word:  word,
				Start: start,
				End:   start + width,
				Pos:   p,
			}
			res = append(res, *token)
			start += width
		}
	} else {
		for p, word := range t.Cut(article, hmm) {
			RWord := []rune(word)
			width := len(RWord)
			if t.specFlag.Has(word) {
				word = strconv.Quote(word)
			}
			if width > 2 {
				for i := 0; i < width-1; i++ {
					Gram2 := string(RWord[i : i+2])
					if t.dict.Freq(Gram2) != 0 {
						token := &Token{
							Word:  Gram2,
							Start: start + i,
							End:   start + i + 2,
							Pos:   p,
						}
						res = append(res, *token)
					}
				}
			}
			if width > 3 {
				for i := 0; i < width-2; i++ {
					Gram3 := string(RWord[i : i+3])
					if t.dict.Freq(Gram3) != 0 {
						token := &Token{
							Word:  Gram3,
							Start: start + i,
							End:   start + i + 3,
							Pos:   p,
						}
						res = append(res, *token)
					}
				}
			}
			token := Token{
				Word:  word,
				Start: start,
				End:   start + width,
				Pos:   p,
			}
			res = append(res, token)
			start += width
		}
	}
	return res
}

func (ts *Tokens) ToJson() string {
	bytes, err := json.Marshal(ts)
	if err != nil {
		fmt.Println(err)
	}
	return utils.ByteToStr(bytes)
}
