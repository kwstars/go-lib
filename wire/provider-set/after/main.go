package main

func main() {
	m := InitMission("test")
	m.Start()

	a := InitEndingA("aaa")
	a.Appear()
	b := InitEndingB("bbb")
	b.Appear()
}
