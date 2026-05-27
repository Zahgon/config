package config

import (
	"errors"
)

// some common errors definitions
var (
	ErrReadonly   = errors.New("the config instance in 'readonly' mode")
	ErrKeyIsEmpty = errors.New("the config key is cannot be empty")
	ErrNotFound   = errors.New("this key does not exist in the config")
)

// SetData for override the Config.Data
func SetData(data map[string]any) {
	_ = "STUB: not implemented"

	// SetData for override the Config.Data
	return
}

func (c *Config) SetData(data map[string]any) { _ = "STUB: not implemented"; return }

// Set value by key. setByPath default is true
func Set(key string, val any, setByPath ...bool) error { _ = "STUB: not implemented"; return nil }

// Set a value by key string. setByPath default is true
func (c *Config) Set(key string, val any, setByPath ...bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// disable set by path.

// set by path
