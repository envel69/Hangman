package main

import (
	"fmt"
	"hangman/fonction"
)

func main() { //fonction main
	RW := fonction.RandomWord                                                 //racourci de RandomWord pour simplifier
	fmt.Println(fonction.Underscore(RW(fonction.LigneParLigne("words.txt")))) // ecriture du mot en tiret
	fmt.Println(fonction.LineByLine("hangman.txt"))                           //affichage de josé
	fonction.User_App()                                                       //input de user app
}
