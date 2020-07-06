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
	res := cards.saveToFile("my_cards")
	fmt.Println(res)
}

// Function declaration in go
func newCard() string {
	return "Five of diamonds"
}
