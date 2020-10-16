//+build wireinject

package main

import "github.com/google/wire"

func injectedMessage() string {
	wire.Build(provideFoo, wire.FieldsOf(new(Foo), "S"))
	return ""
}
