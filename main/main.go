package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 12
	fmt.Println(rand.Intn(max-min+1) + min)
}

func RandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

func ReadFileLineByLine(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
