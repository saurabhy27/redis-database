package mock

import "github.com/saurabhy27/redis-database/model"

type MockDataStore struct {
	GetMocked    bool
	DeleteMocked bool
	ExpireMocked bool
	KeysMocked   bool
	SetMocked    bool
	TtlMocked    bool
	ZAddMocked   bool
	ZRangeMocked bool
}

func (md *MockDataStore) Get(key string) ([]byte, error) {
	md.GetMocked = true
	return []byte("test123"), nil
}

func (mds *MockDataStore) Delete(key string) int {
	mds.DeleteMocked = true
	return 1
}

func (mds *MockDataStore) Expire(key string, seconds int) int {
	mds.ExpireMocked = true
	return 1
}

func (mds *MockDataStore) Keys(key string) ([]string, error) {
	mds.KeysMocked = true
	return []string{"test", "care"}, nil
}

func (mds *MockDataStore) Set(key string, value []byte) {
	mds.SetMocked = true
}

func (mds *MockDataStore) Ttl(key string) int {
	mds.TtlMocked = true
	return 1
}

func (mds *MockDataStore) ZAdd(key string, sorted_set []model.SortedSetByte) (int, error) {
	mds.ZAddMocked = true
	return 2, nil
}

func (mds *MockDataStore) ZRange(key string, start int, stop int) ([]model.SortedSet, error) {
	mds.ZRangeMocked = true
	return []model.SortedSet{{Score: 1, Member: "test123"}}, nil
}
