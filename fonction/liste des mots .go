package fonction

import (
	"bufio"
	"os"
)

func LigneParLigne(fileName string) []string { //liste des mot ligne par ligne
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
