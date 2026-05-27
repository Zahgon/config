package config

import (
	"dario.cat/mergo"
	"github.com/go-viper/mapstructure/v2"
)

// there are some event names for config data changed.
const (
	OnSetValue   = "set.value"
	OnSetData    = "set.data"
	OnLoadData   = "load.data"
	OnReloadData = "reload.data"
	OnCleanData  = "clean.data"
)

// HookFunc on config data changed.
type HookFunc func(event string, c *Config)

// Options config options
type Options struct {
	// ParseEnv parse env in string value and default value. default: false
	//
	//  - like: "${EnvName}" "${EnvName|default}"
	ParseEnv bool
	// ParseTime parses a duration string to `time.Duration`. default: false
	//
	// eg: 10s, 2m
	ParseTime bool
	// ParseDefault tag on binding data to struct. default: false
	//
	//  - tag: default
	//
	// NOTE: If you want to parse a substruct, you need to set the `default:""` flag on the struct,
	// otherwise the fields that will not resolve to it will not be resolved.
	ParseDefault bool
	// Readonly config is readonly. default: false
	Readonly bool
	// EnableCache enable config data cache. default: false
	EnableCache bool
	// ParseKey support key path, allow finding value by key path. default: true
	//
	// - eg: 'key.sub' will find `map[key]sub`
	ParseKey bool
	// TagName tag name for binding data to struct
	//
	// Deprecated: please set tag name by DecoderConfig, or use SetTagName()
	TagName string
	// Delimiter the delimiter char for split key path, on `ParseKey=true`.
	//
	// - default is '.'
	Delimiter byte
	// DumpFormat default write format. default is 'json'
	DumpFormat string
	// ReadFormat default input format. default is 'json'
	ReadFormat string
	// DecoderConfig setting for binding data to struct. such as: TagName
	DecoderConfig *mapstructure.DecoderConfig
	// MergeOptions settings for merge two data
	MergeOptions []func(*mergo.Config)
	// HookFunc on data changed. you can do something...
	HookFunc HookFunc
	// WatchChange bool
}

// OptionFn option func
type OptionFn func(*Options)

func newDefaultOption() *Options { _ = "STUB: not implemented"; return nil }

// for export

// struct decoder config

func newDefaultDecoderConfig(tagName string) *mapstructure.DecoderConfig {
	_ = "STUB: not implemented"
	return nil
}

// tag name for binding struct

// will auto convert string to int/uint

// SetTagName for mapping data to struct
func (o *Options) SetTagName(tagName string) { _ = "STUB: not implemented"; return }

func (o *Options) shouldAddHookFunc() bool { _ = "STUB: not implemented"; return false }

func (o *Options) makeDecoderConfig() *mapstructure.DecoderConfig {
	_ = "STUB: not implemented"
	return nil
}

// copy new config for each binding.

// compatible with previous settings opts.TagName

// add hook on decode value to struct

/*************************************************************
 * config setting
 *************************************************************/

// WithTagName set tag name for export to struct
func WithTagName(tagName string) func(*Options) { _ = "STUB: not implemented"; return nil }

// ParseEnv set parse env value
func ParseEnv(opts *Options) { _ = "STUB: not implemented"; return }

// ParseTime set parse time string.
func ParseTime(opts *Options) { _ = "STUB: not implemented"; return }

// ParseDefault tag value on binding data to struct.
func ParseDefault(opts *Options) { _ = "STUB: not implemented"; return }

// Readonly set readonly
func Readonly(opts *Options) { _ = "STUB: not implemented"; return }

// Delimiter set delimiter char
func Delimiter(sep byte) func(*Options) { _ = "STUB: not implemented"; return nil }

// SaveFileOnSet set hook func, will panic on save error
func SaveFileOnSet(fileName string, format string) func(options *Options) {
	_ = "STUB: not implemented"
	return nil
}

// WithHookFunc set hook func
func WithHookFunc(fn HookFunc) func(*Options) { _ = "STUB: not implemented"; return nil }

// EnableCache set readonly
func EnableCache(opts *Options) { _ = "STUB: not implemented"; return }

// WithOptions with options
func WithOptions(opts ...OptionFn) { _ = "STUB: not implemented"; return }

// WithOptions apply some options
func (c *Config) WithOptions(opts ...OptionFn) *Config { _ = "STUB: not implemented"; return nil }

// apply options

// GetOptions get options
func GetOptions() *Options {
	_ = "STUB: not implemented"

	// Options get
	return nil
}

func (c *Config) Options() *Options {
	_ = "STUB: not implemented"

	// With apply some options
	return nil
}

func (c *Config) With(fn func(c *Config)) *Config {
	_ = "STUB: not implemented"

	// Readonly disable set data to config.
	//
	// Usage:
	//
	//	config.LoadFiles(a, b, c)
	//	config.Readonly()
	return nil
}

func (c *Config) Readonly() { _ = "STUB: not implemented"; return }
