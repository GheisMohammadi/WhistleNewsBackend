package config_test

import (
	"WhistleNewsBackend/config"
	"testing"
)

func TestConfigs(t *testing.T) {
	dbName := config.DataBaseName
	if len(dbName) <= 0 {
		t.Error("wrong config")
	}
}