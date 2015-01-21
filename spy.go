package jasmine

import "github.com/gopherjs/gopherjs/js"

type Spy struct {
	o js.Object
}

func (s *Spy) ReturnValue(value interface{}) {
	s.o.Call("returnValue", value)
}

func (s *Spy) CallFake(fn func()) {
	s.o.Call("callFake", fn)
}

func (s *Spy) ThrowError(err string) *Spy {
	return &Spy{o: s.o.Call("throwError", err)}
}
