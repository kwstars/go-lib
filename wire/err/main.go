package main

import "log"

func main() {
	m, err := InitMission("test")
	if err != nil {
		log.Fatal(err)
	}
	m.Start()
}
