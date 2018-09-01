package common

import (
	"encoding/json"
	"flag"
	. "os"
	log "github.com/sirupsen/logrus"
)

type LoggingString struct {
	Data string `json:"data"`
}

type LoggerConfig struct {
	FilePath string
}

func InitializeLogger(config *LoggerConfig) {

	flag.Parse()
	var file, err1 = OpenFile(config.FilePath, O_RDWR|O_CREATE|O_APPEND, 0666)
	if err1 != nil {
		panic(err1)
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
}

func LogJSON(args ...interface{}) {
	for _, arg := range args {
		dataString, _ := json.Marshal(arg)
		log.Println(string(dataString))
	}

}

func LogString(args ...string) {
	for _, arg := range args {
		dataString, _ := json.Marshal(LoggingString{Data: arg})
		log.Println(string(dataString))
	}

}

func LogStringFatal(args ...string) {
	for _, arg := range args {
		dataString, _ := json.Marshal(LoggingString{Data: arg})
		log.Fatalln(string(dataString))
	}

}
