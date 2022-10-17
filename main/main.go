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

	fmt.Println(underscore(RW(ligneparligne("words.txt"))))

}

func random(n int) int {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}

func RandomWord(words []string) string {
	return words[random(len(words))]
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

func underscore(RW string) string {
	var u string
	//for i := 0; i < len(RW); i++ {
	for range RW {
		u += "_"
	}
	return u
}
