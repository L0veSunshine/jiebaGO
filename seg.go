package jiebaGo

import (
	"math"
	"unicode/utf8"
	"xuan/jiebaGo/finalseg"
	"xuan/jiebaGo/utils"
)

const (
	ReEng     = "[a-zA-Z0-9]"
	RegExp    = "([\u4e00-\u9fd5a-zA-Z0-9+#&\\._%\\-]+)"
	ReSkipExp = "((\r\n)|\\s|\u3000)"
)

type Tokenizer struct {
	dict                 *PrefixDict
	hmmSeg               *finalseg.FinalSeg
	reEng, reHan, reSkip *utils.RegEx
	specFlag             *utils.Set
}

func NewTokenizer() *Tokenizer {
	pd := load()
	fs := finalseg.NewFinalSeg()
	return &Tokenizer{
		dict:   pd,
		hmmSeg: fs,
	}
}

func (t *Tokenizer) AddWord(word string, freq int) {
	var sFreq int
	if freq > 0 {
		sFreq = freq
	} else {
		sFreq = t.suggestFreq(word)
	}
	t.dict.addWord(word, sFreq)
	t.dict.addRecur(word)
	if sFreq == 0 {
		t.hmmSeg.ForceSplit.Add(word)
	}
}

func (t *Tokenizer) suggestFreq(segment string) int {
	total := t.dict.total
	var freq float64
	freq = 1
	word := segment
	for _, seg := range t.cutNoHmm(segment) {
		freq *= float64(t.dict.DefaultFreq(seg, 1)) / float64(total)
	}
	freq = math.Max(freq*float64(total)+1, float64(t.dict.DefaultFreq(word, 1)))
	return int(freq)
}

func (t *Tokenizer) cutNoHmm(sentence string) (res []string) {
	if t.reEng == nil {
		t.reEng = utils.NewRegEx(ReEng)
	}
	res = make([]string, 0, 10)
	SSent := utils.StrToRune(sentence)
	dag := t.getDAG(SSent)
	route := t.calc(SSent, dag)
	var x = 0
	n := len(SSent)
	buf := ""
	for x < n {
		y := route[x].idx + 1
		LWord := string(SSent[x:y])
		if t.reEng.MatchString(LWord) && len(SSent[x:y]) == 1 {
			buf += LWord
			x = y
		} else {
			if buf != "" {
				res = append(res, buf)
				buf = ""
			}
			res = append(res, LWord)
			x = y
		}
	}
	if buf != "" {
		res = append(res, buf)
		buf = ""
	}
	return res
}

func (t *Tokenizer) cutDag(sentence string) (seg []string) {
	seg = make([]string, 0, 10)
	SSent := utils.StrToRune(sentence)
	dag := t.getDAG(SSent)
	route := t.calc(SSent, dag)
	var x = 0
	var buf = ""
	n := len(SSent)
	for x < n {
		y := route[x].idx + 1
		LWord := string(SSent[x:y])
		if y-x == 1 {
			buf += LWord
		} else {
			if buf != "" {
				if utf8.RuneCountInString(buf) == 1 {
					seg = append(seg, buf)
					buf = ""
				} else {
					if t.dict.Freq(buf) == 0 {
						recognized := t.hmmSeg.Cut(buf)
						seg = append(seg, recognized...)
					} else {
						for _, elem := range buf {
							seg = append(seg, string(elem))
						}
					}
					buf = ""
				}
			}
			seg = append(seg, LWord)
		}
		x = y
	}
	if buf != "" {
		if utf8.RuneCountInString(buf) == 1 {
			seg = append(seg, buf)
		} else if t.dict.Freq(buf) == 0 {
			recognized := t.hmmSeg.Cut(buf)
			seg = append(seg, recognized...)
		} else {
			for _, elem := range buf {
				seg = append(seg, string(elem))
			}
		}
	}
	return seg
}

func (t *Tokenizer) Cut(article string, hmm bool) (seg []string) {
	if t.reHan == nil && t.reSkip == nil {
		t.reHan = utils.NewRegEx(RegExp)
		t.reSkip = utils.NewRegEx(ReSkipExp)
	}
	seg = make([]string, 0, 50)
	blocks := t.reHan.SplitAll(article)
	var s []string
	var cuter func(string) []string
	if hmm {
		cuter = t.cutDag
	} else {
		cuter = t.cutNoHmm
	}
	for _, b := range blocks {
		if b == "" {
			continue
		}
		if t.reHan.MatchString(b) {
			s = cuter(b)
			seg = append(seg, s...)
		} else {
			tmp := t.reSkip.SplitAll(b)
			for _, x := range tmp {
				if t.reSkip.MatchString(x) {
					seg = append(seg, x)
				} else {
					for _, segPart := range x {
						seg = append(seg, string(segPart))
					}
				}
			}
		}
	}
	return seg
}

func (t *Tokenizer) CutSearch(article string, hmm bool) (res []string) {
	res = make([]string, 0, 50)
	seg := t.Cut(article, hmm)
	for _, word := range seg {
		RWord := utils.StrToRune(word)
		length := len(RWord)
		if length > 2 {
			for i := 0; i < length-1; i++ {
				Gram2 := string(RWord[i : i+2])
				if t.dict.Freq(Gram2) != 0 {
					res = append(res, Gram2)
				}
			}
		}
		if length > 3 {
			for i := 0; i < length-2; i++ {
				Gram3 := string(RWord[i : i+3])
				if t.dict.Freq(Gram3) != 0 {
					res = append(res, Gram3)
				}
			}
		}
		res = append(res, word)
	}
	return res
}
