package unittest

import (
	"testing"
	"time"

	"github.com/saurabhy27/redis-database/datastore"
	"github.com/saurabhy27/redis-database/utils"
)

func TestGet(t *testing.T) {
	dsStore := datastore.New()
	val, err := dsStore.Get("abc")
	if val != nil {
		t.Errorf("Expected val to be nil, got %s", string(val))
	}
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

func TestSet(t *testing.T) {
	dsStore := datastore.New()
	key, expVal := "test", []byte("test123")
	dsStore.Set(key, expVal)
	actVal, err := dsStore.Get(key)
	if actVal == nil {
		t.Errorf("Expected val to be %s, got %s", string(expVal), string(actVal))
	}

	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}

	if string(expVal) != string(actVal) {
		t.Errorf("Expected val to be %s, got %s", string(expVal), string(actVal))
	}
}

func TestDelete(t *testing.T) {
	dsStore := datastore.New()
	key, expVal := "test", []byte("test123")
	dsStore.Set(key, expVal)
	dsStore.Delete(key)
	val, err := dsStore.Get(key)
	if val != nil {
		t.Errorf("Expected val to be nil, got %s", string(val))
	}
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

func TestKeys(t *testing.T) {
	dsStore := datastore.New()
	key1, expVal1 := "test", []byte("test123")
	key2, expVal2 := "care", []byte("care123")
	dsStore.Set(key1, expVal1)
	dsStore.Set(key2, expVal2)
	actKeys, err := dsStore.Keys("\\\\*")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if len(actKeys) != 2 {
		t.Errorf("Expected keys to be 2, got %d", len(actKeys))
	}
	if !utils.Contains(actKeys, key1) {
		t.Errorf("Expected %s to be %v", key1, actKeys)
	}
	if !utils.Contains(actKeys, key2) {
		t.Errorf("Expected %s to be %v", key2, actKeys)
	}
}

func TestExpire(t *testing.T) {
	dsStore := datastore.New()
	key1, expVal1 := "test", []byte("test123")
	key2, expVal2 := "care", []byte("care123")
	dsStore.Set(key1, expVal1)
	dsStore.Set(key2, expVal2)
	actKeys, _ := dsStore.Keys("\\\\*")
	if len(actKeys) != 2 {
		t.Errorf("Expected keys to be 2, got %d", len(actKeys))
	}
	actExp := dsStore.Expire(key1, 1)
	if actExp != 1 {
		t.Errorf("Expected expire value to be %v, got %d", 1, actExp)
	}
	time.Sleep(2 * time.Second)
	actKeys, _ = dsStore.Keys("\\\\*")
	if utils.Contains(actKeys, key1) {
		t.Errorf("Expected %s not to be %v", key1, actKeys)
	}
	if len(actKeys) != 1 {
		t.Errorf("Expected keys to be 1, got %d", len(actKeys))
	}
	if !utils.Contains(actKeys, key2) {
		t.Errorf("Expected %s to be %v", key2, actKeys)
	}
}

func TestZAdd(t *testing.T) {
	dsStore := datastore.New()
	key, score, expVal := "test", 10.0, []byte("test123")
	actVal, err := dsStore.ZAdd(key, score, expVal)
	if actVal != 1 {
		t.Errorf("Expected val to be 1, got %d", actVal)
	}
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
}

func TestZRange(t *testing.T) {
	dsStore := datastore.New()
	key, score1, expVal1 := "test", 10.0, []byte("test123")
	score2, expVal2 := 21.1, []byte("care123")
	dsStore.ZAdd(key, score1, expVal1)
	dsStore.ZAdd(key, score2, expVal2)
	actVal, err := dsStore.ZRange(key, 0, 1)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if actVal[0].Member != string(expVal1) {
		t.Errorf("Expected act[%f] to be %s, got %s", score1, string(expVal1), actVal[0].Member)
	}
	if actVal[1].Member != string(expVal2) {
		t.Errorf("Expected act[%f] to be %s, got %s", score2, string(expVal2), actVal[1].Member)
	}
}

func TestZRangeNegIdx(t *testing.T) {
	dsStore := datastore.New()
	key, score1, expVal1 := "test", 10.0, []byte("test123")
	score2, expVal2 := 21.1, []byte("care123")
	dsStore.ZAdd(key, score1, expVal1)
	dsStore.ZAdd(key, score2, expVal2)
	actVal, err := dsStore.ZRange(key, 0, -1)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if actVal[0].Member != string(expVal1) {
		t.Errorf("Expected act[%f] to be %s, got %s", score1, string(expVal1), actVal[0].Member)
	}
	if actVal[1].Member != string(expVal2) {
		t.Errorf("Expected act[%f] to be %s, got %s", score2, string(expVal2), actVal[1].Member)
	}
}

func TestTtl(t *testing.T) {
	dsStore := datastore.New()
	key, expVal := "test", []byte("test123")
	dsStore.Set(key, expVal)
	val := dsStore.Ttl(key)
	if val != -1 {
		t.Errorf("Expected err to be -1, got %d", val)
	}
	dsStore.Expire(key, 10)
	val = dsStore.Ttl(key)
	if val == -1 {
		t.Errorf("Expected err to be none -1, got %d", val)
	}
}
