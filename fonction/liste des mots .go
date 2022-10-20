package fonction

import (
	"bufio"
	"os"
)

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
