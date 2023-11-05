package datastore

import (
	"log"
	"regexp"
	"sync"
	"time"

	"github.com/huandu/skiplist"
	"github.com/saurabhy27/redis-database/errs"
	"github.com/saurabhy27/redis-database/model"
	"github.com/saurabhy27/redis-database/utils"
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

func (ds *DataStore) ZAdd(key string, sorted_set []model.SortedSetByte) (int, error) {
	log.Printf("Adding the key %s score %v in sorted set\n", key, sorted_set)
	ds.lock.Lock()
	defer ds.lock.Unlock()
	value, ok := ds.data[key]
	resp := 0
	if ok {
		sList, ok := value.(*skiplist.SkipList)
		if !ok {
			return 0, errs.WrongType
		}
		for _, set := range sorted_set {
			sList.Set(set.Score, set.Member)
			resp += 1
		}
		ds.data[key] = sList
	} else {
		skipList := skiplist.New(skiplist.Float64)
		for _, set := range sorted_set {
			skipList.Set(set.Score, set.Member)
			resp += 1
		}
		ds.data[key] = skipList
	}
	return resp, nil
}

func (ds *DataStore) ZRange(key string, start int, stop int) ([]model.SortedSet, error) {
	log.Printf("Retrieving the key %s from start index %d to stop index %d from sorted set\n", key, start, stop)
	ds.lock.RLock()
	defer ds.lock.RUnlock()
	data := []model.SortedSet{}

	value, ok := ds.data[key]
	if ok {
		sList, ok := value.(*skiplist.SkipList)
		if !ok {
			return nil, errs.WrongType
		}
		start, stop = utils.FormatArrayStartNEndIdx(start, stop, sList.Len())
		if stop >= 0 && start <= stop {
			s := sList.Front()
			currentIndex := 0
			for s != nil && currentIndex <= stop {
				if currentIndex >= start {
					data = append(data, model.SortedSet{Score: s.Score(), Member: string(s.Value.([]byte))})
				}
				currentIndex += 1
				s = s.Next()
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
