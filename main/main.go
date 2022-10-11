package main

import (
	"bufio"
	"math/rand"
	"os"
)

func main() {
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
	fmtprintln.(lines)
	return lines
}
