package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/gookit/goutil"
	"github.com/gookit/goutil/cliutil"
)

func main() {
	config.AddDriver(yaml.Driver)
	config.WithOptions(
		config.ParseEnv,
		config.WithHookFunc(func(event string, c *config.Config) {
			if event == config.OnReloadData {
				cliutil.Cyanln("config reloaded, you can do something ....")
			}
		}),
	)

	// load app config files
	err := config.LoadFiles(
		"testdata/json_base.json",
		"testdata/yml_base.yml",
		"testdata/yml_other.yml",
	)
	if err != nil {
		panic(err)
	}

	// mock server running
	done := make(chan bool)

	// watch loaded config files
	err = watchConfigFiles(config.Default())
	goutil.PanicErr(err)

	cliutil.Infoln("loaded config files is watching ...")
	<-done
}

func watchConfigFiles(cfg *config.Config) error { _ = "STUB: not implemented"; return nil }

//noinspection GoUnhandledErrorResult

// get loaded files

// 'Events' channel is closed

// if event.Op > 0 {

// }

// 'Errors' channel is not closed
