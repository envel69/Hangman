package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println(RandomWord(ligneparligne("words.txt")))
}

func random() {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	fmt.Println(x1)
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
