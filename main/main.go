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
	fmt.Println(underscore(RW(ligneparligne("words.txt"))))
	//fmt.Println(RandomWord(ligneparligne("words.txt")))
	fmt.Println(user_app())
}

func random(n int) int {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}

func RandomWord(words []string) string {
	return words[random(len(words))]
}

func underscore(RW string) string {
	var u string
	fmt.Println(RW)

	for range RW {
		u += "_ "
	}
	return u
}

func ligneparligne(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func user_app() string {
	var lettre string
	fmt.Scanln(&lettre)
	return lettre
}
