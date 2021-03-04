package utils

import "regexp"

type RegEx struct {
	*regexp.Regexp
}

func NewRegEx(expr string) *RegEx {
	reg := regexp.MustCompile(expr)
	return &RegEx{reg}
}

func (r *RegEx) SplitAll(s string) (frags []string) {
	loc := r.FindAllStringIndex(s, -1)
	frags = make([]string, 0, 10)
	if len(loc) == 0 {
		frags = append(frags, s)
		return frags
	}
	var last = -1
	for idx := 0; idx < len(loc); {
		var part = loc[idx]
		if last != -1 {
			if last != part[0] {
				frags = append(frags, s[last:part[0]])
			}
		} else {
			if part[0] != 0 {
				frags = append(frags, s[0:part[0]])
			}
		}
		frags = append(frags, s[part[0]:part[1]])
		last = part[1]
		idx += 1
		if idx == len(loc) && last != len(s) {
			frags = append(frags, s[last:])
		}
	}
	return frags
}
