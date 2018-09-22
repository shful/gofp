package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ParserPosition is a Snapshot of the parsing position in a file
type ParserPosition struct {
	lineNo          int     // >= 0
	currentLineHead string  // currentLineHead is similar to Parser.currentLineHead
	sourceName      *string // this is the sourceName attribute of the parser. sourceName tells the user what is parsed, e.g. a filename.
}

// ColNo1 is the parsing position in the line, starting with 1
func (p *ParserPosition) ColNo1() int {
	return utf8.RuneCountInString(p.currentLineHead) + 1
}

// ColNo1 is the parsing position in the line, starting with 1
// A tabsize >=0 can be specified. This allows to match the column positions when the file is shown in an editor.
func (p *ParserPosition) ColNo1WithTabsize(tabsize int) int {
	tabCnt := strings.Count(p.currentLineHead, "\t")
	return p.ColNo1() + (tabsize-1)*tabCnt
}

func (p *ParserPosition) String() string {
	return fmt.Sprintf("ParserPosition{%d %d}", p.LineNo1(), p.ColNo1())
}

// LineNo1 starts with 1
func (p *ParserPosition) LineNo1() int {
	return p.lineNo + 1
}

// GetCurrentLineHead is the line belonging to lineNo, until -and including- the
// literal starting at the current column.
func (p *ParserPosition) GetCurrentLineHead() string {
	return p.currentLineHead
}

// SourceName - see attribute sourceName
func (p *ParserPosition) SourceName() string {
	return *p.sourceName
}

func (pos *ParserPosition) ErrorfUnexpectedToken(tok Token, lit string, need string) error {
	return pos.Errorf("unexpected %v(literal=%v), need %v", Tokenname(tok), lit, need)
}

func (pos *ParserPosition) Errorf(msg string, fmtargs ...interface{}) error {
	return pos.EnsurePErr(fmt.Errorf(msg, fmtargs...))
}

// EnsurePErr returns err, if this is already a PErr.
// Otherwise, creates a new PErr this this Position, and the message string from err.
func (pos *ParserPosition) EnsurePErr(err error) *PErr {
	if perr, ok := err.(*PErr); ok {
		return perr
	}
	return &PErr{Msg: err.Error(), AfterPos: *pos}
}

// EnrichErrorMsg returns err or a new PErr
// with a modified message, so that msgPrefix is added before the existing message.
func (pos *ParserPosition) EnrichErrorMsg(err error, msgPrefix string) error {
	perr := pos.EnsurePErr(err)
	perr.Msg = fmt.Sprintf("%v:%v", msgPrefix, perr.Msg)
	return perr
}

// ShortenedLineheadMsg is for convenience - a shortened linehead either, or a message saying "at start of line"
func (pos *ParserPosition) ShortenedLineheadMsg() (linehead string) {
	linehead = pos.currentLineHead
	l := len(linehead)

	if l > 0 {
		if l > 20 {
			linehead = "..." + linehead[l-20:]
		}
		linehead = fmt.Sprintf("after '%v'", linehead)
	} else {
		linehead = "at start of line"
	}
	return
}
