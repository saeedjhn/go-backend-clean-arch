package main

import "fmt"

// DialogueReciter known how to recite a dialogue
type DialogueReciter interface {
	// Concrete types should implement this method.
	Recite()
}

func main() {
	newToy := NewToy(SpiderMan{})
	// This performs the spider man dialogue.
	newToy.PerformDialogue()
	// Change the behaviour at runtime.
	newToy.SetSuperHero(SuperMan{})
	// This performs the super man dialogue.
	newToy.PerformDialogue()
}

type toy struct {
	DialogueReciter DialogueReciter
}

func NewToy(dr DialogueReciter) *toy {
	return &toy{
		DialogueReciter: dr,
	}
}

func (t *toy) PerformDialogue() {
	t.DialogueReciter.Recite()
}

func (t *toy) SetSuperHero(dr DialogueReciter) {
	t.DialogueReciter = dr
}

type SpiderMan struct{}

func (spm SpiderMan) Recite() {
	fmt.Println("No Man Can Win Every Battle, " +
		"But No Man Should Fall Without A Struggle")
}

type SuperMan struct{}

func (sum SuperMan) Recite() {
	fmt.Println("There is a superhero in all of us, " +
		"we just need the courage to put on the cape")
}

type BatMan struct{}

func (sum BatMan) Recite() {
	fmt.Println("It's not who I am underneath, " +
		"but what I do that defines me!")
}
