package main

func three() {
	panic("so deep")
}

func two() {
	three()
}

func one() {
	two()
}

func main() {
	one()
}
