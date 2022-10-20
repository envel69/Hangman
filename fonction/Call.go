package fonction

import (
	"fmt"
)

func Call() {
	RW := RandomWord
	// tab := Ligneparligne("../words.txt")
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(Scantxt("words.txt"))
	void := Underscore(words)
	fmt.Println(void)
}
