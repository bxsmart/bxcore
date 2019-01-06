package cache

import (
	"fmt"
	myredis "github.com/bxsmart/bxcore/cache/redis"
)

var cacheMap = make(map[int]Cache, 0)

var cache Cache

type Cache interface {
	Set(key string, value []byte, ttl int64) error

	Get(key string) ([]byte, error)

	Del(key string) error

	Dels(keys []string) error

	Exists(key string) (bool, error)

	Keys(keyFormat string) ([][]byte, error)

	HMSet(key string, ttl int64, args ...[]byte) error

	HMGet(key string, fields ...[]byte) ([][]byte, error)

	HDel(key string, fields ...[]byte) (int64, error)

	HGetAll(key string) ([][]byte, error)

	HVals(key string) ([][]byte, error)

	HExists(key string, field []byte) (bool, error)

	SAdd(key string, ttl int64, members ...[]byte) error

	SCard(key string) (int64, error)

	SRem(key string, members ...[]byte) (int64, error)

	SMembers(key string) ([][]byte, error)

	SIsMember(key string, member []byte) (bool, error)

	ZAdd(key string, ttl int64, args ...[]byte) error

	Incr(key string) (int64, error)

	ExpireAt(key string, expireAt int64) error

	ZRange(key string, start, stop int64, withScores bool) ([][]byte, error)

	ZRemRangeByScore(key string, start, stop int64) (int64, error)

	ZRem(key string, members ...[]byte) (int64, error)

	LPush(key string, ttl int64, args ...[]byte) error

	Publish(channel string, members ...[]byte) (int64, error)

	SScan(key string, count int64) ([][]byte, error)
}

func NewCache(cfg interface{}) Cache {
	options := cfg.(myredis.RedisOptions)
	if _cache, ok := cacheMap[options.Database]; ok {
		return _cache
	}

	redisCache := &myredis.RedisCacheImpl{}
	redisCache.Initialize(cfg)

	cacheMap[redisCache.Database()] = redisCache

	cache = redisCache
}

func GetCache(Db int) (Cache, error) {
	if cache, ok := cacheMap[Db]; ok {
		return cache, nil
	}

	return nil, fmt.Errorf("cache %v not existed", Db)
}

func Set(key string, value []byte, ttl int64) error { return cache.Set(key, value, ttl) }

func Get(key string) ([]byte, error) { return cache.Get(key) }

func Del(key string) error { return cache.Del(key) }

func Dels(keys []string) error { return cache.Dels(keys) }

func Exists(key string) (bool, error) { return cache.Exists(key) }

func Keys(keyFormat string) ([][]byte, error) { return cache.Keys(keyFormat) }

func HMSet(key string, ttl int64, args ...[]byte) error {
	return cache.HMSet(key, ttl, args...)
}

func HMGet(key string, fields ...[]byte) ([][]byte, error) {
	return cache.HMGet(key, fields...)
}

func HGetAll(key string) ([][]byte, error) {
	return cache.HGetAll(key)
}

func HVals(key string) ([][]byte, error) {
	return cache.HVals(key)
}

func HExists(key string, field []byte) (bool, error) {
	return cache.HExists(key, field)
}
func SAdd(key string, ttl int64, members ...[]byte) error {
	return cache.SAdd(key, ttl, members...)
}
func SCard(key string) (int64, error) {
	return cache.SCard(key)
}

func SScan(key string, count int64) ([][]byte, error) {
	return cache.SScan(key, count)
}

func SMembers(key string) ([][]byte, error) {
	return cache.SMembers(key)
}

func SIsMember(key string, member []byte) (bool, error) {
	return cache.SIsMember(key, member)
}

func SRem(key string, members ...[]byte) (int64, error) {
	return cache.SRem(key, members...)
}

func HDel(key string, fields ...[]byte) (int64, error) {
	return cache.HDel(key, fields...)
}

func ZAdd(key string, ttl int64, args ...[]byte) error {
	return cache.ZAdd(key, ttl, args...)
}

func ZRange(key string, start, stop int64, withScores bool) ([][]byte, error) {
	return cache.ZRange(key, start, stop, withScores)
}
func ZRemRangeByScore(key string, start, stop int64) (int64, error) {
	return cache.ZRemRangeByScore(key, start, stop)
}

func ZRem(key string, members ...[]byte) (int64, error) {
	return cache.ZRem(key, members...)
}

func Publish(channel string, members ...[]byte) (int64, error) {
	return cache.Publish(channel, members...)
}

func Incr(key string) (int64, error) {
	return cache.Incr(key)
}

func ExpireAt(key string, expireAt int64) error {
	return cache.ExpireAt(key, expireAt)
}

func LPush(key string, ttl int64, args ...[]byte) error {
	return cache.LPush(key, ttl, args...)
}

func IsInit() bool {
	return nil != cache
}
