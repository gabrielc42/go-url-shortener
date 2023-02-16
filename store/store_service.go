package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// define struct wrapper around Redis client
type StorageService struct {
	redisClient *redis.Client
}

// top level declarations for storeService and Redis context

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// *note: in a real world usage, cache duration shouldn't have
// an expiration time, an LRU policy config should be set
// where the values that are retrieved less often are purged automatically from
// the cache and stored back in RDBMs whenever cache is full

const CacheDuration = 6 * time.Hour

// initialize store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost: 9808",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v 🔬 ", err))
	}

	fmt.Printf("\n Redis started succesfully: pong message = {%s} 📝 ", pong)
	storeService.redisClient = redisClient
	return storeService
}

/*
we want to be able to save mapping between originalUrl and generated shortUrl
*/
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n 🔧", err, shortUrl, originalUrl))
	}
}

/*
we should be able to retrieve the initial long URL
once the short is provided. This is when users will
be calling the shortlink in the url, so what we need
to do is retrieve the long url and think about redirect
*/
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n 📻 ", err, shortUrl))
	}
	return result
}
