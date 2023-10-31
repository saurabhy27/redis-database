package request

import (
	"strings"

	"github.com/saurabhy27/redis-database/constants"
	"github.com/saurabhy27/redis-database/errs"
	"github.com/saurabhy27/redis-database/model"
)

var (
	CMDGet    = model.Command{Cmd: constants.GET, MinReqParams: 1}
	CMDDel    = model.Command{Cmd: constants.DEL, MinReqParams: 1}
	CMDExpire = model.Command{Cmd: constants.EXPIRE, MinReqParams: 2}
	CMDKeys   = model.Command{Cmd: constants.KEYS, MinReqParams: 1}
	CMDSet    = model.Command{Cmd: constants.SET, MinReqParams: 2}
	CMDTtl    = model.Command{Cmd: constants.TTL, MinReqParams: 1}
	CMDZAdd   = model.Command{Cmd: constants.ZADD, MinReqParams: 3}
	CMDZRange = model.Command{Cmd: constants.ZRANGE, MinReqParams: 3}
)

func parseCommand(cmd string) (model.Command, error) {
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
		return model.Command{}, errs.InvalidCommand
	}
}

func ParseProtocol(input string) (model.Request, error) {
	input_splited := strings.Split(input, " ")
	if len(input_splited) == 0 {
		return model.Request{}, errs.EmptyRequest
	}
	command, err := parseCommand(input_splited[0])
	if err != nil {
		return model.Request{}, errs.InvalidCommand
	}
	params := input_splited[1:]
	if len(params) < command.MinReqParams {
		return model.Request{}, errs.MinReqParams
	}
	return model.Request{Command: command, Params: params}, nil
}
