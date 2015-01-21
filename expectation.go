package jasmine

import "github.com/gopherjs/gopherjs/js"

type Expectation struct {
	o   js.Object
	Not *Expectation `js:"not"`
}

func (e *Expectation) ToBe(value interface{}) {
	e.o.Call("toBe", value)
}

func (e *Expectation) ToEqual(value interface{}) {
	e.o.Call("toEqual", value)
}

// ToMatch is not able to match against a regex at the moment
func (e *Expectation) ToMatch(value interface{}) {
	e.o.Call("toMatch", value)
}

// ToBeDefined checks if the expectation is defined
func (e *Expectation) ToBeDefined() {
	e.o.Call("toBeDefined")
}

// ToBeUndefined checks if the expectation is undefined
func (e *Expectation) ToBeUndefined() {
	e.o.Call("toBeUndefined")
}

// ToBeNull checks if the expectation to be null / nil
func (e *Expectation) ToBeNull() {
	e.o.Call("toBeNull")
}

func (e *Expectation) ToBeTruthy() {
	e.o.Call("toBeTruthy")
}

func (e *Expectation) ToBeFalsy() {
	e.o.Call("toBeFalsy")
}

func (e *Expectation) ToContain(value interface{}) {
	e.o.Call("toContain", value)
}

func (e *Expectation) ToBeLessThan(value interface{}) {
	e.o.Call("toBeLessThan", value)
}

func (e *Expectation) ToBeGreaterThan(value interface{}) {
	e.o.Call("toBeGreaterThan", value)
}

func (e *Expectation) ToBeCloseTo(value interface{}, percision int) {
	e.o.Call("toBeCloseTo", value, percision)
}

func (e *Expectation) ToThrow() {
	e.o.Call("toThrow")
}

func (e *Expectation) ToThrowError(value interface{}) {
	e.o.Call("toThrowError", value)
}
func (e *Expectation) ToHaveBeenCalled() {
	e.o.Call("toHaveBeenCalled")
}

func (e *Expectation) ToHaveBeenCalledWith(values ...interface{}) {
	e.o.Call("toHaveBeenCalledWith", values...)
}

func (e *Expectation) TallThrough() {
	e.o.Call("callThrough")
}
