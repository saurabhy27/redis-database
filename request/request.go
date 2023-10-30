package request

import (
	"strings"

	"github.com/saurabhy27/redis-database/constants"
	"github.com/saurabhy27/redis-database/errs"
)

type Command struct {
	Cmd          string
	MinReqParams int
}

type Request struct {
	Command Command
	Params  []string
}

var (
	CMDGet    Command = Command{Cmd: constants.GET, MinReqParams: 1}
	CMDDel    Command = Command{Cmd: constants.DEL, MinReqParams: 1}
	CMDExpire Command = Command{Cmd: constants.EXPIRE, MinReqParams: 2}
	CMDKeys   Command = Command{Cmd: constants.KEYS, MinReqParams: 1}
	CMDSet    Command = Command{Cmd: constants.SET, MinReqParams: 2}
	CMDTtl    Command = Command{Cmd: constants.TTL, MinReqParams: 1}
	CMDZAdd   Command = Command{Cmd: constants.ZADD, MinReqParams: 3}
	CMDZRange Command = Command{Cmd: constants.ZRANGE, MinReqParams: 3}
)

func parseCommand(cmd string) (Command, error) {
	switch cmd {
	case constants.GET:
		return CMDGet, nil
	case constants.DEL:
		return CMDDel, nil
	case constants.EXPIRE:
		return CMDExpire, nil
	case constants.KEYS:
		return CMDKeys, nil
	case constants.SET:
		return CMDSet, nil
	case constants.TTL:
		return CMDTtl, nil
	case constants.ZADD:
		return CMDZAdd, nil
	case constants.ZRANGE:
		return CMDZRange, nil
	default:
		return Command{}, errs.InvalidCommand
	}
}

func ParseProtocol(input string) (Request, error) {
	input_splited := strings.Split(input, " ")
	if len(input_splited) == 0 {
		return Request{}, errs.EmptyRequest
	}
	command, err := parseCommand(input_splited[0])
	if err != nil {
		return Request{}, errs.InvalidCommand
	}
	params := input_splited[1:]
	if len(params) < command.MinReqParams {
		return Request{}, errs.MinReqParams
	}
	return Request{Command: command, Params: params}, nil
}
