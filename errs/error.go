package errs

import "errors"

var (
	NoDataFound    = errors.New("NO DATA FOR THE GIVEN KEY")
	InvalidCommand = errors.New("INVALID COMMAND")
	EmptyRequest   = errors.New("EMPTY REQUESTS")
	WrongType      = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
)
