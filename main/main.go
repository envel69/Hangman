package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// S'il n'y a pas de lettre modifiée, retire 1 à attempts
// S'il n'y a plus de '_' c'est gagné

func main() {
	RW := RandomWord
	// tab := Ligneparligne("../words.txt")
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(Ligneparligne("../words.txt"))
	void := Underscore(words)
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
		void = Remplacement(void, words, lettre)
		fmt.Println(void)
	}
}

func Random(n int) int {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}

func RandomWord(words []string) string {
	return words[Random(len(words))]
}

func Underscore(words string) string {
	Rwords := []rune(words)

	for i := 0; i < len(words); i++ {
		Rwords[i] = '_'
	}
	return string(Rwords)
}

func Ligneparligne(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//fmt.Println(lines)
	return lines
}

func User_app() string {
	var lettre string
	fmt.Scanln(&lettre)
	return lettre
}

func Remplacement(void, words, lettre string) string {
	rune_void := []rune(void)
	fmt.Println("lettre : ", lettre)
	for i := 0; i < len(words); i++ {
		if len(lettre) == 1 {
			if lettre[0] == words[i] && rune_void[i] == '_' {
				rune_void[i] = rune(lettre[0])
			}
		}
	}
	fmt.Println("words : ", words)
	fmt.Println(string(rune_void))
	return string(rune_void)
}
