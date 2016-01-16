package main

import "github.com/arvitaly/gopherjs-jasmine"

func caller(f func()) {
	f()
}
func main() {
	jasmine.RunJasmineTests()
}
