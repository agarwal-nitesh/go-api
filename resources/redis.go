package resources

import "github.com/go-redis/redis"

func NewRedisResource(config *RedisConfig) (ResourceInterface, error) {
	return &RedisResource{config: config}, nil
}

type RedisResource struct {
	config *RedisConfig
	client *redis.Client
}

type RedisConfig struct {
	Address string
}

func (this *RedisResource) Get() (interface{}, error) {
	// create redis session
	this.client = redis.NewClient(&redis.Options{
		Addr: this.config.Address,
	})

	return this.client, nil
}

func (this *RedisResource) Close() bool {
	if this.client != nil {
		this.client.Close()
		return true
	}
	return false
}
