package main

import (
	"fmt"
	"fonction/fonction"
)

// S'il n'y a pas de lettre modifiée, retire 1 à attempts
// S'il n'y a plus de '_' c'est gagné
func main() {
	RW := fonction.RandomWord
	// tab := Ligneparligne("../words.txt")
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(fonction.Ligneparligne("../words.txt"))
	void := fonction.Underscore(words)
	fmt.Println(void)
	// fmt.Println(Underscore(RW(Ligneparligne("../words.txt"))))
	// fmt.Println(RandomWord(ligneparligne("words.txt")))
	attempts := 10
	for attempts > 0 {
		var lettre string
		fmt.Print("Rentrez une lettre : ")
		_, err := fmt.Scanln(&lettre)
		if err != nil {
			attempts--
		}
		void = fonction.Remplacement(void, words, lettre)
		fmt.Println(void)
	}
}
