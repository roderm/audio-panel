package device

import (
	"net/http"
	"time"
	"context"
)

type nanoleafDriver struct {
	cl *http.Client
}
func NewNanoleafDriver(ctx context.Context, config DeviceConfig) *nanoleafDriver {
	return &nanoleafDriver{
		cl: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}