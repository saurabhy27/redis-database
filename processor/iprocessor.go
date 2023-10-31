package processor

import "github.com/saurabhy27/redis-database/model"

type RequestProcessorInterface interface {
	Process(request model.Request) (model.Responce, error)
}
