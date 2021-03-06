// Package http provides a http based message broker
package http

import (
	"github.com/mjlee1983/go-micro/go-micro/v2/broker"
)

// NewBroker returns a new http broker
func NewBroker(opts ...broker.Option) broker.Broker {
	return broker.NewBroker(opts...)
}
