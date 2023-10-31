package errs

import "errors"

var (
	NoDataFound       = errors.New("NO DATA FOR THE GIVEN KEY")
	InvalidCommand    = errors.New("unknown command")
	EmptyRequest      = errors.New("EMPTY REQUESTS")
	WrongType         = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
	MinReqParams      = errors.New("wrong number of arguments for given command")
	InvalidFloatValue = errors.New("value is not a valid float")
	InvalidIntValue   = errors.New("value is not a valid int")
)
