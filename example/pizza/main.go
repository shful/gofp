package main

import (
	"log"
	"os"

	"reifenberg.de/gofp"
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

	ontology, err := gofp.OntologyFromReader(f, "source was pizza-functional.owl")
	if err != nil {
		log.Fatal(gofp.ErrorMsgWithPosition(err))
	}
	log.Println("That's what we parsed:", ontology.About())
}
