package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode"
)

type Token int

const (
	EOF = iota
	NL
	NUM
	COLON
	IDENT
	SEEDS
	MAP
)

var tokens = []string{
	EOF:   "EOF",
	NL:    "NL",
	NUM:   "NUM",
	COLON: "COLON",
	IDENT: "IDENT",
	SEEDS: "SEEDS",
	MAP:   "MAP",
}

func (t Token) String() string {
	return tokens[t]
}

type Position struct {
	Line   int
	Column int
}

func (p Position) String() string {
	return fmt.Sprintf("Line %d Column %d", p.Line, p.Column)
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
			return l.pos, NL, "\\n"
		case ':':
			return l.pos, COLON, ":"
		default:
			if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexNum()
				return startPos, NUM, lit
			}
			if unicode.IsLetter(r) {
				startPos := l.pos
				l.backup()
				return l.lexIdent(startPos)
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

func (l *Lexer) lexIdent(start Position) (Position, Token, string) {
	var buf bytes.Buffer
	r, _, err := l.r.ReadRune()
	if err != nil {
		panic(err)
	}
	buf.WriteRune(r)

	for {
		r, _, err := l.r.ReadRune()
		if err == io.EOF {
			break
		}
		if !unicode.IsLetter(r) && r != '-' {
			l.r.UnreadRune()
			break
		}
		_, _ = buf.WriteRune(r)
	}
	switch lit := buf.String(); lit {
	case "seeds":
		return start, SEEDS, lit
	case "map":
		return start, MAP, lit
	default:
		return start, IDENT, lit
	}
}
