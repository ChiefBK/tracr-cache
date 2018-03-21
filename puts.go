package tracr_cache

import (
	"fmt"
	log "github.com/inconshreveable/log15"
	"tracr-daemon/exchanges"
)

// TODO - enforce client is not null before putting

func (self *CacheClient) PutBalances(key string, balances exchanges.Balances) {
	for currency, balance := range balances {
		err := self.client.HSet(key, currency, balance).Err()
		if err != nil {
			log.Error("error setting balance", "module", "streams", "currency", currency, "balance", balance, "error", err)
		}
	}
}

func (self *CacheClient) PutTicker(key string, ticker exchanges.Ticker) {
	if ticker.LastTrade != nil {
		err := self.client.HSet(key, "lastTrade", *ticker.LastTrade).Err()
		if err != nil {
			log.Error("error setting last trade for ticker", "module", "streams", "key", key, "error", err)
		}
	}

	if ticker.HighestBid != nil {
		err := self.client.HSet(key, "highestBit", *ticker.HighestBid).Err()
		if err != nil {
			log.Error("error setting highest bid for ticker", "module", "streams", "key", key, "error", err)
		}
	}

	if ticker.LowestAsk != nil {
		err := self.client.HSet(key, "lowestAsk", *ticker.LowestAsk).Err()
		if err != nil {
			log.Error("error setting lowest ask for ticker", "module", "streams", "key", key, "error", err)
		}
	}

	if ticker.TwentyFourHourHigh != nil {
		err := self.client.HSet(key, "twentyFourHourHigh", *ticker.TwentyFourHourHigh).Err()
		if err != nil {
			log.Error("error setting 24 hour high for ticker", "module", "streams", "key", key, "error", err)
		}
	}

	if ticker.TwentyFourHourLow != nil {
		err := self.client.HSet(key, "twentyFourHourHigh", *ticker.TwentyFourHourLow).Err()
		if err != nil {
			log.Error("error setting 24 hour low for ticker", "module", "streams", "key", key, "error", err)
		}
	}
}

func (self *CacheClient) PutOrderBook(key string, orderBook exchanges.OrderBook) {
	asksKey := fmt.Sprintf("%s:asks", key)
	bidsKey := fmt.Sprintf("%s:bids", key)

	for price, volume := range orderBook.Asks {
		err := self.client.HSet(asksKey, price, volume).Err()
		if err != nil {
			log.Error("error setting ask for order book", "module", "streams", "key", key, "error", err)
		}
	}

	for price, volume := range orderBook.Bids {
		err := self.client.HSet(bidsKey, price, volume).Err()
		if err != nil {
			log.Error("error setting bid for order book", "module", "streams", "key", key, "error", err)
		}
	}
}

func (self *CacheClient) PutBotEncoding(botKey, botEncoding string) {
	_, err := self.client.HSet("botStorage", botKey, botEncoding).Result()
	if err != nil {
		log.Error("error setting bot", "module", "streams", "botKey", botKey, "error", err)
	}
}
