package finalseg

import "xuan/jiebaGo/utils"

type FinalSeg struct {
	start           HmmStart
	trans, emit     HmmTran
	hanReg, skipReg *utils.RegEx
	ForceSplit      *utils.Set
}

func NewFinalSeg() *FinalSeg {
	start, trans, emit := loadModel()
	return &FinalSeg{
		start:      start,
		trans:      trans,
		emit:       emit,
		hanReg:     utils.NewRegEx(HanRegExp),
		skipReg:    utils.NewRegEx(SkipReg),
		ForceSplit: utils.NewSet(),
	}
}

func (fs *FinalSeg) cut(rs []rune) (seg []string) {
	_, Pos := viterbi(rs, fs.start, fs.trans, fs.emit)
	begin, next := 0, 0
	for idx, w := range rs {
		pos := Pos[idx]
		switch pos {
		case 'B':
			begin = idx
		case 'E':
			seg = append(seg, string(rs[begin:idx+1]))
			next = idx + 1
		case 'S':
			seg = append(seg, string(w))
			next = idx + 1
		}
	}
	if next < len(rs) {
		seg = append(seg, string(rs[next:]))
	}
	return seg
}

func (fs *FinalSeg) Cut(sentence string) (res []string) {
	res = make([]string, 0, 10)
	blocks := fs.hanReg.SplitAll(sentence)
	for _, blk := range blocks {
		if fs.hanReg.MatchString(blk) {
			words := fs.cut([]rune(blk))
			for _, word := range words {
				if !fs.ForceSplit.Has(word) {
					res = append(res, word)
				} else {
					for _, w := range word {
						res = append(res, string(w))
					}
				}
			}
		} else {
			tmp := fs.skipReg.SplitAll(blk)
			for _, seg := range tmp {
				if seg != "" {
					res = append(res, seg)
				}
			}
		}
	}
	return res
}
