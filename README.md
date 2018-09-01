### api
GO based API Gateway (based on go-kit and go-gin)

### Installation

#### go
```
Install go and configure GOROOT & GOPATH
https://golang.org/doc/install
```

#### project setup
```
brew install dep
sudo mkdir -p /var/log/application
sudo touch /var/log/application/e_corp_api.log
sudo chmod -R 777 /var/log/application/e_corp_api.log
sh .bin/build_run.sh develop
```

#### swagger setup
```
swag init
```

#### documentation
```
Each service is initialized with resources it needs and consists of:
-service (logical implementation)
-transport (communication HTTP)
-transportRoutes (Aggregation of all transports).
There is an additional default implementation of transport which implements:
DecodeRequest and EncodeResponse.
Utils and Dao are additional layer added to service to add more structure and code familiarity.
Errors and logging are subtly handled by go error and typecasting at transport level. Errors may also be logged at the time of creation.
```

#### TODO
```
Add unit tests
Add db migration tool
Add ratelimit
Add package metrics
Add tracing
```
