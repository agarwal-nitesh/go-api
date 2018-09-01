package config

import (
	"path/filepath"
	"runtime"

	"gopkg.in/ini.v1"
)

type ConfigManager interface {
	Load(section string, tpl interface{}) error
}

type configManagerImpl struct {
	config *ini.File
}

// returns api/config.ConfigManager
func NewConfigManager() (ConfigManager, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	filename = filepath.Join(filepath.Dir(filename), "config.ini")
	config, err := ini.InsensitiveLoad(filename)
	if err != nil {
		return nil, err
	}

	return &configManagerImpl{config: config}, nil
}

// loads section of ini file into struct
func (this *configManagerImpl) Load(section string, tpl interface{}) error {
	return this.config.Section(section).MapTo(tpl)
}
