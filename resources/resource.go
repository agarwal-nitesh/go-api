package resources

import (
	"github.com/therudite/api/errors"
	"github.com/therudite/api/config"
)

// represents a resource
type ResourceInterface interface {
	Get() (interface{}, error)
	Close() bool
}

type ResourceManagerInterface interface {
	Get(string) (interface{}, error)
	Close()
}

// manages access to resource
type ResourceManagerImpl struct {
	registry map[string]ResourceInterface
}

// returns api/resource.ResourceManager
func NewResourceManager(cfg config.ConfigManager) (ResourceManagerInterface, error) {

	var err error

	mongoconfig := new(MongoConfig)
	err = cfg.Load("mongo", mongoconfig)
	var mongoresource ResourceInterface
	mongoresource, err = NewMongoResource(mongoconfig)
	if err != nil {
		return nil, err
	}

	mysqlconfig := new(MysqlConfig)
	err = cfg.Load("mysql", mysqlconfig)
	var mysqlresource ResourceInterface
	mysqlresource, err = NewMysqlResource(mysqlconfig)
	if err != nil {
		return nil, err
	}

	redisconfig := new(RedisConfig)
	err = cfg.Load("redis", redisconfig)
	var redisresource ResourceInterface
	redisresource, err = NewRedisResource(redisconfig)
	if err != nil {
		return nil, err
	}

	var requestsresource ResourceInterface
	requestsresource, err = NewRequestsResource()
	if err != nil {
		return nil, err
	}

	registry := map[string]ResourceInterface{
		"mongo":    mongoresource,
		"mysql":    mysqlresource,
		"redis":    redisresource,
		"requests": requestsresource,
	}
	return &ResourceManagerImpl{registry: registry}, nil
}

// resource accessor for managed resources
func (this *ResourceManagerImpl) Get(name string) (interface{}, error) {
	if resource, ok := this.registry[name]; ok {
		res, err := resource.Get()
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		// return resource not found error
		return nil, errors.ResourceNotFound
	}
}

// closing all the managed resources
func (this *ResourceManagerImpl) Close() {
	for _, resource := range this.registry {
		resource.Close()
	}
}
