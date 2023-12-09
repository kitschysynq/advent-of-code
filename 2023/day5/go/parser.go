package day5

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	l   *Lexer
	buf struct {
		p   Position
		t   Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{l: NewLexer(r)}
}

func (p *Parser) lex() (Position, Token, string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.p, p.buf.t, p.buf.lit
	}
	pos, tok, lit := p.l.Lex()
	p.buf.p, p.buf.t, p.buf.lit = pos, tok, lit
	return pos, tok, lit
}

func (p *Parser) Parse() (*Map, error) {
	m := new(Map)

	_, tok, lit := p.lex()
	if tok != NL {
		return nil, fmt.Errorf("expected newline, found %q", lit)
	}

	for {
		_, tok, lit := p.lex()
		if tok == EOF {
			break
		}
		if tok != IDENT {
			return nil, fmt.Errorf("expected IDENT, found %q", lit)
		}
		p.unscan()
		i, err := p.parseItemMap()
		if err != nil {
			return nil, err
		}
		m.ItemMaps = append(m.ItemMaps, i)
	}

	return m, nil
}

func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) parseItemMap() (*ItemMap, error) {
	i := new(ItemMap)

	pos, _, lit := p.lex()
	parts := strings.Split(lit, "-to-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid map ident %q found at %s", lit, pos)
	}
	i.From, i.To = parts[0], parts[1]

	if err := p.consume(MAP, COLON, NL); err != nil {
		return nil, err
	}

	for {
		pos, tok, lit := p.lex()
		if tok == NL || tok == EOF {
			return i, nil
		}
		p.unscan()
		nums, err := p.parseNumList()
		if err != nil {
			return nil, err
		}
		if len(nums) != 3 {
			return nil, fmt.Errorf("invalid range definition %q at %s", lit, pos)
		}
		i.Ranges = append(i.Ranges, Range{
			To:   nums[0],
			From: nums[1],
			Len:  nums[2],
		})
	}

	return i, nil
}

func (p *Parser) consume(toks ...Token) error {
	for _, t := range toks {
		pos, tok, lit := p.lex()
		if t != tok {
			return fmt.Errorf(`expected %q, found %q at %s`, t, lit, pos)
		}
	}
	return nil
}

// parseNumList parses a sequence of numbers terminated by a newline
func (p *Parser) parseNumList() ([]int, error) {
	var nums []int
	for {
		pos, tok, lit := p.lex()
		if tok == NL {
			break
		}
		if tok != NUM {
			return nil, fmt.Errorf(`expected number, found %q at %s`, lit, pos)
		}
		n, err := strconv.Atoi(lit)
		if err != nil {
			return nil, fmt.Errorf("invalid number literal %q at %s", lit, pos)
		}
		nums = append(nums, n)
	}
	return nums, nil
}
