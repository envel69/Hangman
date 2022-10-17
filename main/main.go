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

func random(n int) int {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}

// func RandomWord(words []string) string {
// 	return words[random(len(words))]
// }

func RandomWord2(words2 []string) string {
	return words2[random(len(words2))]
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
