package fonction

import "fmt"

func attemps(word string) {
	void := Underscore(word)
	attempts := 10
	for attempts > 0 {
		var lettre string
		fmt.Print("Rentrez une lettre : ")
		_, err := fmt.Scanln(&lettre)
		if err != nil {
			attempts--
		}
		void = Remplacement(void, word, lettre)
		fmt.Println(void)
	}
}
