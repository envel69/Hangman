package fonction

func Underscore(words string) string {
	Rwords := []rune(words)
	for i := 0; i < len(words); i++ {
		Rwords[i] = '_'
	}
	return string(Rwords)
}
