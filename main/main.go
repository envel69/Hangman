package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println(RandomWord(linebyline("words.txt")))
	// r := rand.Intn(3) + 1
}
func RandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}
func linebyline(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
