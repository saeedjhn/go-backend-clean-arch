package setuptest

import "github.com/saeedjhn/go-domain-driven-design/pkg/persistance/cache/redis"

func NewRedisDB(config redis.Config) (*redis.DB, error) {
	db := redis.New(config)
	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}
