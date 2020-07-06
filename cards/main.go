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

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

// Function declaration in go
func newCard() string {
	return "Five of diamonds"
}
