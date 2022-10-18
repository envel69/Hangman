package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	RW := RandomWord
	fmt.Println("Good Luck, you have 10 attempts.")
	fmt.Println(Underscore(RW(Ligneparligne("../words.txt"))))
	//fmt.Println(RandomWord(ligneparligne("words.txt")))
	fmt.Println(User_app())
}

func Random(n int) int {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}

func RandomWord(words []string) string {
	return words[Random(len(words))]
}

func Underscore(RW string) string {
	var u string
	fmt.Println(RW)

	for range RW {
		u += "_ "
	}
	return u
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

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func Remplacement(input, string, ToFind, ModifiedWord) string {
	for i := 0; i < len(input); i++ {
		if ToFind != ModifiedWord {
			i++
		}
	}
}
