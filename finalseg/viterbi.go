package finalseg

func viterbi(obs []rune, start HmmStart, trans, emit HmmTran) (float64, []byte) {
	prevStatus := preState()
	var v = make([]HmmStart, 1)
	var path = Path{}
	for _, y := range HmmState {
		ev, ok := emit[y][obs[0]]
		if !ok {
			ev = MinFloat
		}
		if v[0] == nil {
			v[0] = HmmStart{}
		}
		v[0][y] = start[y] + ev
		path[y] = []byte{y}
	}
	for t := 1; t < len(obs); t++ {
		v = append(v, HmmStart{})
		newPath := Path{}
		var tmpState = make([]byte, 0, 10)
		for _, y := range HmmState {
			emp, ok := emit[y][obs[t]]
			if !ok {
				emp = MinFloat
			}
			var sw = false
			var MaxProb float64
			var state byte
			for _, y0 := range prevStatus[y] {
				tp, ok := trans[y0][rune(y)]
				if !ok {
					tp = MinFloat
				}
				probV := v[t-1][y0] + tp + emp
				if !sw {
					MaxProb = probV
					state = y0
					sw = true
				}
				if probV > MaxProb {
					MaxProb = probV
					state = y0
				} else if probV == MaxProb {
					if y0 > state {
						MaxProb = probV
						state = y0
					}
				}
			}
			v[t][y] = MaxProb
			/*	newPath[y] = append(newPath[y], path[state]...)
				newPath[y] = append(newPath[y], y)  另一种方案
				newPath[y] = tmpState //不能直接赋值 涉及深浅赋值
			*/
			tmpState = append(tmpState, path[state]...)
			tmpState = append(tmpState, y)
			newPath[y] = make([]byte, len(tmpState))
			copy(newPath[y], tmpState)

			tmpState = tmpState[:0]
		}
		path = newPath
	}
	var sw = false
	var maxV float64
	var state byte
	for _, y := range []byte{'E', 'S'} {
		cv := v[len(obs)-1][y]
		if !sw {
			maxV = cv
			state = y
			sw = true
		}
		if cv > maxV {
			maxV = cv
			state = y
		} else if cv == maxV {
			if y > state {
				maxV = cv
				state = y
			}
		}
	}
	return maxV, path[state]
}
