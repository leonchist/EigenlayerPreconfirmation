package config

import (
	"log"

	gcfg "gopkg.in/gcfg.v1"
)

var (
	env = "development"
)

type MainConfig struct {
	Server struct {
		Port     string
		BasePath string
	}

	Database struct {
		SlaveDSN      string
		MasterDSN     string
		RetryInterval int
		MaxIdleConn   int
		MaxConn       int
	}
}

func ReadConfig(cfg interface{}, module string) interface{} {
	ok := ReadModuleConfig(cfg, "/etc", module) || 
	ReadModuleConfig(cfg, "files/etc", module) || 
	ReadModuleConfig(cfg, "../../files/etc", module) || 
	ReadModuleConfig(cfg, "../../../files/etc", module) || 
	ReadModuleConfig(cfg, "/home/app/etc/", module) || 
	ReadModuleConfig(cfg, "/opt/practice-go-unit-test/files/etc", module)

	if !ok {
		log.Fatalln("failed to read config for ", module)
	}
	return cfg
}


func ReadModuleConfig(cfg interface{}, path string, module string) bool {
	environ := env
	fname := path + "/" + module + "." + environ + ".ini"
	err := gcfg.ReadFileInto(cfg, fname)
	if err == nil {
		return true
	}

	return false
}

func ReadModuleConfigWithErr(cfg interface{}, path string, module string) error {
	environ := env
	fname := path + "/" + module + "." + environ + ".ini"
	return gcfg.ReadFileInto(cfg, fname)
}
