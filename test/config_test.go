package test

import (
	"fmt"
	"github.com/leiphp/gokit/pkg/core/config"
	"testing"
)

type AppConfig struct {
	Port int
}

func TestAppConfig(t *testing.T) {
	var cfg AppConfig
	err := config.LoadConfig("config.yaml", &cfg)
	fmt.Println(cfg.Port)
	fmt.Println(err)
}
