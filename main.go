// main shows a lager logging example
package main

type objA struct {
	b *objB
}

func newA(b *objB) *objA {
	return &objA{
		b: b,
	}
}

func (a *objA) doItA() {
	a.b.doItB()
}

type objB struct {
	c *objC
}

func newB(c *objC) *objB {
	return &objB{
		c: c,
	}
}

func (b *objB) doItB() {
	b.c.doItC()
}

type objC struct{}

func newC() *objC {
	return &objC{}
}

func (c *objC) doItC() {

}

func main() {
	c := newC()
	b := newB(c)
	a := newA(b)

	a.doItA()
}
