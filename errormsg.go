package gofp

import (
	"fmt"

	"github.com/shful/gofp/owlfunctional/parser"
)

// ErrorMsgWithPosition produces a user message like "untempting topping on margherita pizza in: pizza.owl 18:32 after '...Pizza ObjectSomeValuesFrom(:hasTopping'"
// for errors or type parser.PErr. For other error types, err.Error() is returned.
func ErrorMsgWithPosition(err error) string {
	if perr, ok := err.(*parser.PErr); ok {
		return fmt.Sprintf("%v in:%v", perr.Msg, ParserPositionMsg(perr.AfterPos))
	}
	return err.Error()
}

// ParserPositionMsg returns a position desciption for the user ,
// like "pizza.owl 18:32 after '...Pizza ObjectSomeValuesFrom(:hasTopping'"
func ParserPositionMsg(pos parser.ParserPosition) string {
	return fmt.Sprintf("%v %d:%d %v", pos.SourceName(), pos.LineNo1(), pos.ColNo1(), pos.ShortenedLineheadMsg())
}
