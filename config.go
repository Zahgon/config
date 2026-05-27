/*
Package config is a go config management implement. support YAML,TOML,JSON,INI,HCL format.

Source code and other details for the project are available at GitHub:

	https://github.com/gookit/config

JSON format content example:

	{
		"name": "app",
		"debug": false,
		"baseKey": "value",
		"age": 123,
		"envKey": "${SHELL}",
		"envKey1": "${NotExist|defValue}",
		"map1": {
			"key": "val",
			"key1": "val1",
			"key2": "val2"
		},
		"arr1": [
			"val",
			"val1",
			"val2"
		],
		"lang": {
			"dir": "res/lang",
			"defLang": "en",
			"allowed": {
				"en": "val",
				"zh-CN": "val2"
			}
		}
	}

Usage please see example(more example please see examples folder in the lib):
*/
package config

import (
	"sync"
)

// There are supported config format
const (
	Ini  = "ini"
	Hcl  = "hcl"
	Yml  = "yml"
	JSON = "json"
	Yaml = "yaml"
	Toml = "toml"
	Prop = "properties"
)

const (
	// default delimiter
	defaultDelimiter byte = '.'
	// default struct tag name for binding data to struct
	defaultStructTag = "mapstructure"
	// struct tag name for set default-value on binding data
	defaultValueTag = "default"
)

// internal vars
// type intArr []int
type strArr []string

// type intMap map[string]int
type strMap map[string]string

// This is a default config manager instance
var dc = New("default")

// Config structure definition
type Config struct {
	// save the latest error, will clear after read.
	err error
	// config instance name
	name string
	lock sync.RWMutex

	// config options
	opts *Options
	// all config data
	data map[string]any

	// loaded config files records
	loadedUrls  []string
	loadedFiles []string
	driverNames []string
	// driver alias to name map.
	aliasMap  map[string]string
	reloading bool

	// TODO Deprecated decoder and encoder, use driver instead
	// drivers map[string]Driver

	// decoders["toml"] = func(blob []byte, v any) (err error){}
	// decoders["yaml"] = func(blob []byte, v any) (err error){}
	decoders map[string]Decoder
	encoders map[string]Encoder

	// cache on got config data
	intCache map[string]int
	strCache map[string]string
	// iArrCache map[string]intArr TODO cache it
	// iMapCache map[string]intMap
	sArrCache map[string]strArr
	sMapCache map[string]strMap
}

// New config instance with custom options, default with JSON driver
func New(name string, opts ...OptionFn) *Config { _ = "STUB: not implemented"; return nil }

// NewGeneric create generic config instance with custom options.
//
//   - default add options: ParseEnv, ParseDefault, ParseTime
func NewGeneric(name string, opts ...OptionFn) *Config { _ = "STUB: not implemented"; return nil }

// NewEmpty create config instance with custom options
func NewEmpty(name string, opts ...OptionFn) *Config { _ = "STUB: not implemented"; return nil }

// don't add any drivers

// NewWith create config instance, and you can call some init func
func NewWith(name string, fn func(c *Config)) *Config { _ = "STUB: not implemented"; return nil }

// NewWithOptions config instance. alias of New()
func NewWithOptions(name string, opts ...OptionFn) *Config { _ = "STUB: not implemented"; return nil }

// Default get the default instance
func Default() *Config {
	_ = "STUB: not implemented"

	/*************************************************************
	 * config drivers
	 *************************************************************/return nil
}

// WithDriver set multi drivers at once.
func WithDriver(drivers ...Driver) { _ = "STUB: not implemented"; return }

// WithDriver set multi drivers at once.
func (c *Config) WithDriver(drivers ...Driver) *Config { _ = "STUB: not implemented"; return nil }

// AddDriver set a decoder and encoder driver for a format.
func AddDriver(driver Driver) { _ = "STUB: not implemented"; return }

// AddDriver set a decoder and encoder driver for a format.
func (c *Config) AddDriver(driver Driver) { _ = "STUB: not implemented"; return }

// HasDecoder has decoder
func (c *Config) HasDecoder(format string) bool { _ = "STUB: not implemented"; return false }

// HasEncoder has encoder
func (c *Config) HasEncoder(format string) bool { _ = "STUB: not implemented"; return false }

// DelDriver delete driver of the format
func (c *Config) DelDriver(format string) { _ = "STUB: not implemented"; return }

/*************************************************************
 * helper methods
 *************************************************************/

// Name get config name
func (c *Config) Name() string {
	_ = "STUB: not implemented"

	// AddAlias add alias for a format(driver name)
	return ""
}

func AddAlias(format, alias string) { _ = "STUB: not implemented"; return }

// AddAlias add alias for a format(driver name)
//
// Example:
//
//	config.AddAlias("ini", "conf")
func (c *Config) AddAlias(format, alias string) { _ = "STUB: not implemented"; return }

// AliasMap get alias map
func (c *Config) AliasMap() map[string]string {
	_ = "STUB: not implemented"

	// Error get last error, will clear after read.
	return nil
}

func (c *Config) Error() error { _ = "STUB: not implemented"; return nil }

// IsEmpty of the config
func (c *Config) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// LoadedUrls get loaded urls list
func (c *Config) LoadedUrls() []string {
	_ = "STUB: not implemented"

	// LoadedFiles get loaded files name
	return nil
}

func (c *Config) LoadedFiles() []string { _ = "STUB: not implemented"; return nil }

// DriverNames get loaded driver names
func (c *Config) DriverNames() []string { _ = "STUB: not implemented"; return nil }

// Reset data and caches
func Reset() {
	_ = "STUB: not implemented"

	// ClearAll data and caches
	return
}

func ClearAll() {
	_ = "STUB: not implemented"

	// ClearAll data and caches
	return
}

func (c *Config) ClearAll() { _ = "STUB: not implemented"; return }

// options

// ClearData clear data
func (c *Config) ClearData() { _ = "STUB: not implemented"; return }

// ClearCaches clear caches
func (c *Config) ClearCaches() { _ = "STUB: not implemented"; return }

/*************************************************************
 * helper methods
 *************************************************************/

// fire hook
func (c *Config) fireHook(name string) { _ = "STUB: not implemented"; return }

// record error
func (c *Config) addError(err error) {
	_ = "STUB: not implemented"

	// format and record error
	return
}

func (c *Config) addErrorf(format string, a ...any) { _ = "STUB: not implemented"; return }
