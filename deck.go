package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}