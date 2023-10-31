package unittest

import (
	"testing"

	"github.com/saurabhy27/redis-database/constants"
	"github.com/saurabhy27/redis-database/errs"
	"github.com/saurabhy27/redis-database/request"
	"github.com/saurabhy27/redis-database/utils"
)

func TestGetValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("GET test")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "test") {
		t.Errorf("Expected %s to be %v", "test", command.Params)
	}
	if command.Command.Cmd != constants.GET {
		t.Errorf("Expected CMD to be %s, got %s", constants.GET, command.Command.Cmd)
	}
}

func TestGetInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("GET")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}

func TestDelValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("DEL test")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "test") {
		t.Errorf("Expected %s to be %v", "test", command.Params)
	}
	if command.Command.Cmd != constants.DEL {
		t.Errorf("Expected CMD to be %s, got %s", constants.DEL, command.Command.Cmd)
	}
}

func TestDelInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("DEL")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}

func TestExpireValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("EXPIRE test 2")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "test") {
		t.Errorf("Expected %s to be %v", "test", command.Params)
	}
	if command.Command.Cmd != constants.EXPIRE {
		t.Errorf("Expected CMD to be %s, got %s", constants.EXPIRE, command.Command.Cmd)
	}
}

func TestExpireInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("EXPIRE test")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}

func TestKeysValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("KEYS *")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "*") {
		t.Errorf("Expected %s to be %v", "*", command.Params)
	}
	if command.Command.Cmd != constants.KEYS {
		t.Errorf("Expected CMD to be %s, got %s", constants.KEYS, command.Command.Cmd)
	}
}

func TestKeysInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("KEYS")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}

func TestSetValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("SET test test123")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "test") {
		t.Errorf("Expected %s to be %v", "test", command.Params)
	}
	if command.Command.Cmd != constants.SET {
		t.Errorf("Expected CMD to be %s, got %s", constants.SET, command.Command.Cmd)
	}
}

func TestSetInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("SET test")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}

func TestTtlValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("TTL test")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "test") {
		t.Errorf("Expected %s to be %v", "test", command.Params)
	}
	if command.Command.Cmd != constants.TTL {
		t.Errorf("Expected CMD to be %s, got %s", constants.TTL, command.Command.Cmd)
	}
}

func TestTtlInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("TTL")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}

func TestZAddValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("ZADD test 10 test123")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "test") {
		t.Errorf("Expected %s to be %v", "test", command.Params)
	}
	if command.Command.Cmd != constants.ZADD {
		t.Errorf("Expected CMD to be %s, got %s", constants.ZADD, command.Command.Cmd)
	}
}

func TestZAddInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("ZADD test 10")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}

func TestZRangeValidParseProtocol(t *testing.T) {
	command, err := request.ParseProtocol("ZRANGE test 0 1")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !utils.Contains(command.Params, "test") {
		t.Errorf("Expected %s to be %v", "test", command.Params)
	}
	if command.Command.Cmd != constants.ZRANGE {
		t.Errorf("Expected CMD to be %s, got %s", constants.ZRANGE, command.Command.Cmd)
	}
}

func TestZRangeInvalidParseProtocol(t *testing.T) {
	_, err := request.ParseProtocol("ZRANGE test")
	if err != errs.MinReqParams {
		t.Errorf("Expected err to be %v, got %v", errs.MinReqParams, err)
	}
}
