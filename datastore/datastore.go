package datastore

import (
	"regexp"
	"sync"
	"time"

	"github.com/huandu/skiplist"
	"github.com/saurabhy27/redis-database/errs"
)

type DataStore struct {
	lock       sync.RWMutex
	data       map[string]any
	expireData map[string]int // Key:expireEpoxTimestamp
}

func NewDataStore() *DataStore {
	return &DataStore{data: make(map[string]any), expireData: make(map[string]int)}
}

func (ds *DataStore) Set(key string, value []byte) {
	ds.lock.Lock()
	defer ds.lock.Unlock()
	ds.data[key] = value
}

func (ds *DataStore) Get(key string) ([]byte, error) {
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
	_, ok := ds.data[key]
	if !ok {
		return 0
	}
	ds.expireData[key] = int(time.Now().Unix()) + seconds
	go ds.expireInBackground(key, seconds)
	return 1
}

func (ds *DataStore) ZAdd(key string, score float64, member []byte) (int, error) {
	ds.lock.Lock()
	defer ds.lock.Unlock()
	value, ok := ds.data[key]
	if ok {
		sList, ok := value.(*skiplist.SkipList)
		if !ok {
			return 0, errs.WrongType
		}
		sList.Set(sList.Back().Score()+1, map[float64][]byte{score: []byte(member)})
		ds.data[key] = sList
	} else {
		skipList := skiplist.New(skiplist.Float64)
		skipList.Set(0, map[float64][]byte{score: []byte(member)})
		ds.data[key] = skipList
	}
	return 1, nil
}

func (ds *DataStore) ZRange(key string, start float64, stop float64) (map[float64]string, error) {
	ds.lock.RLock()
	defer ds.lock.RUnlock()
	data := make(map[float64]string)

	value, ok := ds.data[key]
	if ok {
		sList, ok := value.(*skiplist.SkipList)
		if !ok {
			return nil, errs.WrongType
		}
		elem := sList.Find(start)
		for elem != nil && elem.Score() <= stop {
			val, _ := elem.Value.(map[float64][]byte)
			for k, v := range val {
				data[k] = string(v)
			}
			elem = elem.Next()
		}
	}
	return data, nil
}

func (ds *DataStore) Ttl(key string) int {
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
