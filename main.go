// main shows a lager logging example
package main

import (
	"os"

	"code.cloudfoundry.org/lager"
)

var logger lager.Logger

type objA struct {
	b *objB
}

func newA(b *objB) *objA {
	return &objA{
		b: b,
	}
}

func (a *objA) doItA() {
	logger.Debug("obj-a.do-it-a")
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
	logger.Debug("obj-b.do-it-b")
	b.c.doItC()
}

type objC struct{}

func newC() *objC {
	return &objC{}
}

func (c *objC) doItC() {
	logger.Debug("obj-c.do-it-c")
}

func main() {
	logger = lager.NewLogger("logging-example")
	logger.RegisterSink(lager.NewPrettySink(os.Stderr, lager.DEBUG))

	c := newC()
	b := newB(c)
	a := newA(b)

	a.doItA()
}
