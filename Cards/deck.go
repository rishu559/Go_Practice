package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a custom datatype equivalent to slice of string
type deck []string

// This is called a reciever function to print the cards
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Function to create new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := deck{"Heart", "Spade", "Diamonds", "Club"}
	cardNumbers := deck{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}

	// for i := 0; i < len(cardSuits); i++ {
	// 	for j := 0; j < len(cardNumbers); j++ {
	// 		cards = append(cards, cardNumbers[j]+" of "+cardSuits[i])
	// 	}
	// }

	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+suit)
		}
	}
	return cards
}

// Deal function to divide the deck

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err1 := f.Write([]byte(d.toString()))

	if err1 != nil {
		return err1
	}

	return errors.New("Done")

}

func newDeckFromFile(fileName string) deck {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return deck(strings.Split(string(content), ","))

}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(source)

	for i := 0; i < len(d); i++ {
		randomIndex := myRand.Intn(len(d) - 1)
		// To swap the two numbers
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}

