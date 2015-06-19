package config

import "time"

// Config is a configuration interface for Covenant
type Config interface {
	Get(string) interface{}
	GetBool(string) bool
	GetFloat64(string) float64
	GetInt(string) int
	GetString(string) string
	GetStringMap(string) map[string]interface{}
	GetStringMapString(string) map[string]string
	GetStringSlice(string) []string
	GetTime(string) time.Time
	GetDuration(string) time.Duration

	IsSet(string) bool

	SetDefault(string, interface{})
	Set(string, interface{})

	SetConfigName(string)
	AddConfigPath(string)
	ReadInConfig() error

	RegisterAlias(string, string)
}
