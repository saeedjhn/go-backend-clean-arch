package main

import "log"

type DialogueReciter interface {
	Recite()
}

type Toy struct {
	DialogueReciter DialogueReciter
}

func NewToy(dr DialogueReciter) *Toy {
	return &Toy{
		DialogueReciter: dr,
	}
}

func (t *Toy) PerformDialogue() {
	t.DialogueReciter.Recite()
}

func (t *Toy) SetSuperHero(dr DialogueReciter) {
	t.DialogueReciter = dr
}

type SpiderMan struct{}

func (spm SpiderMan) Recite() {
	log.Println("No Man Can Win Every Battle, " +
		"But No Man Should Fall Without A Struggle")
}

type SuperMan struct{}

func (sum SuperMan) Recite() {
	log.Println("There is a superhero in all of us, " +
		"we just need the courage to put on the cape")
}

type BatMan struct{}

func (sum BatMan) Recite() {
	log.Println("It's not who I am underneath, " +
		"but what I do that defines me!")
}

// func main() {
// 	newToy := NewToy(SpiderMan{})
// 	newToy.PerformDialogue()
// 	newToy.SetSuperHero(SuperMan{})
// 	newToy.PerformDialogue()
// 	newToy.SetSuperHero(BatMan{})
// 	newToy.PerformDialogue()
// }
