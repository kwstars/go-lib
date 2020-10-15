package main

import (
	"errors"
	"time"
)

type Player struct {
	Name string
}

func NewPlayer(name string) (Player, error) {
	if time.Now().Unix()%2 == 0 {
		return Player{}, errors.New("player dead")
	}
	return Player{Name: name}, nil
}
