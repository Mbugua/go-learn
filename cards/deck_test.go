package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) !=36 {
		t.Errorf("Expected deck length of 20, but got %v",len(d))
	}

	if d[0] !="Ace of Spades"{
		t.Errorf("Expected first card of Ace of Spades, but got %v",d[0])
	}

	if d[len(d)-1] != "Queen of Clubs"{
		t.Errorf("Expected last card of Queen of Clubs, but got %v", d[len(d)-1])
	}
}