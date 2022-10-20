package fonction

import "fmt"

func User_app() string {
	var lettre string
	fmt.Println("choisie la lettre: ")
	fmt.Scanln(&lettre)
	return lettre
}
