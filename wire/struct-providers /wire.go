//+build wireinject

package main

import "github.com/google/wire"

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "Monster"))
var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "Player"))

//*表示注入所有字段
//var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "*"))
//var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "*"))

func InitEndingA(name string) EndingA {
	wire.Build(endingASet)
	return EndingA{}
}

func InitEndingB(name string) EndingB {
	wire.Build(endingBSet)
	return EndingB{}
}
