package main

func main() {
	vote(0, 0, 1)
}

func vote(x int, y int, z int) int {
	if (x + y + z) >= 2 {
		return 1
	} else {
		return 0
	}
}
