package fonction

func Underscore(RW string) string { // boucle pour transformé le mot en tiret
	var u string
	//for i := 0; i < len(RW); i++ {
	for range RW {
		u += "_ "
	}
	return u
}
