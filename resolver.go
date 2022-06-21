package rip

import (
	"encoding/json"
	"reflect"

	"github.com/sharpvik/rip/proto"
)

type resolver struct {
	reflect.Value
}

func Use(source interface{}) *resolver {
	return &resolver{reflect.ValueOf(source)}
}

func (r *resolver) Server() *Server {
	return &Server{
		resolver: r,
	}
}

func (r *resolver) Handle(req *proto.Request) (resp *proto.Response) {
	function, err := r.funcByName(req.Function)
	if err != nil {
		return proto.ResponseError(err)
	}
	returnValues, err := invoke(function, req.Argument)
	if err != nil {
		return proto.ResponseError(err)
	}
	return respond(returnValues)
}

func (r *resolver) funcByName(name string) (f reflect.Value, err proto.Error) {
	f = r.MethodByName(name)
	if (f == reflect.Value{}) {
		err = proto.ErrFuncNotFound
	}
	return
}

func invoke(f reflect.Value, arg []byte) (rvs []reflect.Value, err proto.Error) {
	switch argc := f.Type().NumIn(); argc {
	case 0:
		return f.Call([]reflect.Value{}), nil

	case 1:
		return callWithArg(f, arg)

	default:
		return nil, proto.ErrFuncWithBadArgc
	}
}

func callWithArg(f reflect.Value, arg []byte) (rvs []reflect.Value, err proto.Error) {
	ptr, err := unmarshalArg(f.Type().In(0), arg)
	if err != nil {
		return
	}
	return f.Call([]reflect.Value{reflect.Indirect(ptr)}), nil
}

func unmarshalArg(t reflect.Type, arg []byte) (ptr reflect.Value, err proto.Error) {
	ptr = reflect.New(t)
	if err := json.Unmarshal(arg, ptr.Interface()); err != nil {
		err = proto.ErrBadArgUnmarshal
	}
	return
}

func respond(returnValues []reflect.Value) (resp *proto.Response) {
	switch len(returnValues) {
	case 0:
		return proto.ResponseJSON(nil)

	case 1:
		return proto.ResponseJSON(returnValues[0].Interface())

	default:
		return proto.ResponseError(proto.ErrFuncWithBadRetc)
	}
}
