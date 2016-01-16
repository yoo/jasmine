package jasmine

import (
	"time"
)

// +testing
func caller(f func()) {
	f()
}
func RunJasmineTests() {
	Describe("testig basic jasmine functions", func() {
		Describe("nesed siute", func() {
			XIt("schould not be called", func() {
				Expect(1).ToBe(2)
			})
		})

		XDescribe("this suite should not be called", func() {
			It("should not be called", func() {
				Expect(1).ToBe(2)
			})
		})
	})

	Describe("testing basic expectations", func() {
		It("test spec ToBe", func() {
			Expect(true).ToBe(true)
			Expect(true).Not.ToBe(false)
		})

		It("test spec ToEqual", func() {
			Expect(1).ToEqual(1)
			Expect(1).Not.ToEqual(2)
		})

		It("test spec ToMatch", func() {
			Expect("foo").ToMatch("foo")
			Expect("foo").Not.ToMatch("bar")
		})

		It("test spec ToBeDefined", func() {
			Expect(1).ToBeDefined()
		})

		It("test spec ToBeNull", func() {
			Expect(nil).ToBeNull()
		})

		It("test spec ToBeTruthy", func() {
			Expect(true).ToBeTruthy()
			Expect(false).Not.ToBeTruthy()
		})

		It("test spec ToBeFalsy", func() {
			Expect(false).ToBeFalsy()
			Expect(true).Not.ToBeFalsy()
		})

		It("test spec ToContain", func() {
			array := []int{1, 2, 3}
			exp := Expect(array)
			exp.ToContain(2)
			exp.Not.ToContain(5)
		})

		It("test spec ToBeLessThan", func() {
			Expect(1).ToBeLessThan(4)
		})

		It("test spec ToBeGreaterThan", func() {
			Expect(4).ToBeGreaterThan(1)
		})
	})

	Describe("testing setup and teardown methods", func() {
		var a = 0

		BeforeEach(func() {
			a = 1
		})

		AfterEach(func() {
			a = 2
		})

		It("test setup and teardown", func() {
			Expect(a).ToBe(1)
		})

	})

	Describe("testing spy like functinallity with ItAsync", func() {
		var ch chan bool

		BeforeEach(func() {
			ch = make(chan bool)
		})
		ItAsync("test toHaveBeenCalled", func(done func()) {
			caller(func() {
				go func() {
					ch <- true
				}()
			})
			go func() {
				b := <-ch
				Expect(b).ToBeTruthy()
				done()
			}()
		})
	})
	Describe("testing async functions", func() {
		SetDefaultTimeoutInterval(10)
		BeforeEachAsync(func(done func()) {
			time.AfterFunc(time.Millisecond*3, func() {
				done()
			})
		})
		ItAsync("ItAsync with BeforeEachAsync and AfterEachAsync", func(done func()) {
			time.AfterFunc(time.Millisecond*3, func() {
				done()
			})
		})
		AfterEachAsync(func(done func()) {
			time.AfterFunc(time.Millisecond*3, func() {
				done()
			})
		})
	})
}
