package unittest

import (
	"testing"

	"github.com/saurabhy27/redis-database/model"
	"github.com/saurabhy27/redis-database/processor"
	req "github.com/saurabhy27/redis-database/request"
	"github.com/saurabhy27/redis-database/tests/mock"
)

func TestProcessGet(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDGet, Params: []string{"test"}}
	response, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.GetMocked {
		t.Errorf("Mocked Get Function not called")
	}
	val, _ := response.Value.([]byte)
	if string(val) != "test123" {
		t.Errorf("Expected val to be test123, got %v", string(val))
	}
}

func TestProcessSet(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDSet, Params: []string{"test", "test123"}}
	_, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.SetMocked {
		t.Errorf("Mocked Set Function not called")
	}
}

func TestProcessDel(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDDel, Params: []string{"test"}}
	response, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.DeleteMocked {
		t.Errorf("Mocked Del Function not called")
	}
	deleted, _ := response.Value.(int)
	if deleted != 1 {
		t.Errorf("Expected deleted to be 2, got %d", deleted)
	}
}

func TestProcessKeys(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDKeys, Params: []string{"*"}}
	response, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.KeysMocked {
		t.Errorf("Mocked Keys Function not called")
	}
	keys, _ := response.Value.([]string)
	if len(keys) != 2 {
		t.Errorf("Expected keys to be 2, got %d", len(keys))
	}
}

func TestProcessExpire(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDExpire, Params: []string{"test", "1"}}
	response, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.ExpireMocked {
		t.Errorf("Mocked Expire Function not called")
	}
	expired, _ := response.Value.(int)
	if expired != 1 {
		t.Errorf("Expected expire to be 1, got %d", expired)
	}
}

func TestProcessTtl(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDTtl, Params: []string{"test"}}
	response, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.TtlMocked {
		t.Errorf("Mocked TTL Function not called")
	}
	ttl, _ := response.Value.(int)
	if ttl != 1 {
		t.Errorf("Expected ttl to be 1, got %d", ttl)
	}
}

func TestProcessZAdd(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDZAdd, Params: []string{"test", "1", "test123"}}
	response, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.ZAddMocked {
		t.Errorf("Mocked ZADD Function not called")
	}
	zadd, _ := response.Value.(int)
	if zadd != 1 {
		t.Errorf("Expected zadd to be 1, got %d", zadd)
	}
}

func TestProcessZRange(t *testing.T) {
	dataStore := &mock.MockDataStore{}
	reqProcessor := processor.RequestProcessor{DataStore: dataStore}
	request := model.Request{Command: req.CMDZRange, Params: []string{"test", "1", "2"}}
	response, err := reqProcessor.Process(request)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if !dataStore.ZRangeMocked {
		t.Errorf("Mocked ZRANGE Function not called")
	}
	zRange, _ := response.Value.(map[float64]string)
	if zRange[1] != "test123" {
		t.Errorf("Expected zrange[1] to be test123, got %s", zRange[1])
	}
}
