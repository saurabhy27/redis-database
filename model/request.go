package model

type Command struct {
	Cmd          string
	MinReqParams int
}

type Request struct {
	Command Command
	Params  []string
}
