package main

import (
	"bytes"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// DefaultValues is the default configuration
const DefaultValues = `
[Synchronizer]
SyncInterval = "0s"
SyncChunkSize = 100
HackatonAddr = ""
URL = "http://localhost"
PrivateKeyPath = "./test/test.keystore"
PrivateKeyPassword = "testonly"
GenBlockNumber = 0
`

// Duration is a wrapper type that parses time duration from text.
type Duration struct {
	time.Duration `validate:"required"`
}

// UnmarshalText unmarshalls time duration from text.
func (d *Duration) UnmarshalText(data []byte) error {
	duration, err := time.ParseDuration(string(data))
	if err != nil {
		return err
	}
	d.Duration = duration
	return nil
}

// Config represents the configuration of the synchronizer
type Config struct {
	// SyncInterval is the delay interval between reading new rollup information
	SyncInterval Duration `mapstructure:"SyncInterval"`

	// SyncChunkSize is the number of blocks to sync on each chunk
	SyncChunkSize uint64 `mapstructure:"SyncChunkSize"`

	// HackatonAddr is the address of the hackaton smart contract
	HackatonAddr common.Address

	// URL of ethereum node
	URL string `mapstructure:"URL"`

	PrivateKeyPath     string `mapstructure:"PrivateKeyPath"`
	PrivateKeyPassword string `mapstructure:"PrivateKeyPassword"`
	GenBlockNumber     uint64 `mapstructure:"GenBlockNumber"`
}

// Load loads the configuration
func Load() (*Config, error) {
	var cfg Config
	viper.SetConfigType("toml")

	err := viper.ReadConfig(bytes.NewBuffer([]byte(DefaultValues)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc()))
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
