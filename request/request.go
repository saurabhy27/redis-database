package request

import (
	"strings"

	"github.com/saurabhy27/redis-database/constants"
	"github.com/saurabhy27/redis-database/errs"
)

type Command struct {
	Cmd string
}

type Request struct {
	Command Command
	Params  []string
}

var (
	CMDGet    Command = Command{Cmd: constants.GET}
	CMDDel    Command = Command{Cmd: constants.DEL}
	CMDExpire Command = Command{Cmd: constants.EXPIRE}
	CMDKeys   Command = Command{Cmd: constants.KEYS}
	CMDSet    Command = Command{Cmd: constants.SET}
	CMDTtl    Command = Command{Cmd: constants.TTL}
	CMDZAdd   Command = Command{Cmd: constants.ZADD}
	CMDZRange Command = Command{Cmd: constants.ZRANGE}
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
	return Request{Command: command, Params: input_splited[1:]}, nil
}
