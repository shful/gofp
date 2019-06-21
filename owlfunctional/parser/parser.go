// parser takes the concept from:
// https://blog.gopheracademy.com/advent-2014/parsers-lexers/
package parser

import (
	"fmt"
	"io"
	"log"
)

var TokenLog bool // print all parsed tokens

// Parser represents a parser.
type Parser struct {
	s          *Scanner
	pBal       int // parentheses balance starts with 0
	lineNo     int // >= 0
	sourceName string

	// currentLineHead is the line from beginning to, including, the literal starting at colNo
	currentLineHead string

	buf struct {
		tok Token          // last read token
		lit string         // last read literal
		n   int            // buffer size (max=1)
		pos ParserPosition // position where lit comes after
	}
}

// NewParser returns a new instance of Parser.
// sourceName identifies what is parsed.
// The sourceName is shown in error messages. It is never interpreted and must not fulfil any format. Probably, you provide a filename here.
func NewParser(r io.Reader, sourceName string) *Parser {
	return &Parser{s: NewScanner(r), sourceName: sourceName, lineNo: 0} // lineNo internally starts with 0
}

// forwardPos updates colNo and lineNo and currentLineHead,
// after the given token was read
func (p *Parser) forwardPos(tok Token, lit string) {
	if tok == EOL {
		p.lineNo++
		p.currentLineHead = ""
	} else {
		// build current line, with the literal beginning at pos as mostright element.
		p.currentLineHead += lit
	}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) Scan() (tok Token, lit string, pos ParserPosition) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		p.forwardPos(p.buf.tok, p.buf.lit)
		if TokenLog {
			log.Println("Re-read", DescribeToklit(p.buf.tok, p.buf.lit), "after", p.buf.pos)
		}
		return p.buf.tok, p.buf.lit, p.buf.pos
	}

	// pos is what we return and buffer. It is where we are before(!) reading the next literal.
	pos = p.Pos()

	// read the next token from the scanner.
	tok, lit = p.s.scan()
	p.forwardPos(tok, lit)

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit
	p.buf.pos = pos

	if tok == B1 {
		p.pBal++
	} else if tok == B2 {
		p.pBal--
	}

	if TokenLog {
		pos := p.Pos()
		log.Printf("scan reached %v %v {%d}.\n", pos.String(), pos.ShortenedLineheadMsg(), p.pBal)
	}

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) Unscan() {
	p.buf.n = 1

	if p.buf.tok == B1 {
		p.pBal--
	} else if p.buf.tok == B2 {
		p.pBal++
	}

	p.lineNo = p.buf.pos.lineNo
	p.currentLineHead = p.buf.pos.currentLineHead

	if TokenLog {
		log.Printf("Unparsed -> %v\n", p.buf.pos.ShortenedLineheadMsg())
	}
}

// PBal returns the number of ( parsed, minus the number of ). Initially 0.
func (p *Parser) PBal() int {
	return p.pBal
}

// LineNo returns the current line number, starting with 1
func (p *Parser) LineNo() int {
	return p.lineNo
}

// ScanIgnoreWSAndComment scans the next token, and repeats until it finds neither whitespace (including EOL) nor comment.
// litPos is where the returned literal starts.
func (p *Parser) ScanIgnoreWSAndComment() (tok Token, lit string, litPos ParserPosition) {
	tok, lit, litPos = p.Scan()
	for tok == WS || tok == EOL || tok == LINECOMMENT {
		tok, lit, litPos = p.Scan()
	}
	return
}

// ConsumeTokens parses the next tokens and compares to the expected ones
// ignoring SW, EOL and Comment
//todo naming convention for ignoring all these
func (p *Parser) ConsumeTokens(extoks ...Token) (err error) {
	for _, extok := range extoks {
		tok, lit, pos := p.ScanIgnoreWSAndComment()
		if tok != extok {
			return pos.Errorf("expected token \"%v\", found %v", Tokenname(extok), lit)
		}
	}
	return
}

func FmtPrefixedName(prefix, name string) string {
	return fmt.Sprintf("%v:%v", prefix, name)
}

// Pos is the parsing position in the file where scanning will continue.
func (p *Parser) Pos() ParserPosition {
	return ParserPosition{lineNo: p.lineNo, currentLineHead: p.currentLineHead, sourceName: &p.sourceName}
}
