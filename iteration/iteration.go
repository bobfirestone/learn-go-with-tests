package iteration

// Repeat it repeats
func Repeat(times int, character string) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
