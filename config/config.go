package config

import (
	"github.com/wahyunurdian26/util/config"
	"github.com/wahyunurdian26/util/constanta"
)

type Config struct {
	HttpPort           string
	AccountBaseUrl     string
	TransactionBaseUrl string
	AuditBaseUrl       string
}

func LoadConfigs() Config {
	return Config{
		HttpPort:           config.Get(constanta.HttpPort, "8081"),
		AccountBaseUrl:     config.Get(constanta.AccountBaseUrl, "127.0.0.1:6667"),
		TransactionBaseUrl: config.Get(constanta.TransactionBaseUrl, "127.0.0.1:6668"),
		AuditBaseUrl:       config.Get("AUDIT_BASE_URL", "http://127.0.0.1:8083"),
	}
}
