//+build wireinject

package main

import "github.com/google/wire"

func InitMission(p playerParam, m monsterParam) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}
