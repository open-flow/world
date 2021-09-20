package subscribe

import (
	"github.com/open-flow/world-core/werrors"
	"reflect"
)

type Subscriber interface {
	Configuration() Configuration
	Notify(notification Notification) Effect
}

type Notification struct {
	Headers map[string]string
	Model   interface{}
}

type Effect struct {
	Error *werrors.Error
	Model interface{}
}

type Configuration struct {
	Subscribe []reflect.Type
	Publish   []reflect.Type
}
