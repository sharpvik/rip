package rip

import (
	"encoding/json"
	"reflect"
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

func (r *resolver) Handle(req *Request) (resp *Response) {
	function, err := r.funcByName(req.FuncName)
	if err != nil {
		return ResponseError(err)
	}
	returnValues, err := call(function, req.Argument)
	if err != nil {
		return ResponseError(err)
	}
	return respond(returnValues)
}

func (r *resolver) funcByName(name string) (f reflect.Value, err Error) {
	f = r.MethodByName(name)
	if (f == reflect.Value{}) {
		err = ErrFuncNotFound
	}
	return
}

func call(f reflect.Value, arg []byte) (rvs []reflect.Value, err Error) {
	switch argc := f.Type().NumIn(); argc {
	case 0:
		return f.Call([]reflect.Value{}), nil

	case 1:
		return callWithArg(f, arg)

	default:
		return nil, ErrFuncWithBadArgc
	}
}

func callWithArg(f reflect.Value, arg []byte) (rvs []reflect.Value, err Error) {
	ptr, err := unmarshalArg(f.Type().In(0), arg)
	if err != nil {
		return
	}
	return f.Call([]reflect.Value{reflect.Indirect(ptr)}), nil
}

func unmarshalArg(t reflect.Type, arg []byte) (ptr reflect.Value, err Error) {
	ptr = reflect.New(t)
	if err := json.Unmarshal(arg, ptr.Interface()); err != nil {
		err = ErrBadArgUnmarshal
	}
	return
}

func respond(returnValues []reflect.Value) (resp *Response) {
	switch len(returnValues) {
	case 0:
		return ResponseJSON(nil)

	case 1:
		return ResponseJSON(returnValues[0].Interface())

	default:
		return ResponseError(ErrFuncWithBadRetc)
	}
}
