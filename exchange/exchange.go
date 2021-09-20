package exchange

import (
	"github.com/open-flow/world-core/werrors"
	"reflect"
)

type Exchange interface {
	Configuration() Configuration
	Exchange(request Request) Response
}

type Request struct {
	Headers map[string]string
	Payload interface{}
}

type Response struct {
	Error   *werrors.Error
	Headers map[string]string
	Payload interface{}
	Publish []interface{}
}

type Configuration struct {
	Request  reflect.Type
	Response reflect.Type
	Publish  []reflect.Type
}
