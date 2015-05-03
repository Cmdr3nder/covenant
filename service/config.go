package service

import (
    "time"

    "github.com/spf13/viper"
)

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

type viperConfig struct {
    v *viper.Viper
}

func (c *viperConfig) Get(key string) interface{} {
    return c.v.Get(key)
}

func (c *viperConfig) GetBool(key string) bool {
    return c.v.GetBool(key)
}

func (c *viperConfig) GetFloat64(key string) float64 {
    return c.v.GetFloat64(key)
}

func (c *viperConfig) GetInt(key string) int {
    return c.v.GetInt(key)
}

func (c *viperConfig) GetString(key string) string {
    return c.v.GetString(key)
}

func (c *viperConfig) GetStringMap(key string) map[string]interface{} {
    return c.v.GetStringMap(key)
}

func (c *viperConfig) GetStringMapString(key string) map[string]string {
    return c.v.GetStringMapString(key)
}

func (c *viperConfig) GetStringSlice(key string) []string {
    return c.v.GetStringSlice(key)
}

func (c *viperConfig) GetTime(key string) time.Time {
    return c.v.GetTime(key)
}

func (c *viperConfig) GetDuration(key string) time.Duration {
    return c.v.GetDuration(key)
}

func (c *viperConfig) IsSet(key string) bool {
    return c.v.IsSet(key)
}

func (c *viperConfig) SetDefault(key string, value interface{}) {
    c.v.SetDefault(key, value)
}

func (c *viperConfig) Set(key string, value interface{}) {
    c.v.Set(key, value)
}

func (c *viperConfig) SetConfigName(name string) {
    c.v.SetConfigName(name)
}

func (c *viperConfig) AddConfigPath(path string) {
    c.v.AddConfigPath(path)
}

func (c *viperConfig) ReadInConfig() error {
    return c.v.ReadInConfig()
}

func (c *viperConfig) RegisterAlias(a string, b string) {
    c.v.RegisterAlias(a, b)
}
