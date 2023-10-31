package datastore

import (
	"log"
	"regexp"
	"sync"
	"time"

	"github.com/saurabhy27/redis-database/errs"
	"github.com/saurabhy27/redis-database/model"
)

type DataStore struct {
	lock       sync.RWMutex   // to avoid modifing values from multiple goroutines
	data       map[string]any // key:value
	expireData map[string]int // Key:expireEpoxTimestamp
}

func New() *DataStore {
	return &DataStore{data: make(map[string]any), expireData: make(map[string]int)}
}

func (ds *DataStore) Set(key string, value []byte) {
	// setting the values in the map
	log.Printf("Seting the value for key %s\n", key)
	ds.lock.Lock()
	defer ds.lock.Unlock()
	ds.data[key] = value
}

func (ds *DataStore) Get(key string) ([]byte, error) {
	// getting the values in the map
	log.Printf("Fetching the value for key %s\n", key)
	ds.lock.RLock()
	defer ds.lock.RUnlock()
	value, ok := ds.data[key]
	if !ok {
		return nil, nil
	}
	v, ok := value.([]byte)
	if !ok {
		return nil, errs.WrongType
	}
	return v, nil
}

func (ds *DataStore) Delete(key string) int {
	log.Printf("Deleting the key %s\n", key)
	ds.lock.Lock()
	defer ds.lock.Unlock()
	_, ok := ds.data[key]
	if !ok {
		return 0
	}
	delete(ds.data, key)
	delete(ds.expireData, key)
	return 1
}

func (ds *DataStore) Keys(filter string) ([]string, error) {
	log.Printf("Fetching all the keys with filter %s\n", filter)
	ds.lock.RLock()
	defer ds.lock.RUnlock()
	var keys []string
	for k := range ds.data {
		keys = append(keys, k)
	}

	var regMatchKeys []string
	re, err := regexp.Compile(filter)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		matches := re.FindAllString(key, -1)
		if len(matches) != 0 {
			regMatchKeys = append(regMatchKeys, key)
		}
	}
	return regMatchKeys, nil
}

func (ds *DataStore) expireInBackground(key string, seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
	ds.Delete(key)
}

func (ds *DataStore) Expire(key string, seconds int) int {
	log.Printf("Expiring the keys %s in %d seconds\n", key, seconds)
	_, ok := ds.data[key]
	if !ok {
		return 0
	}
	ds.expireData[key] = int(time.Now().Unix()) + seconds
	go ds.expireInBackground(key, seconds)
	return 1
}

func (ds *DataStore) ZAdd(key string, score float64, member []byte) (int, error) {
	log.Printf("Adding the key %s score %f in sorted set\n", key, score)
	ds.lock.Lock()
	defer ds.lock.Unlock()
	value, ok := ds.data[key]
	if ok {
		zList, ok := value.([]model.ZaddModel)
		if !ok {
			return 0, errs.WrongType
		}
		zList = append(zList, model.ZaddModel{Score: score, Member: []byte(member)})
		ds.data[key] = zList
	} else {
		zaddList := []model.ZaddModel{}
		zaddList = append(zaddList, model.ZaddModel{Score: score, Member: []byte(member)})
		ds.data[key] = zaddList
	}
	return 1, nil
}

func (ds *DataStore) ZRange(key string, start int, stop int) (map[float64]string, error) {
	log.Printf("Retrieving the key %s from start index %d to stop index %d from sorted set\n", key, start, stop)
	ds.lock.RLock()
	defer ds.lock.RUnlock()
	data := make(map[float64]string)

	value, ok := ds.data[key]
	if ok {
		sList, ok := value.([]model.ZaddModel)
		if !ok {
			return nil, errs.WrongType
		}
		if len(sList) > start {
			for i := start; i < len(sList) && i <= stop; i++ {
				data[sList[i].Score] = string(sList[i].Member)
			}

		}

	}
	return data, nil
}

func (ds *DataStore) Ttl(key string) int {
	log.Printf("Retrieving the expire of key %s\n", key)
	ds.lock.RLock()
	defer ds.lock.RUnlock()
	expireEpoxTimestamp, ok := ds.expireData[key]
	if ok {
		diff := expireEpoxTimestamp - int(time.Now().Unix())
		if diff > 0 {
			return diff
		}
	}
	return -1
}
