package config

import (
	"io"
)

// Decode all config data to the dst ptr
//
// Usage:
//
//	myConf := &MyConf{}
//	config.Decode(myConf)
func Decode(dst any) error { _ = "STUB: not implemented"; return nil }

// Decode all config data to the dst ptr.
//
// It's equals:
//
//	c.Structure("", dst)
func (c *Config) Decode(dst any) error { _ = "STUB: not implemented"; return nil }

// MapStruct alias method of the 'Structure'
//
// Usage:
//
//	dbInfo := &Db{}
//	config.MapStruct("db", dbInfo)
func MapStruct(key string, dst any) error { _ = "STUB: not implemented"; return nil }

// MapStruct alias method of the 'Structure'
func (c *Config) MapStruct(key string, dst any) error { _ = "STUB: not implemented"; return nil }

// BindStruct alias method of the 'Structure'
func BindStruct(key string, dst any) error { _ = "STUB: not implemented"; return nil }

// BindStruct alias method of the 'Structure'
func (c *Config) BindStruct(key string, dst any) error { _ = "STUB: not implemented"; return nil }

// MapOnExists mapping data to the dst structure only on key exists.
func MapOnExists(key string, dst any) error { _ = "STUB: not implemented"; return nil }

// MapOnExists mapping data to the dst structure only on key exists.
//
//   - Support ParseEnv on mapping
//   - Support ParseDefault on mapping
func (c *Config) MapOnExists(key string, dst any) error { _ = "STUB: not implemented"; return nil }

// Structure get config data and binding to the dst structure.
//
//   - Support ParseEnv on mapping
//   - Support ParseDefault on mapping
//
// Usage:
//
//	dbInfo := Db{}
//	config.Structure("db", &dbInfo)
func (c *Config) Structure(key string, dst any) (err error) {
	_ = "STUB: not implemented"

	// binding all data on key is empty.
	return nil
}

// fix: if c.data is nil, don't need to apply map structure

// init default value by tag: default

// add ParseTime support on parse default value

// binding sub-data of the config

// map structure from data

// set result struct ptr

// init default value by tag: default

// ToJSON string, will ignore error
func (c *Config) ToJSON() string { _ = "STUB: not implemented"; return "" }

// WriteTo a writer
func WriteTo(out io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// WriteTo Write out config data representing the current state to a writer.
		nil
}

func (c *Config) WriteTo(out io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DumpTo a writer and use format
func DumpTo(out io.Writer, format string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// DumpTo use the format(json,yaml,toml) dump config data to a writer
func (c *Config) DumpTo(out io.Writer, format string) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// is empty

// encode data to string

// write content to out

// DumpToFile use the format(json,yaml,toml) dump config data to a writer
func (c *Config) DumpToFile(fileName string, format string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
