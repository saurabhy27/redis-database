package processor

import (
	"strconv"
	"strings"

	"github.com/saurabhy27/redis-database/datastore"
	"github.com/saurabhy27/redis-database/errs"
	"github.com/saurabhy27/redis-database/model"
	req "github.com/saurabhy27/redis-database/request"
)

type RequestProcessor struct {
	DataStore datastore.DataStoreInterface
}

func (rp *RequestProcessor) Process(request model.Request) (model.Responce, error) {
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
		return model.Responce{}, errs.InvalidCommand
	}
}

func (rp *RequestProcessor) processGet(request model.Request) (model.Responce, error) {
	data, err := rp.DataStore.Get(request.Params[0])
	if err != nil {
		return model.Responce{}, err
	}
	return model.Responce{Success: true, Value: data}, nil
}

func (rp *RequestProcessor) processSet(request model.Request) (model.Responce, error) {
	key := request.Params[0]
	value := request.Params[1]
	rp.DataStore.Set(key, []byte(value))
	return model.Responce{Success: true, Value: "OK"}, nil
}

func (rp *RequestProcessor) processDel(request model.Request) (model.Responce, error) {
	deleted := rp.DataStore.Delete(request.Params[0])
	return model.Responce{Success: true, Value: deleted}, nil
}

func (rp *RequestProcessor) processKeys(request model.Request) (model.Responce, error) {
	filter := request.Params[0]
	filter = strings.ReplaceAll(filter, "*", "\\\\*")
	filter = strings.ReplaceAll(filter, "?", "\\\\?")
	data, err := rp.DataStore.Keys(filter)
	if err != nil {
		return model.Responce{}, err
	}
	return model.Responce{Success: true, Value: data}, nil
}

func (rp *RequestProcessor) processExpire(request model.Request) (model.Responce, error) {
	key := request.Params[0]
	seconds := request.Params[1]
	secondsTTL, err := strconv.Atoi(seconds)
	if err != nil {
		return model.Responce{}, err
	}
	expires := rp.DataStore.Expire(key, secondsTTL)
	return model.Responce{Success: true, Value: expires}, nil
}

func (rp *RequestProcessor) processTtl(request model.Request) (model.Responce, error) {
	expireTime := rp.DataStore.Ttl(request.Params[0])
	return model.Responce{Success: true, Value: expireTime}, nil
}

func (rp *RequestProcessor) processZAdd(request model.Request) (model.Responce, error) {
	key := request.Params[0]

	param := request.Params[1:]
	if len(param)%2 != 0 {
		return model.Responce{}, errs.SyntaxError
	}
	var sorted_set []model.SortedSetByte

	for i := 0; i < len(param); i += 2 {
		score := param[i]
		value := param[i+1]
		scoreFloat, err := strconv.ParseFloat(score, 32)
		if err != nil {
			return model.Responce{}, errs.InvalidFloatValue
		}
		sorted_set = append(sorted_set, model.SortedSetByte{Score: scoreFloat, Member: []byte(value)})
	}

	added, err := rp.DataStore.ZAdd(key, sorted_set)
	if err != nil {
		return model.Responce{}, err
	}
	return model.Responce{Success: true, Value: added}, nil
}

func (rp *RequestProcessor) processZRange(request model.Request) (model.Responce, error) {
	key := request.Params[0]
	start := request.Params[1]
	stop := request.Params[2]
	startInt, err := strconv.Atoi(start)
	if err != nil {
		return model.Responce{}, errs.InvalidIntValue
	}

	stopInt, err := strconv.Atoi(stop)
	if err != nil {
		return model.Responce{}, errs.InvalidIntValue
	}
	data, err := rp.DataStore.ZRange(key, startInt, stopInt)
	if err != nil {
		return model.Responce{}, err
	}
	return model.Responce{Success: true, Value: data}, nil
}
