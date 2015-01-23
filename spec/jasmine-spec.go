package main

import "github.com/JohannWeging/jasmine"

func caller(f func()) {
	f()
}
func main() {
	jasmine.RunJasmineTests()
}
