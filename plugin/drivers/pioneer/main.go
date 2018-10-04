package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	pl "github.com/roderm/audio-panel/plugin/iface"
	"io/ioutil"
	"log"
	"os"
)

func NewDriver(ctx context.Context, config interface{}, lgr *log.Logger, identifier string) (pl.IDevice, error) {
	var conf PioneerConfig
	err := mapstructure.Decode(config, &conf)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(conf.Setup); err != nil {
		return nil, fmt.Errorf("path to commandset(%s) not found.", conf.Setup)
	}
	fileContent, err := ioutil.ReadFile(conf.Setup)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read configfile(%s).", conf.Setup)
	}
	var setup CommandSet
	err = json.Unmarshal(fileContent, &setup)
	if err != nil {
		return nil, fmt.Errorf("Bad Syntax in configfile(%s), wasn't able to parse json.", conf.Setup)
	}
	return NewPioneerDriver(ctx, conf, setup, lgr, identifier)
}
