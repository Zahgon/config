package config

// LoadFiles load one or multi files, will fire OnLoadData event
//
// Usage:
//
//	config.LoadFiles(file1, file2, ...)
func LoadFiles(sourceFiles ...string) error { _ = "STUB: not implemented"; return nil }

// LoadFiles load and parse config files, will fire OnLoadData event
func (c *Config) LoadFiles(sourceFiles ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadExists load one or multi files, will ignore not exist
//
// Usage:
//
//	config.LoadExists(file1, file2, ...)
func LoadExists(sourceFiles ...string) error { _ = "STUB: not implemented"; return nil }

// LoadExists load and parse config files, but will ignore not exists file.
func (c *Config) LoadExists(sourceFiles ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FileFilterFn for check file is should load.
type FileFilterFn func(file string, c *Config) (shouldLoad bool, format string)

// LoadFilesByFilter load one or multi files by give filter checked, will fire OnLoadData event
func LoadFilesByFilter(configFiles []string, filter FileFilterFn) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadFilesByFilter load one or multi files by give filter checked, will fire OnLoadData event
//   - `filter` return format can be emtpy, will auto detect it by file extension
func (c *Config) LoadFilesByFilter(configFiles []string, filter FileFilterFn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadRemote load config data from remote URL.
func LoadRemote(format, url string) error { _ = "STUB: not implemented"; return nil }

// LoadRemote load config data from remote URL.
//
// Usage:
//
//	c.LoadRemote(config.JSON, "http://abc.com/api-config.json")
func (c *Config) LoadRemote(format, url string) (err error) {
	_ = "STUB: not implemented"
	// create http client
	return nil
}

//noinspection GoUnhandledErrorResult

// read response content

// LoadOSEnv load data from OS ENV
//
// Deprecated: please use LoadOSEnvs()
func LoadOSEnv(keys []string, keyToLower bool) { _ = "STUB: not implemented"; return }

// LoadOSEnv load data from os ENV
//
// Deprecated: please use Config.LoadOSEnvs()
func (c *Config) LoadOSEnv(keys []string, keyToLower bool) { _ = "STUB: not implemented"; return }

// NOTICE: if is Windows os, os.Getenv() Key is not case-sensitive

// LoadOSEnvs load data from OS ENVs. see Config.LoadOSEnvs
func LoadOSEnvs(nameToKeyMap map[string]string) { _ = "STUB: not implemented"; return }

// LoadOSEnvs load data from os ENVs. format: `{ENV_NAME: config_key}`
//
//   - `config_key` allow use key path. eg: `{"DB_USERNAME": "db.username"}`
func (c *Config) LoadOSEnvs(nameToKeyMap map[string]string) { _ = "STUB: not implemented"; return }

// LoadOSEnvByFilter load OS ENVs by custom fitler func. eg: use for load ENV by prefix.
func LoadOSEnvByFilter(filterFn func(key string) (loadIt bool, cfgKey string)) {
	_ = "STUB: not implemented"
	return
}

// LoadOSEnvByFilter load OS ENVs by custom fitler func. eg: use for load ENV by prefix.
//
//   - `filterFn` return cfgKey can be empty, will use key instead.
func (c *Config) LoadOSEnvByFilter(filterFn func(key string) (loadIt bool, cfgKey string)) {
	_ = "STUB: not implemented"
	return
}

// support bound types for CLI flags vars
var validTypes = map[string]int{
	"int":  1,
	"uint": 1,
	"bool": 1,
	// string is default
	"string": 1,
}

// LoadFlags load data from cli flags. see Config.LoadFlags
func LoadFlags(defines []string) error { _ = "STUB: not implemented"; return nil }

// LoadFlags parse command line arguments, based on provide keys.
//
// Usage:
//
//	// 'debug' flag is bool type
//	c.LoadFlags([]string{"env", "debug:bool"})
//	// can with flag desc message
//	c.LoadFlags([]string{"env:set the run env"})
//	c.LoadFlags([]string{"debug:bool:set debug mode"})
//	// can set value to map key. eg: myapp --map1.sub-key=val
//	c.LoadFlags([]string{"--map1.sub-key"})
func (c *Config) LoadFlags(defines []string) (err error) { _ = "STUB: not implemented"; return nil }

// bind vars

// as string

// parse and collect

// only get name in the keys.

// if f.Value implement the flag.Getter, read typed value

// } else { // TIP: basic type flag always implements Getter interface
// 	_ = c.Set(name, f.Value.String()) // ignore error

// LoadData load one or multi data
func LoadData(dataSource ...any) error { _ = "STUB: not implemented"; return nil }

// LoadData load data from map OR struct
//
// The dataSources type allow:
//   - map[string]any
//   - map[string]string
func (c *Config) LoadData(dataSources ...any) (err error) { _ = "STUB: not implemented"; return nil }

// LoadSMap to config
func (c *Config) LoadSMap(smp map[string]string) { _ = "STUB: not implemented"; return }

// LoadSources load one or multi byte data
func LoadSources(format string, src []byte, more ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadSources load data from byte content.
//
// Usage:
//
//	config.LoadSources(config.Yaml, []byte(`
//	  name: blog
//	  arr:
//		key: val
//
// `))
func (c *Config) LoadSources(format string, src []byte, more ...[]byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadStrings load one or multi string
func LoadStrings(format string, str string, more ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadStrings load data from source string content.
func (c *Config) LoadStrings(format string, str string, more ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadFilesByFormat load one or multi config files by give format, will fire OnLoadData event
func LoadFilesByFormat(format string, configFiles ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadFilesByFormat load one or multi files by give format, will fire OnLoadData event
func (c *Config) LoadFilesByFormat(format string, configFiles ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadExistsByFormat load one or multi files by give format, will fire OnLoadData event
func LoadExistsByFormat(format string, configFiles ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadExistsByFormat load one or multi files by give format, will fire OnLoadData event
func (c *Config) LoadExistsByFormat(format string, configFiles ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadOptions for load config from dir.
type LoadOptions struct {
	// DataKey use for load config from dir.
	// see https://github.com/gookit/config/issues/173
	DataKey string
}

// LoadOptFn type func
type LoadOptFn func(lo *LoadOptions)

func newLoadOptions(loFns []LoadOptFn) *LoadOptions { _ = "STUB: not implemented"; return nil }

// LoadFromDir Load custom format files from the given directory, the file name will be used as the key.
//
// Example:
//
//	// file: /somedir/task.json
//	LoadFromDir("/somedir", "json")
//
//	// after load
//	Config.data = map[string]any{"task": file data}
func LoadFromDir(dirPath, format string, loFns ...LoadOptFn) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadFromDir Load custom format files from the given directory, the file name will be used as the key.
//
// NOTE: will not be reloaded on call ReloadFiles(), if data loaded by the method.
//
// Example:
//
//	// file: /somedir/task.json , will use filename 'task' as key
//	Config.LoadFromDir("/somedir", "json")
//
//	// after load, the data will be:
//	Config.data = map[string]any{"task": {file data}}
func (c *Config) LoadFromDir(dirPath, format string, loFns ...LoadOptFn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// filename without ext.

// TODO use file name as key, it cannot be reloaded. So, cannot append to loadedFiles
// c.loadedFiles = append(c.loadedFiles, fPath)

// ReloadFiles reload config data use loaded files
func ReloadFiles() error { _ = "STUB: not implemented"; return nil }

// ReloadFiles reload config data use loaded files. use on watching loaded files change
func (c *Config) ReloadFiles() (err error) { _ = "STUB: not implemented"; return nil }

// revert to back up data on error

// with lock

// reload config files

// load config file, will fire OnLoadData event
//   - loadExist=false will return error on file not exists
func (c *Config) loadFile(file string, loadExist bool, format string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// skip not exist file

//noinspection GoUnhandledErrorResult

// read file content

// get format for file ext

// parse file content

// parse config source code to Config.
func (c *Config) parseSourceCode(format string, blob []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) loadDataMap(data map[string]any) (err error) {
	_ = "STUB: not implemented"
	// first: init config data
	return nil
}

// again ... will merge data

// parse config source code to Config.
func (c *Config) parseSourceToMap(format string, blob []byte) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decode content to data
