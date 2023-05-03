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
	Words      string
	RandomWord string
	Attempts   int
}

func Call() { // appelle toute les fonctions dans une fonction
	RW := RandomWord
	fmt.Println("Good Luck, you have 10 attempts.")
	words := RW(Scantxt("words.txt"))
	void := Underscore(words)
	fmt.Println(words)
	fmt.Println(void)

	for {

		test := Replace(void, words, User_App())
		void = test
		fmt.Println(void)
		//Win(test, void)
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
func Replace(void, words, lettre string) string { // fonction qui remplace un "_" par la lettre à la bonne place
	rune_void := []rune(void)
	for i := 0; i < len(words); i++ {
		if len(lettre) == 1 {
			if lettre[0] == words[i] && rune_void[i] == '_' {
				rune_void[i] = rune(lettre[0])
			}
		}
	}
	return string(rune_void)
}
func Scantxt(fileName string) []string { // fonction qui scan les mots du txt et affiche un mots aléatoire.
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
func Underscore(words string) string { // fonction qui parcours le mots en "_"
	Rwords := []rune(words)
	for i := 0; i < len(words); i++ {
		Rwords[i] = '_'
	}
	return string(Rwords)
}
func User_App() string { // fonction qui permet de d'ecrire dans le terminal
	var lettre string
	fmt.Println("choisie la lettre: ")
	fmt.Scanln(&lettre)
	return lettre
}

func Compare(lettre string) bool { // renvoie true si le mot est totallement decouvert sinon false
	word := []rune(lettre)
	for i := 0; i < len(word); i++ {
		if word[i] == '_' {
			return false
		}
	}
	return true
}
