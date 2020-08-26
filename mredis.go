package mredis

import "github.com/go-redis/redis"

// NewRedis 初始化连接
func NewRedis(addr, passwd string, n int) (*redis.Client, error) {
	var conn *redis.Client
	if n <= 15 && n >= 0 {
		conn = redis.NewClient(&redis.Options{
			Addr:     addr,   // "localhost:6379"
			Password: passwd, // no password set ""
			DB:       n,      // use default DB
		})
		//defer conn.Close()
		_, err := conn.Ping().Result()
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}
