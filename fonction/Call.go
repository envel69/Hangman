package fonction

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Hangman struct {
	Words      string
	RandomWord string
	Attempts   int
}

var hangman Hangman

func Call() {
	hangman.Attempts = 10

	RW := RandomWord
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(Scantxt("words.txt"))
	void := Underscore(words)
	fmt.Println(void)
	// fmt.Println(words)
	//fmt.Println(void)
	fmt.Println(compare(void))
	Attemp(hangman, words)

}
func Attemp(hangman Hangman, word string) {
	void := Underscore(word)
	//couble de jeu
	for hangman.Attempts > 0 {
		var lettre string
		fmt.Print("Rentrez une lettre : ")
		fmt.Scanln(&lettre)
		hangman, void = Replace(void, word, lettre)
		fmt.Println(void)
		Jose(hangman.Attempts)

		if hangman.Attempts == 0 || compare(void) {
			break
		}
	}
}
func Capitalize(min string) string {
	var maj = strings.ToUpper(min)
	return (maj)
}
func Jose(numberOfAttemps int) { //fonction qui affiche la position du pendu par rapport au nombre de vie restante
	var arrayJose []byte
	pos := 10 - numberOfAttemps
	if pos == 0 {
		fmt.Println("\n\n\n\n\n\n\n\n")
	} else {
		content, _ := ioutil.ReadFile("hangman.txt")
		for i := 0; i < 71; i++ {
			arrayJose = append(arrayJose, content[i+(71*(pos-1))])
		}
		jose := (string(arrayJose))
		fmt.Println(jose)
	}

}
func Random(n int) int {
	y1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(y1)
	return x1.Intn(n)
}
func RandomWord(words []string) string {
	return words[Random(len(words))]
}
func Replace(void, words, lettre string) (Hangman, string) {
	rune_void := []rune(void)
	var faute bool = false
	// fmt.Println("lettre : ", lettre)

	for i := 0; i < len(words); i++ {
		if len(lettre) == 1 {
			if (lettre[0] == words[i] && rune_void[i] == '_') || lettre[0] == byte(rune_void[i]) {
				rune_void[i] = rune(lettre[0])
				faute = true
			}
		}
	}
	if strings.ToUpper(lettre) == strings.ToUpper(words) {
		rune_void = []rune(words)
		faute = true
	}
	if faute == false {
		hangman.Attempts -= 1
	}
	// fmt.Println("words : ", words)
	// fmt.Println(string(rune_void))
	return hangman, string(rune_void)
}
func Scantxt(fileName string) []string {
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
func Underscore(words string) string {
	Rwords := []rune(words)
	for i := 0; i < len(words); i++ {
		Rwords[i] = '_'
	}
	return string(Rwords)
}

func Win(test, void string) string {
	return "nice"
}
func compare(lettre string) bool { // renvoie true si le mot est totallement decouvert sinon false
	word := []rune(lettre)
	// fmt.Println("lettre : ", lettre)
	for i := 0; i < len(word); i++ {
		if word[i] == '_' {
			return false
		}
	}
	return true
}
