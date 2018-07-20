package main

type A struct {
	child [2]A
}

func main() {
	a := A{}
	print(a)
}
