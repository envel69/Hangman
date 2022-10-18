package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() { //fonction main
	RW := RandomWord //racourci de RandomWord pour simplifier
	//fmt.Println(RW(ligneparligne("words.txt")))//marche pas je sais pas pk voir avec envel!!!!!
	fmt.Println(user_app())                                 //a trouvé pk ca marche pas PUTIN!!!
	fmt.Println(underscore(RW(ligneparligne("words.txt")))) // ecriture du mot en tiret

}

func random(n int) int { //fonciont random avec le temp en second
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}

func RandomWord(words []string) string { //liste de mot
	return words[random(len(words))]
}

func ligneparligne(fileName string) []string { //liste des mot ligne par ligne
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func underscore(RW string) string { // boucle pour transformé le mot en tiret
	var u string
	//for i := 0; i < len(RW); i++ {
	for range RW {
		u += "_"
	}
	return u
}

func user_app() string {
	var lettre string
	fmt.Println("choisie la lettre: ")
	fmt.Scanln(&lettre)
	return lettre
}
