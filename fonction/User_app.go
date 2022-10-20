package fonction

import "fmt"

func User_App() string {
	var lettre string
	fmt.Println("choisie la lettre: ")
	fmt.Scanln(&lettre)
	return lettre
}
