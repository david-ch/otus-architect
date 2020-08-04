package cache

import (
	"context"
	"encoding/json"
	"github.com/david-ch/otus-architect/product-service/product/model"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

const ttl = 10 * time.Minute
const promotedProductsIdsKey = "promotedProductsIds"
var cacheEnabled = os.Getenv("CACHE_ENABLED") != "false"

var ctx = context.Background()

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}



func GetProduct(id int64) (p model.Product, found bool) {
	if !cacheEnabled {
		found = false
		return
	}

	found = refreshingGet(
		productCacheKey(id),
		&p,
		ttl,
	)

	if found {
		log.Print("DEBUG found in cache - product with id = ", p.ID)
	}

	return
}

func CacheProduct(p model.Product) {
	if !cacheEnabled {
		return
	}

	json, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = rdb.Set(ctx, productCacheKey(p.ID), json, ttl).Result()
	if err != nil {
		log.Println(err)
	} else {
		log.Print("DEBUG cached - product with id = ", p.ID)
	}
}

func EvictProduct(id int64) {
	if !cacheEnabled {
		return
	}

	_, err := rdb.Del(ctx, productCacheKey(id)).Result()
	if err != nil {
		log.Println(err)
	} else {
		log.Print("DEBUG evicted from cache - product with id = ", id)
	}
}


func CachePromoted(ids []int64) {
	if !cacheEnabled {
		return
	}

	json, err := json.Marshal(ids)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = rdb.Set(ctx, promotedProductsIdsKey, json, ttl).Result()
	if err != nil {
		log.Println(err)
	} else {
		log.Print("DEBUG cached - promotedProducts")
	}
}

func GetPromoted() (ids []int64, found bool) {
	if !cacheEnabled {
		found = false
		return
	}

	found = refreshingGet(
		promotedProductsIdsKey,
		&ids,
		ttl,
	)

	if found {
		log.Print("DEBUG found in cache - promotedProducts")
	}

	return
}

func EvictPromoted() {
	if !cacheEnabled {
		return
	}

	_, err := rdb.Del(ctx, promotedProductsIdsKey).Result()
	if err != nil {
		log.Println(err)
	} else {
		log.Print("DEBUG evicted from cache - promotedProducts")
	}
}

func productCacheKey(id int64) string {
	return "product:" + string(id)
}

func refreshingGet(key string, scanTo interface{}, ttl time.Duration) (found bool) {
	found = true

	pipe := rdb.Pipeline()
	pipe.Expire(ctx, key, ttl)
	result := pipe.Get(ctx, key)

	_, err := pipe.Exec(ctx)
	if result.Err() == redis.Nil {
		found = false
		return
	}

	bytes, err := result.Bytes()
	if err != nil {
		log.Println("ERROR while deserializing value from cache", err)
		found = false
		return
	}

	err = json.Unmarshal(bytes, &scanTo)
	if err != nil {
		log.Println("ERROR while deserializing value from cache", err)
		found = false
		return
	}

	return
}
