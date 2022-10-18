package fonction

func RandomWord(words []string) string {
	return words[Random(len(words))]
}
