package reports

import "go/token"

// Report captures the position of the current token and the next token, used as
// the diagnostic's start/end range.
type Report struct {
	Pos          token.Pos
	NextTokenPos token.Pos
	Category     string
	Message      string
}
