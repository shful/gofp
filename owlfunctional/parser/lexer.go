package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"unicode"
)

//https://blog.gopheracademy.com/advent-2014/parsers-lexers/

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	EOL
	// CR
	// LF
	WS
	DOUBLECIRCUM // ^^ optional type separator for literal values
	BOOLLIT      // boolean literal, true oder false
	STRINGLIT    // string literal
	INTLIT       // integer literal, can be signed
	FLOATLIT     // floating point number literal, can be signed
	LINECOMMENT  // comment until and uncluding line end
	IRI          // e.g.<http://www.w3.org/2000/01/rdf-schema#>

	// Literals
	IDENT //

	// Misc characters
	AT     // @
	COLON  // :
	EQUALS // =
	SHARP  // #
	COMMA  // ,
	QUOTE  // "
	B1     // (
	B2     // )
	PLUS   // +
	MINUS  // -

	// Keywords
	AnnotationAssertion
	AnnotationProperty
	AsymmetricObjectProperty
	Class
	ClassAssertion
	DataAllValuesFrom
	DataComplementOf
	DataExactCardinality
	DataHasValue
	DataIntersectionOf
	DataMaxCardinality
	DataMinCardinality
	DataOneOf
	DataProperty
	DataPropertyAssertion
	DataPropertyDomain
	DataPropertyRange
	DataSomeValuesFrom
	Datatype
	DatatypeRestriction
	DataUnionOf
	Declaration
	DifferentIndividuals
	DisjointClasses
	EquivalentClasses
	FunctionalDataProperty
	FunctionalObjectProperty
	InverseFunctionalObjectProperty
	InverseObjectProperties
	IrreflexiveObjectProperty
	NamedIndividual
	ObjectComplementOf
	ObjectAllValuesFrom
	ObjectExactCardinality
	ObjectHasSelf
	ObjectHasValue
	ObjectIntersectionOf
	ObjectInverseOf
	ObjectMaxCardinality
	ObjectMinCardinality
	ObjectOneOf
	ObjectProperty
	ObjectPropertyDomain
	ObjectPropertyRange
	ObjectSomeValuesFrom
	ObjectUnionOf
	Ontology
	OWLBottomDataProperty
	OWLNothing
	OWLFalse
	OWLThing
	OWLTrue
	OWLTopDataProperty
	Prefix
	ReflexiveObjectProperty
	SubClassOf
	SubDataPropertyOf
	SubObjectPropertyOf
	SymmetricObjectProperty
	TransitiveObjectProperty
)

var keywords map[string]Token = map[string]Token{
	"AnnotationAssertion":             AnnotationAssertion,
	"AnnotationProperty":              AnnotationProperty,
	"AsymmetricObjectProperty":        AsymmetricObjectProperty,
	"Class":                           Class,
	"ClassAssertion":                  ClassAssertion,
	"DataAllValuesFrom":               DataAllValuesFrom,
	"DataComplementOf":                DataComplementOf,
	"DataExactCardinality":            DataExactCardinality,
	"DataHasValue":                    DataHasValue,
	"DataIntersectionOf":              DataIntersectionOf,
	"DataMaxCardinality":              DataMaxCardinality,
	"DataMinCardinality":              DataMinCardinality,
	"DataOneOf":                       DataOneOf,
	"DataProperty":                    DataProperty,
	"DataPropertyAssertion":           DataPropertyAssertion,
	"DataPropertyDomain":              DataPropertyDomain,
	"DataPropertyRange":               DataPropertyRange,
	"DataSomeValuesFrom":              DataSomeValuesFrom,
	"Datatype":                        Datatype,
	"DatatypeRestriction":             DatatypeRestriction,
	"DataUnionOf":                     DataUnionOf,
	"Declaration":                     Declaration,
	"DifferentIndividuals":            DifferentIndividuals,
	"DisjointClasses":                 DisjointClasses,
	"EquivalentClasses":               EquivalentClasses,
	"false":                           OWLFalse,
	"FunctionalDataProperty":          FunctionalDataProperty,
	"FunctionalObjectProperty":        FunctionalObjectProperty,
	"InverseFunctionalObjectProperty": InverseFunctionalObjectProperty,
	"IrreflexiveObjectProperty":       IrreflexiveObjectProperty,
	"InverseObjectProperties":         InverseObjectProperties,
	"NamedIndividual":                 NamedIndividual,
	"ObjectAllValuesFrom":             ObjectAllValuesFrom,
	"ObjectComplementOf":              ObjectComplementOf,
	"ObjectExactCardinality":          ObjectExactCardinality,
	"ObjectHasSelf":                   ObjectHasSelf,
	"ObjectHasValue":                  ObjectHasValue,
	"ObjectIntersectionOf":            ObjectIntersectionOf,
	"ObjectInverseOf":                 ObjectInverseOf, //not test covered
	"ObjectMaxCardinality":            ObjectMaxCardinality,
	"ObjectMinCardinality":            ObjectMinCardinality,
	"ObjectOneOf":                     ObjectOneOf,
	"ObjectProperty":                  ObjectProperty,
	"ObjectPropertyDomain":            ObjectPropertyDomain,
	"ObjectPropertyRange":             ObjectPropertyRange,
	"ObjectSomeValuesFrom":            ObjectSomeValuesFrom,
	"ObjectUnionOf":                   ObjectUnionOf,
	"Ontology":                        Ontology,
	"Prefix":                          Prefix,
	"ReflexiveObjectProperty":         ReflexiveObjectProperty,
	"SubClassOf":                      SubClassOf,
	"SubDataPropertyOf":               SubDataPropertyOf,
	"SubObjectPropertyOf":             SubObjectPropertyOf,
	"SymmetricObjectProperty":         SymmetricObjectProperty,
	"TransitiveObjectProperty":        TransitiveObjectProperty,
	"true": OWLTrue,
}

func Tokenname(t Token) string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case AT:
		return "AT"
	case DOUBLECIRCUM:
		return "DOUBLECIRCUM"
	case EOF:
		return "EOF"
	case EOL:
		return "EOL"
	case WS:
		return "WS"
	case COLON:
		return "COLON"
	case EQUALS:
		return "EQUALS"
	case COMMA:
		return "COMMA"
	case FLOATLIT:
		return "FLOATLIT"
	case IDENT:
		return "IDENT"
	case INTLIT:
		return "INTLIT"
	case IRI:
		return "IRI"
	case LINECOMMENT:
		return "LINECOMMENT"
	case B1:
		return "B1"
	case B2:
		return "B2"
	case STRINGLIT:
		return "STRINGLIT"
	case SHARP:
		return "SHARP"
	case QUOTE:
		return "QUOTE"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	}
	for k, v := range keywords {
		if v == t {
			return k
		}
	}
	return fmt.Sprintf("%d", t)
}

func isSign(ch rune) bool {
	return ch == '+' || ch == '-'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' //|| ch == '\n'
}

func isLineOrFileEnd(ch rune) bool {
	return ch == cr || ch == lf || ch == eof
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

// func isLetter(ch rune) bool {
// 	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
// }

var eof = rune(0)
var cr = '\r'
var lf = '\n'

// Scanner //////////////
// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// scan returns the next token and literal value.
func (s *Scanner) scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if ch == lf || ch == cr {
		s.unread()
		return s.scanEOL()
	} else if unicode.IsLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if ch == '"' {
		s.unread()
		return s.scanStringliteral()
	} else if ch == '#' {
		s.unread()
		return s.scanLinecomment()
	} else if ch == '^' {
		s.unread()
		return s.scanDoubleCircum()
	} else if ch == '<' {
		s.unread()
		return s.scanIRI()
	} else if isSign(ch) || isDigit(ch) {
		s.unread()
		return s.scanNumber()
	}
	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case ':':
		return COLON, string(ch)
	case '=':
		return EQUALS, string(ch)
	case ',':
		return COMMA, string(ch)
	case '"':
		return QUOTE, string(ch)
	case '(':
		return B1, string(ch)
	case ')':
		return B2, string(ch)
	case '+':
		return PLUS, string(ch)
	case '-':
		return MINUS, string(ch)
	case '@':
		return AT, string(ch)
	}

	return ILLEGAL, string(ch)
}

// End Scanner //////////////////

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanStringliteral consumes the current rune and all contiguous literal tokens.
// Both surrounding quotes are excluded.
//todo: are Line Breaks in literals allowed ?
func (s *Scanner) scanStringliteral() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	s.read() // first quote
	// buf.WriteRune(s.read())

	// Read every subsequent string literal character into the buffer.
	// Quote end EOF will cause the loop to exit.
	var chPrev rune
	for {

		ch := s.read()
		if ch == eof {
			s.unread()
			break
		}

		// break at "
		if ch == '"' && chPrev != '\\' {
			break
		}
		buf.WriteRune(ch)
		chPrev = ch
	}

	return STRINGLIT, buf.String()
}

// scanIRI consumes the current rune and all contiguous literal tokens until >
// Both surrounding <> are included.
//todo Same as string literal, are line breaks possible?
func (s *Scanner) scanIRI() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent string literal character into the buffer.
	// Quote end EOF/EOL will cause the loop to exit.
	// var chPrev rune
	for {

		ch := s.read()

		if isLineOrFileEnd(ch) {
			s.unread()
			break
		}
		buf.WriteRune(ch)

		// break at >
		if ch == '>' {
			break
		}
		// chPrev = ch
	}

	return IRI, buf.String()
}

// scanEOL consumes the current CR or LF.
// In case of CR, also consumes an optionally following LF.
func (s *Scanner) scanEOL() (tok Token, lit string) {

	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	ch := s.read()
	if ch == eof {
		s.unread()
		return EOL, buf.String()
	}

	buf.WriteRune(ch)

	// break at LF
	if ch == lf {
		return EOL, buf.String()
	}

	// optionally, consume an LF after CR
	ch = s.read()
	if ch == lf {
		buf.WriteRune(ch)
	} else {
		s.unread()
	}
	return EOL, buf.String()
}

// scanLinecomment consumes the current rune and all contiguous tokens until End-Of-Line.
func (s *Scanner) scanLinecomment() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent string literal character into the buffer.
	// EOL and EOF will cause the loop to exit. Both are not consumed.
	for {

		ch := s.read()
		if isLineOrFileEnd(ch) {
			s.unread()
			break
		}

		buf.WriteRune(ch)
	}

	return LINECOMMENT, buf.String()
}

// scanEOL consumes the current double circumflex
// if not found, as expected, returns the next character  which is single circumflex of type ILLEGAL.
func (s *Scanner) scanDoubleCircum() (tok Token, lit string) {

	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	// first circumflex
	ch := s.read()
	buf.WriteRune(ch)

	// second, probably also a circumflex
	ch = s.read()
	if ch != '^' {
		s.unread()
		return ILLEGAL, buf.String()
	}

	buf.WriteRune(ch)

	return DOUBLECIRCUM, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters including EOF/EOL will cause the loop to exit.
	for {
		if ch := s.read(); !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	// switch buf.String() {
	// case "Declaration":
	// 	return Declaration, buf.String()
	if tok, ok := keywords[buf.String()]; ok {
		return tok, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

// scanNumber consumes the current rune and all contiguous number runes.
// Numbers can be floating point, i.e. contain a single dot. Numbers can have one leading + or - sign.
func (s *Scanner) scanNumber() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read()) // sign or first digit

	// Read every subsequent ident character into the buffer.
	// Non-number chars cause exit.
	for {
		// no need to stop with a second dot. There is no regular way for . to directly follow a number, i.e.
		// a second dot surely is an error.
		if ch := s.read(); !unicode.IsDigit(ch) && ch != '.' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a number then return the right number type
	// There could still be an unparseable number with multiple dots in the buffer.
	bs := buf.String()
	var err error
	if _, err = strconv.Atoi(bs); err == nil {
		return INTLIT, bs
	}
	if _, err = strconv.ParseFloat(bs, 64); err == nil {
		return FLOATLIT, bs
	}

	// Otherwise give up:
	return ILLEGAL, bs
}
