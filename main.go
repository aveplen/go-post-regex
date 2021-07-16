package main

import (
	"fmt"
	"os"

	match "github.com/aveplen/go-post-regex/checker"
	compiler "github.com/aveplen/go-post-regex/compiler"
	p "github.com/aveplen/go-post-regex/postfix"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: match 'regex' 'str'\n")
	os.Exit(2)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	r := p.Regex(os.Args[1])
	str := os.Args[2]
	postfixReg := r.ExplicitConcat().ToPostfix()
	nfa := compiler.Compile(postfixReg)
	fmt.Println(match.Match(nfa, str))
}
