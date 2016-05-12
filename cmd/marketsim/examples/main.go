package main

import (
	"flag"
)

// An example provides a run() method so that this main
// function can execute it depending on the command line
// arguments.
type Example interface {
	run()
}

func run(ex Example) {
	ex.run()
}

var flagMapping = map[string]Example {
	"simple": SimpleExample{},
	"full": FullNasdaqSim{}}

// User can pass -example=full to run the full nasdaq sim
// otherwise it defaults to the simple example.
func main() {
	var example = flag.String("example", "simple", "Pass either -example=simple or -example=full")
	flag.Parse()

	var ex = flagMapping[*example]
	run(ex)
}
