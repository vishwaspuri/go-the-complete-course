package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// If no error
	sOS := strings.Split(string(bs), ",") // Slice of string
	return deck(sOS)
}

func (d deck) shuffle() {
	// Truly random random number generator
	// Creating a new source object
	source := rand.NewSource(time.Now().UnixNano()) // Using time to get a slightly different value for seed
	r := rand.New(source)                           // Creating a new rand variable using that seed
	for index := range d {
		newPosition := r.Intn(len(d) - 1) // Sudo random generator, uses same seed everytime
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
