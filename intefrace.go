package world_core

import (
	"github.com/open-flow/world-core/exchange"
	"github.com/open-flow/world-core/subscribe"
)

type World interface {
	RegisterExchange(exchange exchange.Exchange) error
	RegisterSubscriber(subscriber subscribe.Subscriber) error
	Connect() error
	Drain() error
}
