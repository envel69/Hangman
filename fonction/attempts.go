package fonction

import "fmt"

func attemps() {
	void := fonction.Underscore(words)
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
