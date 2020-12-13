package pipe

func New(input <-chan string, output chan<- string, f func(s string) string) {
	for inputString := range input {
		output <- f(inputString)
	}
}
