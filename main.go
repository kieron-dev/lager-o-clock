// main shows a lager logging example
package main

import (
	"os"

	"code.cloudfoundry.org/lager"
)

type objA struct {
	b      *objB
	logger lager.Logger
}

func newA(b *objB, logger lager.Logger) *objA {
	return &objA{
		b:      b,
		logger: logger.Session("obj-a"),
	}
}

func (a *objA) doItA(target string) {
	a.logger.Debug("do-it-a", lager.Data{"target": target})
	a.b.doItB()
	a.b.doAnotherB()
}

type objB struct {
	c      *objC
	logger lager.Logger
}

func newB(c *objC, logger lager.Logger) *objB {
	return &objB{
		c:      c,
		logger: logger.Session("obj-b"),
	}
}

func (b *objB) doItB() {
	b.logger.Debug("do-it-b")
	b.c.doItC()
	b.c.doItC()
}

func (b *objB) doAnotherB() {
	b.logger.Debug("do-another-b")
}

type objC struct {
	logger lager.Logger
}

func newC(logger lager.Logger) *objC {
	return &objC{
		logger: logger.Session("obj-c"),
	}
}

func (c *objC) doItC() {
	c.logger.Debug("do-it-c")
}

func main() {
	logger := lager.NewLogger("logging-example")
	logger.RegisterSink(lager.NewPrettySink(os.Stderr, lager.DEBUG))

	c := newC(logger)
	b := newB(c, logger)
	a := newA(b, logger)

	a.doItA("foo")
	a.doItA("bar")
}
