package main

type monsterParam string

type Monster struct {
	Name string
}

func NewMonster(name monsterParam) Monster {
	return Monster{Name: string(name)}
}
