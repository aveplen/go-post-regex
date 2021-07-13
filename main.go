package main

import (
	"fmt"

	match "github.com/aveplen/go-post-regex/checker"
	compiler "github.com/aveplen/go-post-regex/compiler"
	p "github.com/aveplen/go-post-regex/postfix"
)

const r p.Regex = "a"

func main() {
	postfixReg := r.ExplicitConcat().ToPostfix()
	println(postfixReg)
	nfa := compiler.Compile(postfixReg)
	fmt.Println(match.Match(nfa, "a"))
}
