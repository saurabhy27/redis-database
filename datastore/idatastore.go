package datastore

type DataStoreInterface interface {
	Get(key string) ([]byte, error)
	Delete(key string) int
	Expire(key string, seconds int) int
	Keys(filter string) ([]string, error)
	Set(key string, value []byte)
	Ttl(key string) int
	ZAdd(key string, score float64, member []byte) (int, error)
	ZRange(key string, start float64, stop float64) (map[float64]string, error)
}
