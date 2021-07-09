package main

import (
	"fmt"

	stack "github.com/aveplen/go-weak-stack"
)

type regex string

const r regex = "https?+:(//)*twitch\\.tv/(verni_shavermy)"

func (r regex) toRuneArr() []rune {
	res := make([]rune, len(r))
	for i, tok := range r {
		res[i] = tok
	}
	return res
}

func fromRuneArr(runeArr []rune) regex {
	return regex(string(runeArr))
}

func explicitConcat(r regex) regex {
	runeArr := r.toRuneArr()
	temp := make([]rune, len(runeArr))
	for i, token := range runeArr {
		temp = append(temp, token)
		if token == '(' || token == '|' {
			continue
		}
		if i+1 < len(runeArr) {
			lookahead := runeArr[i+1]
			if lookahead == '*' || lookahead == '?' || lookahead == '+' || lookahead == '|' || lookahead == ')' {
				continue
			}
			temp = append(temp, '.')
		}
	}
	res := fromRuneArr(temp)
	return res
}

type postRegex regex

func (r regex) toPostfix() postRegex {
	// not implemented
	return postRegex("")
}

type stringerInt int

func (strInt stringerInt) String() string {
	return fmt.Sprint(int(strInt))
}

func main() {
	fmt.Println("Before execution I got:", r)
	fmt.Println("After execution:", explicitConcat(r))

	st := stack.New()
	var i stringerInt = 0
	for ; i < 10; i++ {
		st.Push(i)
		fmt.Println(st)
	}

	for j := 0; j < 8; j++ {
		x, _ := st.Pop()
		fmt.Println(st, "   ", x)
	}
}
