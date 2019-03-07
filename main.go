// package is a collection of common sourse code files
package main

func main() {
	// // this is a 1 way of declaring a variable that is statically
	// // var card string = "ace of spade"
	// // 2nd way Dynamically declaring a variable

	// card := "Ace of Spades"
	// // Note only first time this collan is used after that it is not like what we see down
	// card = "Five of Dimonds"
	// fmt.Println(card)
	//cards := deck{"Ace of Dimonds", newCard()}
	//cards = append(cards, "Six of Spade")
	// cards := newDeck()
	//cards.print()
	// 	hand, remainingCards := deal(cards, 5)
	// 	hand.print()
	// 	remainingCards.print()
	// }

	// func newCard() string {
	// 	return "Five of Dimonds"
	//fmt.Println(cards.toString())
	//cards.savetoFile("my_cards")
	// cards := newDeckFromFile("my_cards")

	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
