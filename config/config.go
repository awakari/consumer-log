package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Api struct {
		Port uint16 `envconfig:"API_PORT" default:"8080" required:"true"`
	}
	Log struct {
		Level int `envconfig:"LOG_LEVEL" default:"-4" required:"true"`
	}
	Queue QueueConfig
}

type QueueConfig struct {
	BatchSize uint32 `envconfig:"QUEUE_BATCH_SIZE" default:"100" required:"true"`
	FallBack  struct {
		Enabled bool   `envconfig:"QUEUE_FALLBACK_ENABLED" default:"true" required:"false"`
		Suffix  string `envconfig:"QUEUE_FALLBACK_SUFFIX" default:"fallback" required:"true""`
	}
	Limit              uint32 `envconfig:"QUEUE_LIMIT" default:"1000" required:"true"`
	Name               string `envconfig:"QUEUE_NAME" default:"consumer-log" required:"true"`
	SleepOnEmptyMillis uint32 `envconfig:"QUEUE_SLEEP_ON_EMPTY_MILLIS" default:"1000" required:"true"`
	SleepOnErrorMillis uint32 `envconfig:"QUEUE_SLEEP_ON_ERROR_MILLIS" default:"1000" required:"true"`
	Uri                string `envconfig:"QUEUE_URI" default:"queue:8080" required:"true"`
}

func NewConfigFromEnv() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return
}
