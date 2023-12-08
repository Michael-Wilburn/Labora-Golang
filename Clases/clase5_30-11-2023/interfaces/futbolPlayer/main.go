package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall() int
	Name() string
}

type FootbalPlayer struct {
	stamina int
	power   int
}

func (f FootbalPlayer) KickBall() int {
	return f.stamina + f.power

}

func (f FootbalPlayer) Name() string {
	return "Random"
}

type CR7 struct {
	name    string
	stamina int
	power   int
	SUI     int
}

func (f CR7) KickBall() int {
	return f.stamina + f.power*f.SUI

}

func (f CR7) Name() string {
	return "CR7"
}

type Messi struct {
	stamina int
	power   int
	SUI     int
}

func (f Messi) KickBall() int {
	return f.stamina + f.power*f.SUI
}

func (f Messi) Name() string {
	return "Messi"
}

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-2; i++ {
		team[i] = FootbalPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[len(team)-1] = CR7{
		stamina: 10,
		power:   10,
		SUI:     10,
	}
	team[len(team)-2] = Messi{
		stamina: 10,
		power:   10,
		SUI:     10,
	}

	for i := 0; i < len(team); i++ {
		team[i].KickBall()
		fmt.Printf("%s is kicking the ball %d\n", team[i].Name(), team[i].KickBall())
	}
}
