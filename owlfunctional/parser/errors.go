package parser

import (
	"fmt"
)

type PErr struct {
	Msg string

	// AfterPos is the Position where the error is directly behind.
	AfterPos ParserPosition
}

func NewErr(msg string, pos ParserPosition) error {
	return &PErr{Msg: msg, AfterPos: pos}
}

// Error satifies the error interface.
func (s *PErr) Error() string {
	return fmt.Sprintf("%v after %v", s.Msg, s.AfterPos.String())
	// return s.String()
}

// String returns a readable representation of error and position.
// func (s *PErr) String() string {
// 	return fmt.Sprintf("%v after %v", s.Msg, s.AfterPos.String())
// }

// DescribeToklit returns a readable presentation of both the last token and, if printable, the literal.
func DescribeToklit(tok Token, lit string) string {
	// tokens where we can omit the literal value without information loss, or where literal is unprintable:
	if tok == EOF || tok == EOL || tok == WS || tok == COLON || tok == COMMA || tok == EQUALS || tok == QUOTE || tok == B1 || tok == B2 {
		lit = ""
	}
	return fmt.Sprintf("%v %v", Tokenname(tok), lit)
}
