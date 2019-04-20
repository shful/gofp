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
You will notice there's not much of an API yet. We get an ontology.Ontology instance from the parser, and
all parsed structures are found in (slice and map) members therein, accessible with methods like AllSubClassOfs(), or AllClassDecls(), or ClassDecl("example.com/Pizza"). These method definitions are in the "Axioms" and "Decls" interfaces of the gofp/store package.


#### Caveats
The implementation is not complete. The "import" statement is unknown and breaks parsing.
Annotations and free text inside an Ontology element are unknown and break parsing.
Some more statements and datatypes are unknown; most of these come from the "Individual" and "Annotation" categories.

The current Gofp has further expectations that clash with OWL. Gofp needs all classes, individuals and properties explicitly declared with "Declaration(...)" before the first reference. Further, all input must be UTF-8.


### Recent API changes
The current API may change and Gofp does not use a semantic version number yet. Note that since commit 8589ce from Apr 17 2019, the parsed elements are not in Ontology.All* - slices and maps anymore. Instead, we got functions prefixed with "All*" which return slices for both axioms and declarations - for example, Ontology.AllClassDecls(). Declarations, additionally, can be taken by IRI - for example, with Ontology.ClassDecl(iri string).


#### Prerequisites
A Go compiler and the standard library.


#### Remarkable OWL parser projects, choosen arbitrarily
* owlapi (Interacting with OWL Ontologies, Java) [https://github.com/owlcs/owlapi]
* jopa (Java OWL Persistence API, Java) [https://github.com/kbss-cvut/jopa]
* jop (Jena OWL Parser, Java) [https://github.com/daveti/jop]
* pronto (Python frontend to Ontologies, Python) [https://github.com/althonos/pronto]


#### License
Gofp is MIT licensed.