package main

func add() int {
	var a, b, c int = 1, 2, 0
	c = (a + b)
	return c
}
func main() {
	println(add())
}
