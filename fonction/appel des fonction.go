package fonction

import "fmt"

func appel() {
	RW := fonction.RandomWord
	// tab := Ligneparligne("../words.txt")
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(fonction.Ligneparligne("../words.txt"))
	void := fonction.Underscore(words)
	fmt.Println(void)
	// fmt.Println(Underscore(RW(Ligneparligne("../words.txt"))))
	// fmt.Println(RandomWord(ligneparligne("words.txt")))
}
