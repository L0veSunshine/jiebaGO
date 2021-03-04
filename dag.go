package jiebaGo

import "math"

type (
	Prob struct {
		Value float64
		idx   int
	}
)

func (t *Tokenizer) getDAG(rs []rune) map[int][]int {
	dag := map[int][]int{}
	var i int
	for k := 0; k < len(rs); k++ {
		var frag string
		var loc []int
		i = k
		frag = string(rs[k])
		for i < len(rs) && t.dict.Has(frag) {
			if t.dict.Freq(frag) != 0 {
				loc = append(loc, i)
			}
			i += 1
			if i+1 > len(rs) {
				frag = string(rs[k:])
			} else {
				frag = string(rs[k : i+1])
			}
		}
		if loc == nil {
			loc = append(loc, k)
		}
		dag[k] = loc
	}
	return dag
}

func (t *Tokenizer) calc(rs []rune, DAG map[int][]int) (route map[int]Prob) {
	n := len(rs)
	route = map[int]Prob{}
	route[n] = Prob{
		Value: 0,
		idx:   0,
	}
	logTotal := math.Log(float64(t.dict.total))
	for idx := n - 1; idx >= 0; idx-- {
		tmp := &Prob{}
		var sw = false
		for _, x := range DAG[idx] {
			var a float64
			a = float64(t.dict.Freq(string(rs[idx : x+1])))
			if a < 1 {
				a = 1
			}
			r := math.Log(a) - logTotal + route[x+1].Value
			if !sw {
				tmp.Value = r
				tmp.idx = x
				sw = true
			}
			if r > tmp.Value {
				tmp.Value = r
				tmp.idx = x
			}
		}
		route[idx] = *tmp
	}
	return route
}
