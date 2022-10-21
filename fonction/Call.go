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

type hangman struct {
	words      string
	RandomWord string
	attempts   int
}

func Call() {
	RW := RandomWord
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(Scantxt("words.txt"))
	void := Underscore(words)
	fmt.Println(void)
	// fmt.Println(words)
	//fmt.Println(void)
	fmt.Println(compare(void))

	for {

		test := Replace(void, words, User_App())
		void = test
		fmt.Println(void)
		//Win(test, void)
		if compare(void) {
			break
		}
	}

}
func Attemp(word string) {
	void := Underscore(word)
	attempts := 10
	for attempts > 0 {
		var lettre string
		fmt.Print("Rentrez une lettre : ")
		_, err := fmt.Scanln(&lettre)
		if err != nil {
			attempts--
		}
		void = Replace(void, word, lettre)
		fmt.Println(void)
	}
}
func Capitalize(min string) string {
	var maj = strings.ToUpper(min)
	return (maj)
}
func Jose(numberOfAttemps int) { //fonction qui affiche la position du pendu par rapport au nombre de vie restante
	var arrayJose []byte
	initNumberOfLife := 10
	pos := initNumberOfLife - numberOfAttemps
	content, err := ioutil.ReadFile("josé.txt")

	if err != nil {
		fmt.Println("re essaye")
	} else {
		for i := 0; i < 71; i++ {
			arrayJose = append(arrayJose, content[i+(71*pos)])

		}
		josé := (string(arrayJose))
		fmt.Println(josé)

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
func Replace(void, words, lettre string) string {
	rune_void := []rune(void)
	// fmt.Println("lettre : ", lettre)
	for i := 0; i < len(words); i++ {
		if len(lettre) == 1 {
			if lettre[0] == words[i] && rune_void[i] == '_' {
				rune_void[i] = rune(lettre[0])
			}
		}
	}
	// fmt.Println("words : ", words)
	// fmt.Println(string(rune_void))
	return string(rune_void)
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
func User_App() string {
	var lettre string
	fmt.Println("choisie la lettre: ")
	fmt.Scanln(&lettre)
	return lettre
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
