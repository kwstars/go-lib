//+build wireinject

package main

import "github.com/google/wire"

var kitty = Monster{Name: "big boss"}

func InitEndingA(name string) EndingA {
	wire.Build(NewPlayer, wire.Value(kitty), NewEndingA)
	return EndingA{}
}

func InitEndingB(name string) EndingB {
	wire.Build(NewPlayer, wire.Value(kitty), NewEndingB)
	return EndingB{}
}
