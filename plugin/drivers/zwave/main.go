package main

import (
	"context"
	"log"

	pl "github.com/roderm/audio-panel/plugin/iface"
)

func NewDriver(ctx context.Context, config interface{}, lgr *log.Logger, identifier string) (pl.IDevice, error) {
}
