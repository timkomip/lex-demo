package lex

import (
	"bufio"
	"io"
	"unicode"
)

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{Line: 1, Column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, err := l.nextRune()

		if err == io.EOF {
			return l.pos, EOF, ""
		} else if err != nil {
			panic(err)
		}

		switch r {
		case '\n':
			l.newLine()
		case ';':
			return l.pos, SEMI, ";"
		case '+':
			return l.pos, ADD, "+"
		case '-':
			return l.pos, SUB, "-"
		case '*':
			if l.peek() == '*' {
				startPos := l.pos
				l.nextRune()
				return startPos, POWER, "**"
			}
			return l.pos, MUL, "*"
		case '/':
			return l.pos, DIV, "/"
		case '=':
			return l.pos, ASSIGN, "="
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, INT, lit
			} else if unicode.IsLetter(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexIdent()
				t := l.lookupIdentType(lit)
				return startPos, t, lit
			} else {
				return l.pos, ILLEGAL, string(r)
			}
		}
	}
}

func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _ := l.nextRune()

		if unicode.IsDigit(r) {
			lit += string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexIdent() string {
	var lit string
	for {
		r, _ := l.nextRune()

		if unicode.IsLetter(r) {
			lit += string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lookupIdentType(lit string) Token {
	if lit == "var" {
		return VAR
	}

	return IDENT
}

func (l *Lexer) nextRune() (rune, error) {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return r, err
		}
	}

	l.pos.Column++
	return r, nil
}

func (l *Lexer) peek() rune {
	nextValue, _, _ := l.reader.ReadRune()
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
	return nextValue
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Column--
}

func (l *Lexer) newLine() {
	l.pos.Line++
	l.pos.Column = 0
}
