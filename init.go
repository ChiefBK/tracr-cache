package tracr_cache

import (
	"github.com/go-redis/redis"
	log "github.com/inconshreveable/log15"
	"fmt"
	"errors"
)

func Init() error {
	log.Info("Initializing tracr-cache")

	// initialize connection to redis server using database 7
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       7,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		return errors.New("error sending ping to redis")
	}

	if pong != "PONG" {
		return errors.New(fmt.Sprintf("pinging redis responded with error: %s", pong))
	}

	return nil
}

