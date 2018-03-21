package tracr_cache

import (
	log "github.com/inconshreveable/log15"
	"tracr-daemon/exchanges"
	"tracr-daemon/keys"
	"fmt"
	"strconv"
)

// TODO - enforce client is not null before getting

func (self *CacheClient) GetBalance(exchange, currency string) float64 {
	log.Debug("reading balance", "module", "streams", "exchange", exchange, "currency", currency)
	key := keys.BuildBalancesKey(exchange)
	balance, err := self.client.HGet(key, currency).Float64()
	if err != nil {
		log.Error("error getting balance", "module", "streams", "key", key, "error", err)
	}

	return balance
}

func (self *CacheClient) GetOrderBook(exchange, pair string) exchanges.OrderBook {
	log.Debug("reading order book", "module", "streams", "exchange", exchange, "pair", pair)
	key := keys.BuildOrderBookKey(exchange, pair)
	asksKey := fmt.Sprintf("%s:asks", key)
	bidsKey := fmt.Sprintf("%s:bids", key)

	orderBook := exchanges.NewOrderBook(exchange, pair)

	asks, err := self.client.HGetAll(asksKey).Result()
	if err != nil {
		log.Error("error getting order book asks", "module", "streams", "key", asksKey, "error", err)
	}

	for price, volume := range asks {
		orderBook.Asks[price], _ = strconv.ParseFloat(volume, 64)
	}

	bids, err := self.client.HGetAll(bidsKey).Result()
	if err != nil {
		log.Error("error getting order book bids", "module", "streams", "key", bidsKey, "error", err)
	}

	for price, volume := range bids {
		orderBook.Bids[price], _ = strconv.ParseFloat(volume, 64)
	}

	return *orderBook
}

func (self *CacheClient) GetTicker(exchange, pair string) exchanges.Ticker {
	log.Debug("reading ticker", "module", "streams", "exchange", exchange, "pair", pair)
	panic("please implement")
	//key := fmt.Sprintf("%s-Ticker-%s", exchange, pair)
	//
	//
	//return ticker
}

func (self *CacheClient) GetBotEncoding(botKey string) (string, error) {
	botEncoding, err := self.client.HGet("botStorage", botKey).Result()
	if err != nil {
		log.Error("error getting bot", "module", "streams", "botKey", botKey, "error", err)
		return "", err
	}

	return botEncoding, nil
}
