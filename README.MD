### Gofp Owl-Functional Parser


Gofp reads OWL files with functional syntax.

The resulting structures strictly resemble the OWL-Functional structures. For example, a SubClassOf axiom
ends up in a struct instance of type "axioms.SubClassOf". Everything is read into memory at once.

The package structure, as well as package and type names in Gofp depict the OWL documentation from here: https://www.w3.org/2007/OWL/refcard

The implementation idea comes from https://blog.gopheracademy.com/advent-2014/parsers-lexers/. 



#### Example for parsing a file:
```	go
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
```

To run that example, cd into the gofp/example/pizza/ directory and type
> go run main.go


#### How to access the parsed data ?
We get an ontology.Ontology instance from the parser, and, by default, this has an Ontology.K attribute with exactly all parsed knowledge (OWL axioms and declarations).
All elements are returned by the "All"-prefixed methods therein, like "AllSubClassOfs()." Additionally, Declarations can be taken by their IRI - for example, ClassDecl("example.com/Pizza").
While this is the default, Gofp can parse directly into custom types, alternatively. See also the parameter documentation of the ontologies.NewOntology() function.


#### Caveats
The implementation is not complete. The "import" statement is unknown and breaks parsing.
Annotations and free text inside an Ontology element are unknown and break parsing.
Some more statements and datatypes are unknown; most of these come from the "Individual" and "Annotation" categories.

The current Gofp has further expectations that clash with OWL. Gofp needs all classes, individuals and properties explicitly declared with "Declaration(...)" before the first reference. Further, all input must be UTF-8.


#### Recent API changes
Note that the API may continue to change, and Gofp does not use a semantic version yet. 
Since commit 018e1d from Apr 23 2019, the parsed elements are not found in Ontology.All* - slices and maps anymore. See above, "How to access the parsed data" for details.
That API change was made to (optionally) parse directly into custom types.


#### Prerequisites
A Go compiler and the standard library.


#### Remarkable OWL parser projects, choosen arbitrarily
* owlapi (Interacting with OWL Ontologies, Java) [https://github.com/owlcs/owlapi]
* jopa (Java OWL Persistence API, Java) [https://github.com/kbss-cvut/jopa]
* jop (Jena OWL Parser, Java) [https://github.com/daveti/jop]
* pronto (Python frontend to Ontologies, Python) [https://github.com/althonos/pronto]


#### License
Gofp is MIT licensed.