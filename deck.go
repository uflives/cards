package main

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, main := range cardValues {
		for _, sub := range cardSuits {
			name := main + " of " + sub
			cards = append(cards, name)

		}
	}

	return cards
}
