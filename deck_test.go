package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 30 {
		t.Failed()
	}

	if d[0] != "One of Spades" {
		t.Failed()
	}

	if d[len(d)-1] != "Ten of Diamonds" {
		t.Failed()
	}
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	newD := newDeckFromFile("_decktesting")

	if len(newD) != 30 {
		t.Failed()
	}
}
