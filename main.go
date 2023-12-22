package main

import "fmt"



func main(){
	cards :=  newDeck()

	cards = append(cards, "Six of Spades")

	hand, _ := deal(cards, 5)
	

	hand.saveToFile("hand.txt")

	fmt.Println(newDeckFromFile("hand.txt").toString())

}

