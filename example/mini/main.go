package main

import (
	"fmt"
	"log"
	"strings"

	"reifenberg.de/gofp"
	"reifenberg.de/gofp/owlfunctional/ontologies"
)

// main parses an example string and prints what classes and properties were parsed.
// This is the example from the README.TXT.
func main() {
	var err error
	var o *ontologies.Ontology

	// parse an OWL string into an ontology struct:
	o, err = gofp.OntologyFromReader(strings.NewReader(`
			Prefix(:=<http://www.example.org/gofphelloworld#>)
			Prefix(rdfs:=<http://www.w3.org/2000/01/rdf-schema#>)
			
			Ontology(<urn:absolute:example.org>
			
				# Classes
				Declaration(Class(:Pizza))
				Declaration(Class(:MozzarellaTopping))
				Declaration(Class(:MargheritaPizza))
			
				# Object Properties
				Declaration(ObjectProperty(:hasTopping))
			
				# some Axioms
				AnnotationAssertion(rdfs:comment :MargheritaPizza "Pizza with Tomato and Mozzarella on top")
				ObjectPropertyDomain(:hasTopping :Pizza)
			
				SubClassOf(:MargheritaPizza :Pizza)
				SubClassOf(:MargheritaPizza ObjectSomeValuesFrom(:hasTopping :MozzarellaTopping))
			)
		`),
		"pizza demo data",
	)

	if err != nil {
		log.Fatal(gofp.ErrorMsgWithPosition(err))
	}

	fmt.Println("All declared class names are:")
	for name, _ := range o.AllClassDecls {
		fmt.Println(name)
	}
}
