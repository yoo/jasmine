package jasmine

import "github.com/gopherjs/gopherjs/js"

// Describe discribes a test suit containing multiple specs.
// The specs are run inside the callback function fn.
func Describe(name string, fn func()) {
	js.Global.Call("describe", name, fn)
}

// XDescribe is a canncled or commented out suit which will not be run.
func XDescribe(name string, fn func()) {
	js.Global.Call("xdescribe", name, fn)
}

// It is a spec to be tested with Expectations.
func It(behavior string, fn func()) {
	js.Global.Call("it", behavior, fn)
}

// XIt is a cancled or commented out spec which will not be run.
func XIt(behavior string, fn func()) {
	js.Global.Call("xit", behavior, fn)
}

func ItAsync(behavior string, fn func(func())) {

	var cb = js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		return js.Global.Get("eval").Invoke("var a = function(cb){ return function(done){ return cb(done); }; }; a")
	}).Invoke().Invoke(fn)
	js.Global.Call("it", behavior, cb)
}

func XitAsync(behavior string, fn func(func())) {

	js.Global.Call("xit", behavior, createFuncWithVisibleDone(fn))
}

func BeforeEach(fn func()) {
	js.Global.Call("beforeEach", fn)
}

func BeforeEachAsync(fn func(func())) {
	js.Global.Call("beforeEach", createFuncWithVisibleDone(fn))
}

func AfterEach(fn func()) {
	js.Global.Call("afterEach", fn)
}

func AfterEachAsync(fn func(func())) {
	js.Global.Call("afterEach", createFuncWithVisibleDone(fn))
}

func Expect(value interface{}) *Expectation {
	return &Expectation{o: js.Global.Call("expect", value)}
}

func SetDefaultTimeoutInterval(interval int) {
	js.Global.Get("jasmine").Set("DEFAULT_TIMEOUT_INTERVAL", interval)
}

func Fail(message string) {
	js.Global.Call("fail", message)
}

func createFuncWithVisibleDone(fn func(func())) *js.Object {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		return js.Global.Get("eval").Invoke("var a = function(cb){ return function(done){ return cb(done); }; }; a")
	}).Invoke().Invoke(fn)
}
