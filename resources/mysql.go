package resources

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlResource(config *MysqlConfig) (ResourceInterface, error) {
	return &MysqlResource{Config: config}, nil
}

type MysqlResource struct {
	Config *MysqlConfig
	DB     *sql.DB
}

type MysqlConfig struct {
	URI            string
	MaxConnections int
}

func (this *MysqlResource) Get() (interface{}, error) {
	// creating mysql session
	var err error
	this.DB, err = sql.Open("mysql", this.Config.URI)
	if err != nil {
		return nil, err
	}
	this.DB.SetMaxOpenConns(this.Config.MaxConnections)
	this.DB.SetMaxIdleConns(0)

	// start background thread for tracking open db connections
	const INTERVAL = 10
	ticker := time.NewTicker(INTERVAL * time.Second)
	go func(DB *sql.DB) {
		for {
			select {
			case <-ticker.C:
				// check if connections have exceeded max limit
				if DB.Stats().OpenConnections >= this.Config.MaxConnections {
					panic("One of db instance has exceeded allotted maximum connection limit")
					return
				}
			}
		}
	}(this.DB)

	return this.DB, nil
}

func (this *MysqlResource) Close() bool {
	if this.DB != nil {
		this.DB.Close()
		return true
	}
	return false
}
