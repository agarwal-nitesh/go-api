A note from the author:
Recreated/copied it when I was trying to write Java in all languages.


### api
GO based API server (based on go-kit and go-gin).

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


