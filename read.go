package config

import (
	"time"

	"github.com/gookit/goutil/maputil"
)

// Exists key exists check
func Exists(key string, findByPath ...bool) bool { _ = "STUB: not implemented"; return false }

// Exists key exists check
func (c *Config) Exists(key string, findByPath ...bool) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// disable find by path.

// has sub key? eg. "lang.dir"

// find top item data based on top key

// is map(from Set)

// is map(from Set)

// is map(decode from toml/json/yaml.v3)

// is map(decode from yaml.v2)

// is array(is from Set)

// is array(is from Set)

// is array(load from file)

// error

/*************************************************************
 * read config data
 *************************************************************/

// Data return all config data
func Data() map[string]any {
	_ = "STUB: not implemented"

	// Data get all config data.
	//
	// Note: will don't apply any options, like ParseEnv
	return nil
}

func (c *Config) Data() map[string]any {
	_ = "STUB: not implemented"

	// Sub return a map config data by key
	return nil
}

func Sub(key string) map[string]any {
	_ = "STUB: not implemented"

	// Sub get a map config data by key
	//
	// Note: will don't apply any options, like ParseEnv
	return nil
}

func (c *Config) Sub(key string) map[string]any { _ = "STUB: not implemented"; return nil }

// Keys return all config data
func Keys() []string {
	_ = "STUB: not implemented"

	// Keys get all config data
	return nil
}

func (c *Config) Keys() []string { _ = "STUB: not implemented"; return nil }

// Get config value by key string, support get sub-value by key path(eg. 'map.key'),
func Get(key string, findByPath ...bool) any { _ = "STUB: not implemented"; return *new(any) }

// Get config value by key, findByPath default is true.
func (c *Config) Get(key string, findByPath ...bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// GetValue get value by given key string. findByPath default is true.
func GetValue(key string, findByPath ...bool) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// GetValue get value by given key string. findByPath default is true.
//
// Return:
//   - ok is true, find value from config
//   - ok is false, not found or error
func (c *Config) GetValue(key string, findByPath ...bool) (value any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// if not is readonly

// is top key

// disable find by path.

// c.addError(ErrNotFound)

// has sub key? eg. "lang.dir"

// c.addError(ErrNotFound)

// find top item data based on top key

// c.addError(ErrNotFound)

// find child
// NOTICE: don't merge case, will result in an error.
// e.g. case []int, []string
// OR
// case []int:
// case []string:

// is map(from Set)

// is map(from Set)

// is map(decode from toml/json)

// is map(decode from yaml)

// is array(is from Set)

// check slice index

// is array(is from Set)

// is array(load from file)

// error

/*************************************************************
 * read config (basic data type)
 *************************************************************/

// String get a string by key
func String(key string, defVal ...string) string { _ = "STUB: not implemented"; return "" }

// String get a string by key, if not found return default value
func (c *Config) String(key string, defVal ...string) string { _ = "STUB: not implemented"; return "" }

// give default value

// MustString get a string by key, will panic on empty or not exists
func MustString(key string) string { _ = "STUB: not implemented"; return "" }

// MustString get a string by key, will panic on empty or not exists
func (c *Config) MustString(key string) string { _ = "STUB: not implemented"; return "" }

func (c *Config) getString(key string) (value string, ok bool) {
	_ = "STUB: not implemented"
	// find from cache
	return "", false
}

// from json `int` always is float64

// add cache

// Int get an int by key
func Int(key string, defVal ...int) int { _ = "STUB: not implemented"; return 0 }

// Int get a int value, if not found return default value
func (c *Config) Int(key string, defVal ...int) (value int) { _ = "STUB: not implemented"; return 0 }

// Uint get a uint value, if not found return default value
func Uint(key string, defVal ...uint) uint { _ = "STUB: not implemented"; return 0 }

// Uint get a int value, if not found return default value
func (c *Config) Uint(key string, defVal ...uint) (value uint) { _ = "STUB: not implemented"; return 0 }

// Int64 get a int value, if not found return default value
func Int64(key string, defVal ...int64) int64 { _ = "STUB: not implemented"; return 0 }

// Int64 get a int value, if not found return default value
func (c *Config) Int64(key string, defVal ...int64) (value int64) {
	_ = "STUB: not implemented"
	return 0
}

// try to get an int64 value by given key
func (c *Config) tryInt64(key string) (value int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Duration get a time.Duration type value. if not found return default value
func Duration(key string, defVal ...time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Duration get a time.Duration type value. if not found return default value
func (c *Config) Duration(key string, defVal ...time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Float get a float64 value, if not found return default value
func Float(key string, defVal ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// Float get a float64 by key
func (c *Config) Float(key string, defVal ...float64) (value float64) {
	_ = "STUB: not implemented"
	return 0
}

// Bool get a bool value, if not found return default value
func Bool(key string, defVal ...bool) bool { _ = "STUB: not implemented"; return false }

// Bool looks up a value for a key in this section and attempts to parse that value as a boolean,
// along with a boolean result similar to a map lookup.
//
// of following(case insensitive):
//   - true
//   - yes
//   - false
//   - no
//   - 1
//   - 0
//
// The `ok` boolean will be false in the event that the value could not be parsed as a bool
func (c *Config) Bool(key string, defVal ...bool) (value bool) {
	_ = "STUB: not implemented"
	return false
}

/*************************************************************
 * read config (complex data type)
 *************************************************************/

// Ints get config data as an int slice/array
func Ints(key string) []int {
	_ = "STUB: not implemented"

	// Ints get config data as an int slice/array
	return nil
}

func (c *Config) Ints(key string) (arr []int) { _ = "STUB: not implemented"; return nil }

// iv, err := strconv.Atoi(fmt.Sprintf("%v", v))

// reset

// IntMap get config data as a map[string]int
func IntMap(key string) map[string]int { _ = "STUB: not implemented"; return nil }

// IntMap get config data as a map[string]int
func (c *Config) IntMap(key string) (mp map[string]int) { _ = "STUB: not implemented"; return nil }

// from Set

// decode from json,toml

// iv, err := strconv.Atoi(fmt.Sprintf("%v", v))

// reset

// if decode from yaml

// iv, err := strconv.Atoi(fmt.Sprintf( "%v", v))

// reset

// sk := fmt.Sprintf("%v", k)

// Strings get strings by key
func Strings(key string) []string { _ = "STUB: not implemented"; return nil }

// Strings get config data as a string slice/array
func (c *Config) Strings(key string) (arr []string) {
	_ = "STUB: not implemented"

	// find from cache
	return nil
}

// arr = append(arr, fmt.Sprintf("%v", v))

// add cache

// StringsBySplit get []string by split a string value.
func StringsBySplit(key, sep string) []string { _ = "STUB: not implemented"; return nil }

// StringsBySplit get []string by split a string value.
func (c *Config) StringsBySplit(key, sep string) (ss []string) {
	_ = "STUB: not implemented"
	return nil
}

// StringMap get config data as a map[string]string
func StringMap(key string) map[string]string { _ = "STUB: not implemented"; return nil }

// StringMap get config data as a map[string]string
func (c *Config) StringMap(key string) (mp map[string]string) {
	_ = "STUB: not implemented"

	// find from cache
	return nil
}

// from Set

// decode from json,toml,yaml.v3

// decode from yaml v2

// add cache

// SubDataMap get sub config data as maputil.Map
func SubDataMap(key string) maputil.Map {
	_ = "STUB: not implemented"
	return *

	// SubDataMap get sub config data as maputil.Map
	//
	// TIP: will not enable parse Env and more
	new(maputil.Map)
}

func (c *Config) SubDataMap(key string) maputil.Map {
	_ = "STUB: not implemented"
	return *new(maputil.Map)
}

// keep is not nil
