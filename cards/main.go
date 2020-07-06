// package main
// import "fmt"
// func main() {
// 	// var card string = "Ace of Spades"
// 	card := "Ace of Spades"
// 	fmt.Println(card)
// 	card = "Five of Diamonds"
// 	fmt.Println(card)
// }

package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}

// Function declaration in go
func newCard() string {
	return "Five of diamonds"
}
