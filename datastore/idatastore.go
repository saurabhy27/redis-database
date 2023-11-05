package datastore

import "github.com/saurabhy27/redis-database/model"

type DataStoreInterface interface {
	Get(key string) ([]byte, error)
	Delete(key string) int
	Expire(key string, seconds int) int
	Keys(filter string) ([]string, error)
	Set(key string, value []byte)
	Ttl(key string) int
	ZAdd(key string, sorted_set []model.SortedSetByte) (int, error)
	ZRange(key string, start int, stop int) ([]model.SortedSet, error)
}
