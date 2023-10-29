package datastore

type DataStoreInterface interface {
	Get(key string) ([]byte, error)
	Delete(key string) error
	Expire(key string, seconds int) error
	Keys(filter string) ([]string, error)
	Set(key string, value []byte)
	Ttl(key string) int
	ZAdd(key string, score float64, member []byte) error
	ZRange(key string, start float64, stop float64) (map[float64][]byte, error)
}
