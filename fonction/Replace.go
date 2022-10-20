package fonction

import "fmt"

func Replace(void, words, lettre string) string {
	rune_void := []rune(void)
	fmt.Println("lettre : ", lettre)
	for i := 0; i < len(words); i++ {
		if len(lettre) == 1 {
			if lettre[0] == words[i] && rune_void[i] == '_' {
				rune_void[i] = rune(lettre[0])
			}
		}
	}
	fmt.Println("words : ", words)
	fmt.Println(string(rune_void))
	return string(rune_void)
}
