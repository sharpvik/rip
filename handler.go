package rip

import (
	"encoding/json"
	"reflect"

	"github.com/sharpvik/rip/proto"
)

type Handler interface {
	Handle(*proto.Request) *proto.Response
}

type handler struct {
	reflect.Value
}

func NewHanlder(source interface{}) Handler {
	return &handler{Value: reflect.ValueOf(source)}
}

func (h *handler) Handle(req *proto.Request) *proto.Response {
	function, e := h.funcByName(req.Function)
	if e != nil {
		return proto.ResponseError(e)
	}
	returnValues, e := invoke(function, req.Argument)
	if e != nil {
		return proto.ResponseError(e)
	}
	return respond(returnValues)
}

func (h *handler) funcByName(name string) (f reflect.Value, e proto.Error) {
	f = h.MethodByName(name)
	if (f == reflect.Value{}) {
		e = proto.ErrFuncNotFound
	}
	return
}

func invoke(f reflect.Value, arg []byte) ([]reflect.Value, proto.Error) {
	switch argc := f.Type().NumIn(); argc {
	case 0:
		return f.Call([]reflect.Value{}), nil

	case 1:
		return callWithArg(f, arg)

	default:
		return nil, proto.ErrFuncWithBadArgc
	}
}

func callWithArg(f reflect.Value, arg []byte) (rvs []reflect.Value, e proto.Error) {
	ptr, e := unmarshalArg(f.Type().In(0), arg)
	if e != nil {
		return
	}
	return f.Call([]reflect.Value{reflect.Indirect(ptr)}), nil
}

func unmarshalArg(t reflect.Type, arg []byte) (ptr reflect.Value, e proto.Error) {
	ptr = reflect.New(t)
	if err := json.Unmarshal(arg, ptr.Interface()); err != nil {
		e = proto.ErrBadArgUnmarshal
	}
	return
}

func respond(returnValues []reflect.Value) *proto.Response {
	switch len(returnValues) {
	case 0:
		return proto.ResponseJSON(nil)

	case 1:
		return proto.ResponseJSON(returnValues[0].Interface())

	default:
		return proto.ResponseError(proto.ErrFuncWithBadRetc)
	}
}
