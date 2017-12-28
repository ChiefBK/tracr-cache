package tracr_cache

import (
	"fmt"
	log "github.com/inconshreveable/log15"
	"tracr-daemon/exchanges"
	"github.com/go-redis/redis"
)

var client *redis.Client

func PutBalances(key string, balances exchanges.Balances) {
	for currency, balance := range balances {
		err := client.HSet(key, currency, balance).Err()
		if err != nil {
			log.Error("error setting balance", "module", "streams", "currency", currency, "balance", balance)
		}
	}
}

func PutTicker(key string, ticker exchanges.Ticker) {
	if ticker.LastTrade != nil {
		err := client.HSet(key, "lastTrade", ticker.LastTrade).Err()
		if err != nil {
			log.Error("error setting last trade for ticker", "module", "streams", "key", key)
		}
	}

	if ticker.HighestBid != nil {
		err := client.HSet(key, "highestBit", ticker.HighestBid).Err()
		if err != nil {
			log.Error("error setting highest bid for ticker", "module", "streams", "key", key)
		}
	}

	if ticker.LowestAsk != nil {
		err := client.HSet(key, "lowestAsk", ticker.LowestAsk).Err()
		if err != nil {
			log.Error("error setting lowest ask for ticker", "module", "streams", "key", key)
		}
	}

	if ticker.TwentyFourHourHigh != nil {
		err := client.HSet(key, "twentyFourHourHigh", ticker.TwentyFourHourHigh).Err()
		if err != nil {
			log.Error("error setting 24 hour high for ticker", "module", "streams", "key", key)
		}
	}

	if ticker.TwentyFourHourLow != nil {
		err := client.HSet(key, "twentyFourHourHigh", ticker.TwentyFourHourLow).Err()
		if err != nil {
			log.Error("error setting 24 hour low for ticker", "module", "streams", "key", key)
		}
	}
}

func PutOrderBook(key string, orderBook exchanges.OrderBook) {
	asksKey := fmt.Sprintf("%s:asks", key)
	bidsKey := fmt.Sprintf("%s:bids", key)

	for price, volume := range orderBook.Asks {
		err := client.HSet(asksKey, price, volume).Err()
		if err != nil {
			log.Error("error setting ask for order book", "module", "streams", "key", key)
		}
	}

	for price, volume := range orderBook.Bids {
		err := client.HSet(bidsKey, price, volume).Err()
		if err != nil {
			log.Error("error setting bid for order book", "module", "streams", "key", key)
		}
	}
}
