package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

type Database struct {
	Client Repository
}

func newRedis(address string, username string, password string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       0,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: RedisRepo{db: client},
	}, nil
}
func GetRedisClient() *Database {
	r := os.Getenv("REDIS_ADDRESS")
	u := os.Getenv("REDIS_USERNAME")
	p := os.Getenv("REDIS_PASSWORD")
	if r != "" {
		dbRedis, err := newRedis(r, u, p)
		if err != nil {
			println(err.Error(), err)
		}
		return dbRedis
	}
	return nil
}

/*func NewPostgres(address string) (*Database, error) {
	db, err := sql.Open("pgx", address)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{Client: PostgresRepo{db: db}}, nil
}*/
