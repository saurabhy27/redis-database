package processor

import (
	"strconv"
	"strings"

	"github.com/saurabhy27/redis-database/datastore"
	"github.com/saurabhy27/redis-database/errs"
	req "github.com/saurabhy27/redis-database/request"
)

type RequestProcessor struct {
	DataStore datastore.DataStoreInterface
}

func (rp *RequestProcessor) Process(request req.Request) (req.Responce, error) {
	switch request.Command {
	case req.CMDGet:
		return rp.processGet(request)
	case req.CMDSet:
		return rp.processSet(request)
	case req.CMDDel:
		return rp.processDel(request)
	case req.CMDKeys:
		return rp.processKeys(request)
	case req.CMDExpire:
		return rp.processExpire(request)
	case req.CMDTtl:
		return rp.processTtl(request)
	case req.CMDZAdd:
		return rp.processZAdd(request)
	case req.CMDZRange:
		return rp.processZRange(request)

	default:
		return req.Responce{}, errs.InvalidCommand
	}
}

func (rp *RequestProcessor) processGet(request req.Request) (req.Responce, error) {
	data, err := rp.DataStore.Get(request.Params[0])
	if err != nil {
		return req.Responce{}, err
	}
	return req.Responce{Success: true, Value: data}, nil
}

func (rp *RequestProcessor) processSet(request req.Request) (req.Responce, error) {
	key := request.Params[0]
	value := request.Params[1]
	rp.DataStore.Set(key, []byte(value))
	return req.Responce{Success: true, Value: "OK"}, nil
}

func (rp *RequestProcessor) processDel(request req.Request) (req.Responce, error) {
	deleted := rp.DataStore.Delete(request.Params[0])
	return req.Responce{Success: true, Value: deleted}, nil
}

func (rp *RequestProcessor) processKeys(request req.Request) (req.Responce, error) {
	filter := request.Params[0]
	filter = strings.ReplaceAll(filter, "*", "\\\\*")
	filter = strings.ReplaceAll(filter, "?", "\\\\?")
	data, err := rp.DataStore.Keys(filter)
	if err != nil {
		return req.Responce{}, err
	}
	return req.Responce{Success: true, Value: data}, nil
}

func (rp *RequestProcessor) processExpire(request req.Request) (req.Responce, error) {
	key := request.Params[0]
	seconds := request.Params[1]
	secondsTTL, err := strconv.Atoi(seconds)
	if err != nil {
		return req.Responce{}, err
	}
	expires := rp.DataStore.Expire(key, secondsTTL)
	return req.Responce{Success: true, Value: expires}, nil
}

func (rp *RequestProcessor) processTtl(request req.Request) (req.Responce, error) {
	expireTime := rp.DataStore.Ttl(request.Params[0])
	return req.Responce{Success: true, Value: expireTime}, nil
}

func (rp *RequestProcessor) processZAdd(request req.Request) (req.Responce, error) {
	key := request.Params[0]
	score := request.Params[1]
	value := request.Params[2]
	scoreFloat, err := strconv.ParseFloat(score, 32)
	if err != nil {
		return req.Responce{}, errs.InvalidFloatValue
	}
	added, err := rp.DataStore.ZAdd(key, scoreFloat, []byte(value))
	if err != nil {
		return req.Responce{}, err
	}
	return req.Responce{Success: true, Value: added}, nil
}

func (rp *RequestProcessor) processZRange(request req.Request) (req.Responce, error) {
	key := request.Params[0]
	start := request.Params[1]
	stop := request.Params[2]
	startInt, err := strconv.Atoi(start)
	if err != nil {
		return req.Responce{}, errs.InvalidIntValue
	}

	stopInt, err := strconv.Atoi(stop)
	if err != nil {
		return req.Responce{}, errs.InvalidIntValue
	}
	data, err := rp.DataStore.ZRange(key, startInt, stopInt)
	if err != nil {
		return req.Responce{}, err
	}
	return req.Responce{Success: true, Value: data}, nil
}
