package main

type Example interface {
	run()
}

func run(ex Example) {
	ex.run()
}

func main() {
	var ex1 FullNasdaqSim
	var ex2 SimpleExample
	run(ex1)
	run(ex2)
}
