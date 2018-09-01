package resources

import (
	"github.com/parnurzeal/gorequest"
)

func NewRequestsResource() (ResourceInterface, error) {
	return &RequestsResource{}, nil
}

type RequestsResource struct {
	requests *gorequest.SuperAgent
}

func (this *RequestsResource) Get() (interface{}, error) {
	this.requests = gorequest.New()

	return this.requests, nil
}

func (this *RequestsResource) Close() bool {
	this.requests = nil
	return true
}
