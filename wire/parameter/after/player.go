package main

type playerParam string

type Player struct {
	Name string
}

func NewPlayer(name playerParam) Player {
	return Player{Name: string(name)}
}
