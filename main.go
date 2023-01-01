package main

import (
	abstractfactory "github.com/design-patterns/creational/abstractfactory"
	"github.com/design-patterns/creational/singleton"
	"github.com/design-patterns/structural/adapter"
)

func main() {
	adapter.RunDemo()
	abstractfactory.Demo()
	singleton.Demo()
}
