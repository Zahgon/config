package config

import (
	"github.com/go-viper/mapstructure/v2"
)

// ValDecodeHookFunc returns a mapstructure.DecodeHookFunc
// that parse ENV var, and more custom parse
func ValDecodeHookFunc(parseEnv, parseTime bool) mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}

// https://docs.docker.com/compose/environment-variables/env-file/

// feat: support parse time or duration string. eg: 10s

// resolve format, check is alias
func (c *Config) resolveFormat(f string) string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * Deprecated methods
 *************************************************************/

// SetDecoder add/set a format decoder
//
// Deprecated: please use driver instead
func SetDecoder(format string, decoder Decoder) { _ = "STUB: not implemented"; return }

// SetDecoder set decoder
//
// Deprecated: please use driver instead
func (c *Config) SetDecoder(format string, decoder Decoder) { _ = "STUB: not implemented"; return }

// SetDecoders set decoders
//
// Deprecated: please use driver instead
func (c *Config) SetDecoders(decoders map[string]Decoder) { _ = "STUB: not implemented"; return }

// SetEncoder set a encoder for the format
//
// Deprecated: please use driver instead
func SetEncoder(format string, encoder Encoder) { _ = "STUB: not implemented"; return }

// SetEncoder set a encoder for the format
//
// Deprecated: please use driver instead
func (c *Config) SetEncoder(format string, encoder Encoder) { _ = "STUB: not implemented"; return }

// SetEncoders set encoders
//
// Deprecated: please use driver instead
func (c *Config) SetEncoders(encoders map[string]Encoder) { _ = "STUB: not implemented"; return }

/*************************************************************
 * helper methods/functions
 *************************************************************/

// LoadENVFiles load
// func LoadENVFiles(filePaths ...string) error {
// 	return dotenv.LoadFiles(filePaths...)
// }

// GetEnv get os ENV value by name
func GetEnv(name string, defVal ...string) (val string) { _ = "STUB: not implemented"; return "" }

// Getenv get os ENV value by name. like os.Getenv, but support default value
//
// Notice:
// - Key is not case-sensitive when getting
func Getenv(name string, defVal ...string) (val string) { _ = "STUB: not implemented"; return "" }

func parseVarNameAndType(key string) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// can set var type: int, uint, bool

// if type is not valid and has multi words, as desc message.

// format key
func formatKey(key, sep string) string { _ = "STUB: not implemented"; return "" }
