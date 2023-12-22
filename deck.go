package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func (d deck) print(){
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs" }

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck){
 return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
   return strings.Join([]string(d), ",");
}

func (d deck) saveToFile(filename string) error {

	return  os.WriteFile(filename, []byte(d.toString()), 0666)
	
}

func newDeckFromFile(filename string) deck{
   bs, err :=	os.ReadFile(filename)

   if err != nil {
	 fmt.Println("Error: ", err)
	 os.Exit(1)
   }
   return deck(strings.Split(string(bs), ","))
}