package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// fmt.Println(ligneparligne("words.txt"))
	fmt.Println(RandomWord(ligneparligne("words.txt")))
}

func RandomWord(words []string) string {
	return words[rand.Intn(len(words))]
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
