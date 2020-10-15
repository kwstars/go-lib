// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitEndingA(name string) EndingA {
	player := NewPlayer(name)
	monster := _wireMonsterValue
	endingA := NewEndingA(player, monster)
	return endingA
}

var (
	_wireMonsterValue = kitty
)

func InitEndingB(name string) EndingB {
	player := NewPlayer(name)
	monster := _wireMainMonsterValue
	endingB := NewEndingB(player, monster)
	return endingB
}

var (
	_wireMainMonsterValue = kitty
)

// wire.go:

var kitty = Monster{Name: "big boss"}
