// main shows a lager logging example
package main

import (
	"os"

	"code.cloudfoundry.org/lager"
)

type objA struct {
	b *objB
}

func newA(b *objB) *objA {
	return &objA{
		b: b,
	}
}

func (a *objA) doItA(logger lager.Logger, target string) {
	logger = logger.Session("obj-a", lager.Data{"target": target})
	logger.Debug("do-it-a")
	a.b.doItB(logger)
}

type objB struct {
	c *objC
}

func newB(c *objC) *objB {
	return &objB{
		c: c,
	}
}

func (b *objB) doItB(logger lager.Logger) {
	logger = logger.Session("obj-b")
	logger.Debug("do-it-b")
	b.c.doItC(logger)
	b.c.doItC(logger)
}

type objC struct{}

func newC() *objC {
	return &objC{}
}

func (c *objC) doItC(logger lager.Logger) {
	logger = logger.Session("obj-c")
	logger.Debug("do-it-c")
}

func main() {
	logger := lager.NewLogger("logging-example")
	logger.RegisterSink(lager.NewPrettySink(os.Stderr, lager.DEBUG))

	c := newC()
	b := newB(c)
	a := newA(b)

	a.doItA(logger, "foo")
	a.doItA(logger, "bar")
}
