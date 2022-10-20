package fonction

import (
	"fmt"
)

func Call() {
	RW := RandomWord
	// tab := Ligneparligne("../words.txt")
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(Ligneparligne("../words.txt"))
	void := Underscore(words)
	fmt.Println(void)
	// fmt.Println(Underscore(RW(Ligneparligne("../words.txt"))))
	// fmt.Println(RandomWord(ligneparligne("words.txt")))
}
