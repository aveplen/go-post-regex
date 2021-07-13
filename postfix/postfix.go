package postfix

import (
	"strings"
)

func (r Regex) New(s string) Regex {
	return Regex(s).ExplicitConcat()
}

func (r Regex) ExplicitConcat() (res Regex) {
	for i, token := range r {
		res = res + Regex(token)
		if token == '(' || token == '|' {
			continue
		}
		if i+1 < len(r) {
			lookahead := r[i+1]
			if strings.Contains("*?+|)", string(lookahead)) {
				continue
			}
			res = res + "."
		}
	}
	return res
}
