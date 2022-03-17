package lex

type Token int

const (
	EOF = iota
	ILLEGAL
	IDENT
	INT
	SEMI // ;

	// Infix ops
	ADD   // +
	SUB   // -
	MUL   // *
	DIV   // /
	POWER // **

	ASSIGN // =

	VAR // var
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",

	ADD:   "+",
	SUB:   "-",
	MUL:   "*",
	DIV:   "/",
	POWER: "**",

	ASSIGN: "=",

	VAR: "VAR",
}

func (t Token) String() string {
	return tokens[t]
}
