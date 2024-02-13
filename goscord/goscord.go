package goscord

import (
	"github.com/tomoconnor/goscord/goscord/gateway"
)

func New(options *gateway.Options) *gateway.Session {
	return gateway.NewSession(options)
}
