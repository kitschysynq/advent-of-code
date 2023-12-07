package day5

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type Token int

const (
	EOF = iota
	NUM
	COLON
	IDENT
	TO
	MAP
)

var tokens = []string{
	EOF:  "EOF",
	NUM:  "NUM",
	COLON: "COLON",
	IDENT: "IDENT",
	TO: "TO",
	MAP: "MAP",
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
		case 'C':
			startPos := l.pos
			l.backup()
			lit := l.lexCardID()
			return startPos, CARD, lit
		case '|':
			return l.pos, PIPE, "|"
		default:
			if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexNum()
				return startPos, NUM, lit
			}
			if !unicode.IsSpace(r) {
				panic("unexpected token: " + string(r))
			}
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

func (l *Lexer) lexCardID() string {
	var lit string
	c, err := l.r.ReadString(':')
	if err != nil {
		panic("reading from input: " + err.Error())
	}
	if _, err := fmt.Sscanf(c[:len(c)-1], "Card %s", &lit); err != nil {
		panic(err)
	}
	return lit
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
