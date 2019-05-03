package main

import (
	"log"
	"os"

	"github.com/shful/gofp"
)

// main parses a pizza ontology.
// The pizza ontology is from https://github.com/owlcs/pizza-ontology
// which was converted to OWL-Functional with Protègè (https://protege.stanford.edu/)
func main() {
	f, err := os.Open("pizza-functional.owl")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	o, err := gofp.OntologyFromReader(f, "the pizza-functional.owl file")
	if err != nil {
		log.Fatal(gofp.ErrorMsgWithPosition(err))
	}
	log.Println("That's what we parsed:", o.About())
}
