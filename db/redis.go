package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const key = "log-manager"

type RedisRepo struct {
	db *redis.Client
}

func (r RedisRepo) SaveFilter(filter string) {
	enc, err := Encrypt(filter, secret)
	if err != nil {
		enc = filter
		fmt.Println(err.Error(), err)
	}
	s, err := r.db.HSet(context.Background(), key, enc, 24*time.Hour).Result()
	if err != nil {
		fmt.Println(s)
	}
}

func (r RedisRepo) LoadFilters() []string {
	enc, err := r.db.HGetAll(context.Background(), key).Result()
	if err != nil {
		fmt.Println(err.Error(), err)
		return []string{}
	}
	result := make([]string, 0, len(enc))

	for key := range enc {
		res, err := Decrypt(key, secret)
		if err != nil {
			res = key
		}
		result = append(result, res)
	}
	return result
}
