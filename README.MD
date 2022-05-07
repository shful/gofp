### Gofp Owl-Functional Parser


Gofp reads OWL2 files with functional syntax.

The resulting structures strictly resemble the OWL-Functional structures. For example, a SubClassOf axiom
is parsed into an `axioms.SubClassOf` instance. Everything is read into memory at once.

The project structure and the package and type names depict the OWL documentation from here: https://www.w3.org/2007/OWL/refcard

The implementation idea comes from https://blog.gopheracademy.com/advent-2014/parsers-lexers/. 



#### Example for parsing a file:
```	go
f, err := os.Open("pizza-functional.owl")
if err != nil {
	panic(err)
}
defer f.Close()

o, err := gofp.OntologyFromReader(f, "source was pizza-functional.owl")
if err != nil {
	log.Fatal(gofp.ErrorMsgWithPosition(err))
}
log.Println("That's what we parsed:", o.About())
```

To run that example, cd into the gofp/example/pizza/ directory and type
> go run main.go


#### How to access the parsed data ?
We get an `owlfunctional.Ontology` instance from the parser. By default, this has an `Ontology.K` attribute with all parsed knowledge, which is made up of OWL axioms and declarations.
All parsed elements are accessible by the "All"-prefixed methods here, like `AllSubClassOfs()`. Additionally, all declarations are accessible by their IRI, for example `ClassDecl("example.com/Pizza")`.

```
// Example: print all parsed class declaration IRIs:
for _, declaration := range o.K.AllClassDecls() {
	fmt.Println(declaration.IRI)
}
```

While this is the default, Gofp can parse directly into custom types, alternatively. See also the parameter documentation of the `owlfunctional.NewOntology` function.


#### Caveats
The implementation is not complete. The "import" statement is unknown and breaks parsing.
Annotations and free text inside an Ontology element are unknown and break parsing.
Some more statements and datatypes are unknown; most of these come from the "Individual" and "Annotation" categories. A  management for IRI prefixes is missing.
Further, all input must be UTF-8.


#### Recent API changes
Note that the API may continue to change. Gofp, by intention, has a v0.* version (see https://blog.golang.org/publishing-go-modules for golang versioning).
* With commit 6e35dd from Oct 08 2019, the import path of "owlfunctional/declarations" was changed to "owlfunctional/decl", and "owlfunctional/ontologies" was shortened to "owlfunctional".
* Since commit 018e1d from Apr 23 2019, the parsed elements are not found in Ontology.All* - slices and maps anymore. See above, "How to access the parsed data" for details. That API change was made to (optionally) parse directly into custom types.


#### Prerequisites
A Go compiler and the standard library.


#### Remarkable OWL parser projects, choosen arbitrarily
* owlapi (Interacting with OWL Ontologies, Java) [https://github.com/owlcs/owlapi]
* jopa (Java OWL Persistence API, Java) [https://github.com/kbss-cvut/jopa]
* jop (Jena OWL Parser, Java) [https://github.com/daveti/jop]
* pronto (Python frontend to Ontologies, Python) [https://github.com/althonos/pronto]


#### License
Gofp is MIT licensed.
