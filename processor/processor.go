package processor

import (
	"strconv"

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
	return req.Responce{Success: true}, nil
}

func (rp *RequestProcessor) processDel(request req.Request) (req.Responce, error) {
	rp.DataStore.Delete(request.Params[0])
	return req.Responce{Success: true}, nil
}

func (rp *RequestProcessor) processKeys(request req.Request) (req.Responce, error) {
	filter := ""
	if len(request.Params) != 0 {
		filter = request.Params[0]
	}
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
	err = rp.DataStore.Expire(key, secondsTTL)
	if err != nil {
		return req.Responce{}, err
	}
	return req.Responce{Success: true}, nil
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
		return req.Responce{}, err
	}

	err = rp.DataStore.ZAdd(key, scoreFloat, []byte(value))
	if err != nil {
		return req.Responce{}, err
	}
	return req.Responce{Success: true}, nil
}

func (rp *RequestProcessor) processZRange(request req.Request) (req.Responce, error) {
	key := request.Params[0]
	start := request.Params[1]
	stop := request.Params[2]
	startFloat, err := strconv.ParseFloat(start, 32)
	if err != nil {
		return req.Responce{}, err
	}

	stopFloat, err := strconv.ParseFloat(stop, 32)
	if err != nil {
		return req.Responce{}, err
	}
	data, err := rp.DataStore.ZRange(key, startFloat, stopFloat)
	if err != nil {
		return req.Responce{}, err
	}
	return req.Responce{Success: true, Value: data}, nil
}
