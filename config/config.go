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
		HttpPort:           config.Get(constanta.HttpPort, "8988"),
		AccountBaseUrl:     config.Get(constanta.AccountBaseUrl, "service-account:6667"),
		TransactionBaseUrl: config.Get(constanta.TransactionBaseUrl, "service-transaction:6668"),
		AuditBaseUrl:       config.Get("AUDIT_BASE_URL", "http://service-audit:8083"),
	}
}
