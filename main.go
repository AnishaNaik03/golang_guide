package main

func main() {
	// var card string = "ace of spades"
	// card := "hello"
	// card = "hey"
	// card := Newcard()
	// i := 123
	// fmt.Println(i)

	//type deck
	// cards := deck{"ACe of Diamonds", Newcard()}
	// cards := []string{"ACe of Diamonds", Newcard()}
	// cards = append(cards, "six of heart")
	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards.toString())
	// cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// greeting := "hi there!"
	// fmt.Println([]byte(greeting))

	//saving to file
	// cards := newDeck()
	// cards.saveTofile("my_cards")

	//reading from file
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}

// func Newcard() string {
// 	return "Five Of Diamonds"
// }
