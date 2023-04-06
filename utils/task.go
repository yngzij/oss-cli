package utils

import (
	"fmt"
	"reflect"
)

type FuncHandler struct {
	f    reflect.Value
	args []reflect.Value
}

var Funcs []*FuncHandler

func NewFuncHandler(f interface{}, args ...interface{}) *FuncHandler {
	fv := reflect.ValueOf(f)

	if fv.Kind() != reflect.Func {
		panic("event handler must be a func.")
	}

	ft := fv.Type()

	if len(args) > 0 {
		if len(args) != ft.NumIn() {
			panic("event handler args number not match.")
		}
		fargs := make([]reflect.Value, len(args))
		for i, arg := range args {
			fargs[i] = reflect.ValueOf(arg)
		}
		return &FuncHandler{f: fv, args: fargs}
	} else {
		return &FuncHandler{f: fv}
	}
}

func (h *FuncHandler) Call() (ret []reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("event call error: %s", r)
			}
		}
	}()

	if len(h.args) > 0 {
		ret = h.f.Call(h.args)
	} else {
		ret = h.f.Call(nil)
	}

	return
}
