package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Contender interface {
	GetName() string
	ThrowAttack() int
	RecieveAttack(intensity int)
}

func main() {
	var criminal Criminal = newCriminal()
	var police Police = newPolice()

	var contenders [2]Contender

	areBothAlive := criminal.IsAlive() && police.IsAlive()
	oneOrZero := rand.Intn(2)
	contenders[oneOrZero] = &police
	contenders[(oneOrZero+1)%2] = &criminal
	var turns float64 = 0
	var currentAttacker int = 0
	for areBothAlive {
		attackIntensity := contenders[currentAttacker].ThrowAttack()
		fmt.Println(contenders[currentAttacker].GetName(), "ataca con intensidad", attackIntensity)
		contenders[(currentAttacker+1)%2].RecieveAttack(attackIntensity)
		currentAttacker = (currentAttacker + 1) % 2

		areBothAlive = criminal.IsAlive() && police.IsAlive()
		fmt.Printf("Police(Life=%d,Armour=%d), CriminalLife=%d, turno número %d\n", police.Life, police.Armour, criminal.Life, int(turns))
		turns += 0.5

		time.Sleep(2 * time.Second)
	}

	if police.IsAlive() {
		fmt.Println("Gano policia")
	} else {
		fmt.Println("Ganó ladron")
	}
}

type Criminal struct {
	Fighter // PROMOTED MEMBERS!!
}

func newCriminal() Criminal {
	return Criminal{
		Fighter: Fighter{
			Life: 100,
		},
	}
}

func (c Criminal) ThrowAttack() int {
	return rand.Intn(60)
}

func (c *Criminal) RecieveAttack(intensity int) {
	c.Life -= intensity
}

func (c Criminal) GetName() string {
	return "Criminal"
}

type Police struct {
	Fighter
	Armour int // 0..50
}

func newPolice() Police {
	return Police{
		Fighter: Fighter{
			Life: 100,
		},
		Armour: 50,
	}
}

func (p Police) ThrowAttack() int {
	return rand.Intn(40)
}

func (p *Police) RecieveAttack(intensity int) {
	if p.Armour > 0 {
		diff := (p.Armour - intensity)
		hasArmour := diff > 0
		if hasArmour {
			p.Armour -= intensity
			intensity = 0
		} else {
			p.Armour = 0
			intensity = -diff // intensity -= p.Armour
		}
	}
	p.Life -= intensity
}

func (p Police) GetName() string {
	return "Policia"
}

type Fighter struct {
	Life int
}

func (f Fighter) IsAlive() bool {
	return f.Life > 0
}

func (f Fighter) GetLife() int {
	return f.Life
}
