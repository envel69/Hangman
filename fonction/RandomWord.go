package fonction

func RandomWord(words []string) string { //liste de mots en tableau
	return words[Random(len(words))]
}
