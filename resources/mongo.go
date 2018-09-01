package resources

import (
	"gopkg.in/mgo.v2"
)

func NewMongoResource(config *MongoConfig) (ResourceInterface, error) {
	return &MongoResource{config: config}, nil
}

type MongoResource struct {
	config  *MongoConfig
	session *mgo.Session
}

type MongoConfig struct {
	URI string
}

func (this *MongoResource) Get() (interface{}, error) {
	// creating mongo session
	var err error
	this.session, err = mgo.Dial(this.config.URI)
	if err != nil {
		return nil, err
	}
	this.session.SetMode(mgo.Monotonic, true)

	return this.session.DB("e_corp_api"), nil
}

func (this *MongoResource) Close() bool {
	if this.session != nil {
		this.session.Close()
		return true
	}
	return false
}
