package processor

import req "github.com/saurabhy27/redis-database/request"

type RequestProcessorInterface interface {
	Process(request req.Request) (req.Responce, error)
}
