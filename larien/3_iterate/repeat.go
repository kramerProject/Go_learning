package iteracao

func Repetir(txt string, maxRep int) string {
	var repetitions string
	for i := 0; i < maxRep; i++ {
		repetitions += txt
	}
	return repetitions
}
