package main

func main() {

}

func sumInt(n ...int) (int, int) {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return len(n), sum
}
