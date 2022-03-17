package main

import (
	"fmt"
	"os"

	"github.com/timkomip/simple-lexer/lex"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	lexer := lex.NewLexer(file)

	for {
		pos, tok, lit := lexer.Lex()
		if tok == lex.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
	}
}
