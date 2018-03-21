package tracr_cache

import (
	"github.com/go-redis/redis"
	log "github.com/inconshreveable/log15"
	"fmt"
	"errors"
)

type CacheClient struct {
	client *redis.Client
}

func NewCacheClient() (*CacheClient, error) {
	log.Info("creating cache client", "module", "cache")

	// initialize connection to redis server using database 7
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       7,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		return nil, errors.New("error sending ping to redis")
	}

	if pong != "PONG" {
		return nil, errors.New(fmt.Sprintf("pinging redis responded with error: %s", pong))
	}

	cacheClient := &CacheClient{
		client: client,
	}

	return cacheClient, nil
}

func (self *CacheClient) ClearCache() error {
	err := self.client.FlushDB().Err()

	if err != nil {
		return err
	}

	return nil
}
