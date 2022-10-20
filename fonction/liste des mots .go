package fonction

import (
	"bufio"
	"fmt"
	"os"
)

func Ligneparligne(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//fmt.Println(lines)
	return lines
}
