package rfip

import (
	"encoding/json"
	"reflect"
)

type Resolver struct {
	reflect.Value
}

func NewResolver(master interface{}) Resolver {
	return Resolver{reflect.ValueOf(master)}
}

func (resr *Resolver) Resolve(req *Request) (resp *Response, err error) {
	function, err := resr.funcByName(req.FuncName)
	if err != nil {
		return
	}
	returnValues, err := call(function, req.Argument)
	if err != nil {
		return
	}
	return respond(returnValues)
}

func (resr *Resolver) funcByName(name string) (f reflect.Value, err error) {
	f = resr.MethodByName(name)
	if (f == reflect.Value{}) {
		err = ErrFuncNotFound
	}
	return
}

func call(f reflect.Value, arg []byte) (rvs []reflect.Value, err error) {
	switch argc := f.Type().NumIn(); argc {
	case 0:
		return f.Call([]reflect.Value{}), nil

	case 1:
		return callWithArg(f, arg)

	default:
		return nil, ErrFuncWithBadArgc
	}
}

func callWithArg(f reflect.Value, arg []byte) (rvs []reflect.Value, err error) {
	ptr, err := unmarshalArg(f.Type().In(0), arg)
	if err != nil {
		return
	}
	return f.Call([]reflect.Value{reflect.Indirect(ptr)}), nil
}

func unmarshalArg(t reflect.Type, arg []byte) (ptr reflect.Value, err error) {
	ptr = reflect.New(t)
	if err = json.Unmarshal(arg, ptr.Interface()); err != nil {
		err = ErrBadArgUnmarshal
	}
	return
}

func respond(returnValues []reflect.Value) (resp *Response, err error) {
	switch len(returnValues) {
	case 0:
		return ResponseJSON(nil)

	case 1:
		return ResponseJSON(returnValues[0].Interface())

	default:
		return nil, ErrFuncWithBadRetc
	}
}