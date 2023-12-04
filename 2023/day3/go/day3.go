package day3

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

const (
	EOF = iota
	NUM
	SYM
)

var tokens = []string{
	EOF: "EOF",
	NUM: "NUM",
	SYM: "SYM",
}

func (t Token) String() string {
	return tokens[t]
}

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	pos Position
	r   *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		pos: Position{0, 0},
		r:   bufio.NewReader(r),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}

		l.pos.Column++

		switch r {
		case '\n':
			l.resetPosition()
		case '.':
		default:
			if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexNum()
				return startPos, NUM, lit
			}
			return l.pos, SYM, string(r)
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}

func (l *Lexer) backup() {
	if err := l.r.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Column--
}

func (l *Lexer) lexNum() string {
	var lit string
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.Column++
		if !unicode.IsDigit(r) {
			l.backup()
			return lit
		}

		lit = lit + string(r)
	}
}
