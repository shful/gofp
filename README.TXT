Gofp Owl-Functional Parser
===================================

Gofp reads OWL files with Functional syntax into memory structures for further processing.
The resulting structures strictly resemble the OWL-Functional structures. For example, any SubClassOf axiom
ends up in a struct instance of type "axioms.SubClassOf".
The package structure, as well as package and type names in Gofp depict the OWL documentation from here: https://www.w3.org/2007/OWL/refcard
The implementation idea comes fro https://blog.gopheracademy.com/advent-2014/parsers-lexers/. Everything is read into memory at once.


Example for parsing a file:
	
    f, _ := os.Open("pizza-functional.owl")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ontology, err := gofp.OntologyFromReader(f, "source was pizza-functional.owl")
	if err != nil {
		log.Fatal(gofp.ErrorMsgWithPosition(err))
	}
	log.Println("That's what we parsed:", ontology.About())


To run that example, cd into the gofp/example/pizza/ directory and type
  go run main.go


How to access the parsed data ?
You will notice there's not much of an API yet. We get an ontology.Ontology instance from the parser, and
all parsed structures are found in (array) members therein, for example in "AllSubClassOfs".

The implementation is not complete. The "import" statement is unknown and breaks parsing.
Multiple IRIs, Annotations and free text inside an Ontology element are unknown and break parsing.
Some more statements and datatypes are unknown; most of these come from the "Individual" and "Annotation" categories.
The current Gofp has further expectations that clash with OWL. Gofp needs all classes, individuals and properties explicitly declared with "Declaration(...)" before the first reference. Further, all input must be UTF-8.