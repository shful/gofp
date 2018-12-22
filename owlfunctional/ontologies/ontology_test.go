package ontologies

import (
	"testing"

	"reifenberg.de/gofp/mock"
	"reifenberg.de/gofp/owlfunctional/annotations"
	"reifenberg.de/gofp/owlfunctional/axioms"
	"reifenberg.de/gofp/owlfunctional/builtindatatypes"
	"reifenberg.de/gofp/owlfunctional/declarations"
	"reifenberg.de/gofp/owlfunctional/facets"
	"reifenberg.de/gofp/owlfunctional/parser"
)

func TestParsePizzaOntology(t *testing.T) {
	var p *parser.Parser
	var err error
	p = mock.NewTestParser(ontologyTestString)
	o := NewOntology(map[string]string{"": "localprefix", "hello": "hello.de", "xsd": builtindatatypes.PRE_XSD, "rdfs": builtindatatypes.PRE_RDFS})

	parser.TokenLog = true
	err = o.Parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if o.IRI != "<urn:absolute:test.de>" {
		t.Fatal(o.IRI)
	}

	// === Decls
	if len(o.AllClassDecls) != 50 {
		t.Fatal(o.AllClassDecls)
	}
	if !o.ClassDeclExists("localprefix#AmericanHotPizza") {
		t.Fatal()
	}
	if !o.ClassDeclExists("hello.de#FishbonePizza") {
		t.Fatal()
	}
	if o.ClassDeclExists("notexisting") {
		t.Fatal()
	}

	if len(o.AllDataPropertyDecls) != 1 {
		t.Fatal(o.AllDataPropertyDecls)
	}
	if !o.DataPropertyDeclExists("localprefix#hasCaloricContentValue") {
		t.Fatal()
	}
	if o.DataPropertyDeclExists("localprefix#hasTopping") {
		t.Fatal()
	}
	if o.DataPropertyDeclExists("localprefix#HasTopping") { // case differs
		t.Fatal()
	}
	if len(o.AllObjectPropertyDecls) != 7 {
		t.Fatal(o.AllObjectPropertyDecls)
	}
	if !o.ObjectPropertyDeclExists("localprefix#hasTopping") {
		t.Fatal()
	}

	if len(o.AllNamedIndividualDecls) != 4 {
		t.Fatal(o.AllNamedIndividualDecls)
	}
	if !o.NamedIndividualDeclExists("localprefix#MyQuattroFormaggio") {
		t.Fatal()
	}

	// === SubObjectPropertyOfs
	if len(o.AllSubObjectPropertyOfs) != 2 {
		t.Fatal(o.AllSubObjectPropertyOfs)
	}
	{
		s := o.AllSubObjectPropertyOfs[0]
		switch x := s.P1.(type) {
		case *declarations.ObjectPropertyDecl:
			if x.IRI != "localprefix#hasBase" {
				t.Fatal(x)
			}
		default:
			t.Fatal(s.P1)
		}
	}

	// InverseObjectProperties
	if len(o.AllInverseObjectProperties) != 3 {
		t.Fatal(o.AllInverseObjectProperties)
	}
	{
		s := o.AllInverseObjectProperties[0] // the slice preserves the statement order
		x := s.P1.(*declarations.ObjectPropertyDecl)
		if x.IRI != "localprefix#hasBase" {
			t.Fatal(x)
		}
	}

	// ObjectPropertyDomain
	if len(o.AllObjectPropertyDomains) != 2 {
		t.Fatal(o.AllObjectPropertyDomains)
	}
	{
		s := o.AllObjectPropertyDomains[0]
		x := s.C.(*declarations.ClassDecl)
		if x.IRI != "localprefix#Pizza" {
			t.Fatal(x)
		}
		y := s.P.(*declarations.ObjectPropertyDecl)
		if y.IRI != "localprefix#hasBase" {
			t.Fatal(y)
		}
	}

	// ObjectPropertyRange
	if len(o.AllObjectPropertyRanges) != 3 {
		t.Fatal(o.AllObjectPropertyRanges)
	}
	{
		s := o.AllObjectPropertyRanges[0]
		x := s.C.(*declarations.ClassDecl)
		if x.IRI != "localprefix#PizzaBase" {
			t.Fatal(x)
		}
		y := s.P.(*declarations.ObjectPropertyDecl)
		if y.IRI != "localprefix#hasBase" {
			t.Fatal(y)
		}
	}

	// DataPropertyRange
	if len(o.AllDataPropertyRanges) != 1 {
		t.Fatal(o.AllDataPropertyRanges)
	}
	{
		s := o.AllDataPropertyRanges[0]
		x := s.R.(*declarations.DataPropertyDecl)
		if x.IRI != "localprefix#hasCaloricContentValue" {
			t.Fatal(x)
		}
		y := s.D.(*facets.BuiltinDatatype)
		if y.DatatypeIRI != builtindatatypes.PRE_XSD+"#integer" {
			t.Fatal(y)
		}
	}

	// SubClassOf
	if len(o.AllSubClassOfs) != 67 {
		t.Fatal(len(o.AllSubClassOfs))
	}
	{
		s := o.AllSubClassOfs[0]
		x := s.C1.(*declarations.ClassDecl)
		if x.IRI != "hello.de#FishbonePizza" {
			t.Fatal(x)
		}
	}
}

func TestParseEquivalentClasses(t *testing.T) {
	var p *parser.Parser
	var err error
	var o *Ontology = NewOntology(map[string]string{})
	o.Prefixes[""] = "localprefix"
	o.AllClassDecls[`localprefix#Pizza`] = &declarations.ClassDecl{declarations.Declaration{IRI: "localprefix#Pizza"}}
	o.AllClassDecls[`localprefix#InterestingPizza`] = &declarations.ClassDecl{declarations.Declaration{IRI: "localprefix#InterestingPizza"}}
	o.AllObjectPropertyDecls[`localprefix#hasTopping`] = &declarations.ObjectPropertyDecl{Declaration: declarations.Declaration{IRI: "localprefix#hasTopping"}}

	p = mock.NewTestParser(`EquivalentClasses(:InterestingPizza ObjectIntersectionOf(:Pizza ObjectMinCardinality(3 :hasTopping)))	`)
	// parser.TokenLog = true

	err = o.parseEquivalentClasses(p)
	if err != nil {
		t.Fatal(err)
	}
	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}

	if len(o.AllEquivalentClasses) != 1 {
		t.Fatal(o.AllEquivalentClasses)
	}

	var expr axioms.EquivalentClasses
	expr = o.AllEquivalentClasses[0]
	if len(expr.EquivalentClasses) != 2 {
		t.Fatal(o.AllEquivalentClasses)
	}
}

func TestParseAnnotationAssertion(t *testing.T) {

	var p *parser.Parser
	var err error
	var expr annotations.AnnotationAssertion
	var o *Ontology = NewOntology(map[string]string{})
	o.Prefixes[""] = "The local ns"
	o.Prefixes["xsd"] = "The xsd-ns"
	o.Prefixes["rdfs"] = "The rdfs-ns"
	o.Prefixes["pizza"] = "The-Pizza-Namespace"
	o.AllClassDecls[`The local ns:MargheritaPizza`] = &declarations.ClassDecl{declarations.Declaration{IRI: "The local ns#MargheritaPizza"}}

	// 3rd param in AnnotationAssertion can be IRI, literal, or anonymous individual

	// with 3rd param = literal:
	p = mock.NewTestParser(`AnnotationAssertion(rdfs:comment :MargheritaPizza "Pizza from Tomato and Mozzarella"^^xsd:string)`)
	parser.TokenLog = true
	err = o.parseAnnotationAssertion(p)
	if err != nil {
		t.Fatal(err)
	}

	if err = p.ConsumeTokens(parser.EOF); err != nil {
		t.Fatal(err)
	}
	if len(o.AllAnnotationAssertions) != 1 {
		t.Fatal(o.AllAnnotationAssertions)
	}
	expr = o.AllAnnotationAssertions[0]
	if expr.T != `"Pizza from Tomato and Mozzarella"^^The xsd-ns#string` {
		t.Fatal(expr.T)
	}

	// with 3rd param = IRI without prefix:
	p = mock.NewTestParser(`AnnotationAssertion(rdfs:seeAlso pizza:Pizza <https://en.wikipedia.org/wiki/Pizza>)`)
	// parser.TokenLog = true
	err = o.parseAnnotationAssertion(p)
	if err != nil {
		t.Fatal(err)
	}

	expr = o.AllAnnotationAssertions[1]
	if expr.S != `The-Pizza-Namespace#Pizza` {
		t.Fatal(expr.S)
	}
	if expr.T != `https://en.wikipedia.org/wiki/Pizza` {
		t.Fatal(expr.T)
	}
}

const ontologyTestString = `
Ontology(<urn:absolute:test.de>

	Declaration(Class(hello:FishbonePizza))
	Declaration(Class(:AmericanHotPizza))
	Declaration(Class(:AmericanaPizza))
	Declaration(Class(:AnchovyTopping))
	Declaration(Class(:CaloriePizza))
	Declaration(Class(:CaperTopping))
	Declaration(Class(:Cheese))
	Declaration(Class(:CheesePizza))
	Declaration(Class(:CheeseTopping))
	Declaration(Class(:CheesyPizza))
	Declaration(Class(:DeepPanBase))
	Declaration(Class(:Food))
	Declaration(Class(:GreenPepperTopping))
	Declaration(Class(:HamTopping))
	Declaration(Class(:HighCaloriePizza))
	Declaration(Class(:Hot))
	Declaration(Class(:InterestingPizza))
	Declaration(Class(:JalapenoPepperTopping))
	Declaration(Class(:LowCaloriePizza))
	Declaration(Class(:MargheritaPizza))
	Declaration(Class(:MeatTopping))
	Declaration(Class(:Medium))
	Declaration(Class(:Mild))
	Declaration(Class(:Mozzarella))
	Declaration(Class(:MozzarellaTopping))
	Declaration(Class(:MushroomTopping))
	Declaration(Class(:NamedPizza))
	Declaration(Class(:NonVegetarianPizza))
	Declaration(Class(:OliveTopping))
	Declaration(Class(:OnionTopping))
	Declaration(Class(:Parmesan))
	Declaration(Class(:ParmesanTopping))
	Declaration(Class(:PepperTopping))
	Declaration(Class(:PepperoniTopping))
	Declaration(Class(:Pizza))
	Declaration(Class(:PizzaBase))
	Declaration(Class(:PrawnTopping))
	Declaration(Class(:RedPepperTopping))
	Declaration(Class(:SalamiTopping))
	Declaration(Class(:SeafoodTopping))
	Declaration(Class(:SohoPizza))
	Declaration(Class(:SpicyBeefTopping))
	Declaration(Class(:SpicyPizza))
	Declaration(Class(:SpicynessValuePartition))
	Declaration(Class(:ThinAndCrispyBase))
	Declaration(Class(:TomatoTopping))
	Declaration(Class(:Topping))
	Declaration(Class(:TunaTopping))
	Declaration(Class(:VegetableTopping))
	Declaration(Class(:VegetarianPizza))
	Declaration(ObjectProperty(:hasBase))
	Declaration(ObjectProperty(:hasIngredient))
	Declaration(ObjectProperty(:hasSpicyness))
	Declaration(ObjectProperty(:hasTopping))
	Declaration(ObjectProperty(:isBaseOf))
	Declaration(ObjectProperty(:isIngredientOf))
	Declaration(ObjectProperty(:isToppingOf))
	Declaration(DataProperty(:hasCaloricContentValue))
	Declaration(NamedIndividual(:MyKäseEiPizza))
	Declaration(NamedIndividual(:MyMargherita))
	Declaration(NamedIndividual(:MyQuattroFormaggio))
	Declaration(NamedIndividual(:MySauerkrautpizza))
	
	############################
	#   Object Properties
	############################
	
	# Object Property: :hasBase (:hasBase)
	
	SubObjectPropertyOf(:hasBase :hasIngredient)
	InverseObjectProperties(:hasBase :isBaseOf)
	FunctionalObjectProperty(:hasBase)
	ObjectPropertyDomain(:hasBase :Pizza)
	ObjectPropertyRange(:hasBase :PizzaBase)
	
	# Object Property: :hasIngredient (:hasIngredient)
	
	InverseObjectProperties(:hasIngredient :isIngredientOf)
	TransitiveObjectProperty(:hasIngredient)
	
	# Object Property: :hasSpicyness (:hasSpicyness)
	
	FunctionalObjectProperty(:hasSpicyness)
	ObjectPropertyRange(:hasSpicyness :SpicynessValuePartition)
	
	# Object Property: :hasTopping (:hasTopping)
	
	SubObjectPropertyOf(:hasTopping :hasIngredient)
	InverseObjectProperties(:hasTopping :isToppingOf)
	ObjectPropertyDomain(:hasTopping :Pizza)
	ObjectPropertyRange(:hasTopping :Topping)
	
	# Object Property: :isBaseOf (:isBaseOf)
	
	IrreflexiveObjectProperty(:isBaseOf)
	
	# Object Property: :isToppingOf (:isToppingOf)
	
	IrreflexiveObjectProperty(:isToppingOf)
	
	
	############################
	#   Data Properties
	############################
	
	# Data Property: :hasCaloricContentValue (:hasCaloricContentValue)
	
	FunctionalDataProperty(:hasCaloricContentValue)
	DataPropertyDomain(:hasCaloricContentValue :Food)
	DataPropertyRange(:hasCaloricContentValue xsd:integer)
	
	
	
	############################
	#   Classes
	############################
	
	# Class: hello:FishbonePizza (hello:FishbonePizza)
	
	SubClassOf(hello:FishbonePizza :NamedPizza)
	SubClassOf(hello:FishbonePizza ObjectSomeValuesFrom(:hasTopping :TomatoTopping))
	
	# Class: :AmericanHotPizza (:AmericanHotPizza)
	
	SubClassOf(:AmericanHotPizza :CheesyPizza)
	SubClassOf(:AmericanHotPizza :NamedPizza)
	SubClassOf(:AmericanHotPizza ObjectSomeValuesFrom(:hasTopping :JalapenoPepperTopping))
	SubClassOf(:AmericanHotPizza ObjectSomeValuesFrom(:hasTopping :ParmesanTopping))
	SubClassOf(:AmericanHotPizza ObjectSomeValuesFrom(:hasTopping :PepperoniTopping))
	SubClassOf(:AmericanHotPizza ObjectSomeValuesFrom(:hasTopping :TomatoTopping))
	
	# Class: :AmericanaPizza (:AmericanaPizza)
	
	SubClassOf(:AmericanaPizza :CheesyPizza)
	SubClassOf(:AmericanaPizza :NamedPizza)
	SubClassOf(:AmericanaPizza ObjectSomeValuesFrom(:hasTopping :ParmesanTopping))
	SubClassOf(:AmericanaPizza ObjectSomeValuesFrom(:hasTopping :PepperoniTopping))
	SubClassOf(:AmericanaPizza ObjectSomeValuesFrom(:hasTopping :TomatoTopping))
	SubClassOf(:AmericanaPizza ObjectAllValuesFrom(:hasTopping ObjectUnionOf(:ParmesanTopping :PepperoniTopping :TomatoTopping)))
	
	# Class: :AnchovyTopping (:AnchovyTopping)
	
	SubClassOf(:AnchovyTopping :SeafoodTopping)
	
	# Class: :CaloriePizza (:CaloriePizza)
	
	EquivalentClasses(:CaloriePizza ObjectIntersectionOf(:Pizza DataHasValue(:hasCaloricContentValue "150"^^xsd:int)))
	
	# Class: :CaperTopping (:CaperTopping)
	
	SubClassOf(:CaperTopping :VegetableTopping)
	
	# Class: :Cheese (:Cheese)
	
	SubClassOf(:Cheese :Food)
	
	# Class: :CheesePizza (:CheesePizza)
	
	SubClassOf(:CheesePizza :NamedPizza)
	SubClassOf(:CheesePizza ObjectExactCardinality(4 :hasTopping :CheeseTopping))
	
	# Class: :CheeseTopping (:CheeseTopping)
	
	EquivalentClasses(:CheeseTopping ObjectIntersectionOf(:Topping ObjectSomeValuesFrom(:hasIngredient :Cheese)))
	
	# Class: :CheesyPizza (:CheesyPizza)
	
	EquivalentClasses(:CheesyPizza ObjectIntersectionOf(:Pizza ObjectSomeValuesFrom(:hasTopping :CheeseTopping)))
	
	# Class: :DeepPanBase (:DeepPanBase)
	
	SubClassOf(:DeepPanBase :PizzaBase)
	DisjointClasses(:DeepPanBase :ThinAndCrispyBase)
	
	# Class: :GreenPepperTopping (:GreenPepperTopping)
	
	SubClassOf(:GreenPepperTopping :PepperTopping)
	SubClassOf(:GreenPepperTopping ObjectSomeValuesFrom(:hasSpicyness :Hot))
	
	# Class: :HamTopping (:HamTopping)
	
	SubClassOf(:HamTopping :MeatTopping)
	
	# Class: :HighCaloriePizza (:HighCaloriePizza)
	
	EquivalentClasses(:HighCaloriePizza ObjectIntersectionOf(:Pizza DataSomeValuesFrom(:hasCaloricContentValue DatatypeRestriction(xsd:integer xsd:minInclusive "400"^^xsd:integer))))
	DisjointClasses(:HighCaloriePizza :LowCaloriePizza)
	
	# Class: :Hot (:Hot)
	
	SubClassOf(:Hot :SpicynessValuePartition)
	
	# Class: :InterestingPizza (:InterestingPizza)
	
	EquivalentClasses(:InterestingPizza ObjectIntersectionOf(:Pizza ObjectMinCardinality(3 :hasTopping)))
	
	# Class: :JalapenoPepperTopping (:JalapenoPepperTopping)
	
	SubClassOf(:JalapenoPepperTopping :PepperTopping)
	SubClassOf(:JalapenoPepperTopping ObjectSomeValuesFrom(:hasSpicyness :Hot))
	
	# Class: :LowCaloriePizza (:LowCaloriePizza)
	
	EquivalentClasses(:LowCaloriePizza ObjectIntersectionOf(:Pizza DataSomeValuesFrom(:hasCaloricContentValue DatatypeRestriction(xsd:integer xsd:maxInclusive "402"^^xsd:integer))))
	
	# Class: :MargheritaPizza (:MargheritaPizza)
	
	AnnotationAssertion(rdfs:comment :MargheritaPizza "Pizza aus Tomate und Mozzarella"^^xsd:string)
	SubClassOf(:MargheritaPizza :Food)
	SubClassOf(:MargheritaPizza :NamedPizza)
	SubClassOf(:MargheritaPizza ObjectSomeValuesFrom(:hasTopping :MozzarellaTopping))
	SubClassOf(:MargheritaPizza ObjectSomeValuesFrom(:hasTopping :TomatoTopping))
	SubClassOf(:MargheritaPizza ObjectAllValuesFrom(:hasTopping ObjectUnionOf(:MozzarellaTopping :TomatoTopping)))
	
	# Class: :MeatTopping (:MeatTopping)
	
	SubClassOf(:MeatTopping :Topping)
	
	# Class: :Medium (:Medium)
	
	SubClassOf(:Medium :SpicynessValuePartition)
	
	# Class: :Mild (:Mild)
	
	SubClassOf(:Mild :SpicynessValuePartition)
	
	# Class: :Mozzarella (:Mozzarella)
	
	SubClassOf(:Mozzarella :Cheese)
	DisjointClasses(:Mozzarella :Parmesan)
	
	# Class: :MozzarellaTopping (:MozzarellaTopping)
	
	SubClassOf(:MozzarellaTopping :Topping)
	SubClassOf(:MozzarellaTopping ObjectSomeValuesFrom(:hasIngredient :Mozzarella))
	
	# Class: :MushroomTopping (:MushroomTopping)
	
	SubClassOf(:MushroomTopping :VegetableTopping)
	
	# Class: :NamedPizza (:NamedPizza)
	
	SubClassOf(:NamedPizza :Pizza)
	
	# Class: :NonVegetarianPizza (:NonVegetarianPizza)
	
	EquivalentClasses(:NonVegetarianPizza ObjectIntersectionOf(:Pizza ObjectComplementOf(:VegetarianPizza)))
	DisjointClasses(:NonVegetarianPizza :VegetarianPizza)
	
	# Class: :OliveTopping (:OliveTopping)
	
	SubClassOf(:OliveTopping :VegetableTopping)
	
	# Class: :OnionTopping (:OnionTopping)
	
	SubClassOf(:OnionTopping :VegetableTopping)
	
	# Class: :Parmesan (:Parmesan)
	
	SubClassOf(:Parmesan :Cheese)
	
	# Class: :ParmesanTopping (:ParmesanTopping)
	
	SubClassOf(:ParmesanTopping :CheeseTopping)
	SubClassOf(:ParmesanTopping ObjectSomeValuesFrom(:hasIngredient :Parmesan))
	
	# Class: :PepperTopping (:PepperTopping)
	
	SubClassOf(:PepperTopping :VegetableTopping)
	
	# Class: :PepperoniTopping (:PepperoniTopping)
	
	SubClassOf(:PepperoniTopping :MeatTopping)
	
	# Class: :Pizza (:Pizza)
	
	SubClassOf(:Pizza :Food)
	SubClassOf(:Pizza ObjectExactCardinality(1 :hasBase :PizzaBase))
	SubClassOf(:Pizza DataSomeValuesFrom(:hasCaloricContentValue xsd:integer))
	
	# Class: :PizzaBase (:PizzaBase)
	
	SubClassOf(:PizzaBase :Food)
	
	# Class: :PrawnTopping (:PrawnTopping)
	
	SubClassOf(:PrawnTopping :SeafoodTopping)
	
	# Class: :RedPepperTopping (:RedPepperTopping)
	
	SubClassOf(:RedPepperTopping :PepperTopping)
	
	# Class: :SalamiTopping (:SalamiTopping)
	
	SubClassOf(:SalamiTopping :MeatTopping)
	
	# Class: :SeafoodTopping (:SeafoodTopping)
	
	SubClassOf(:SeafoodTopping :Topping)
	
	# Class: :SohoPizza (:SohoPizza)

	AnnotationAssertion(rdfs:comment :SohoPizza "Pizza aus Tomate und Mozzarella"^^xsd:string)
	SubClassOf(:SohoPizza :Food)
	SubClassOf(:SohoPizza :NamedPizza)
	SubClassOf(:SohoPizza ObjectSomeValuesFrom(:hasTopping :MozzarellaTopping))
	SubClassOf(:SohoPizza ObjectSomeValuesFrom(:hasTopping :OliveTopping))
	SubClassOf(:SohoPizza ObjectSomeValuesFrom(:hasTopping :ParmesanTopping))
	SubClassOf(:SohoPizza ObjectSomeValuesFrom(:hasTopping :TomatoTopping))
	SubClassOf(:SohoPizza ObjectAllValuesFrom(:hasTopping ObjectUnionOf(:MozzarellaTopping :OliveTopping :ParmesanTopping :TomatoTopping)))
	
	# Class: :SpicyBeefTopping (:SpicyBeefTopping)
	
	SubClassOf(:SpicyBeefTopping :MeatTopping)
	
	# Class: :SpicyPizza (:SpicyPizza)
	
	EquivalentClasses(:SpicyPizza ObjectIntersectionOf(:Pizza ObjectSomeValuesFrom(:hasTopping ObjectSomeValuesFrom(:hasSpicyness :Hot))))
	
	# Class: :SpicynessValuePartition (:SpicynessValuePartition)
	
	EquivalentClasses(:SpicynessValuePartition ObjectUnionOf(:Hot :Medium :Mild))
	
	# Class: :ThinAndCrispyBase (:ThinAndCrispyBase)
	
	SubClassOf(:ThinAndCrispyBase :PizzaBase)
	
	# Class: :TomatoTopping (:TomatoTopping)
	
	SubClassOf(:TomatoTopping :VegetableTopping)
	
	# Class: :Topping (:Topping)
	
	SubClassOf(:Topping :Food)
	
	# Class: :TunaTopping (:TunaTopping)
	
	SubClassOf(:TunaTopping :SeafoodTopping)
	
	# Class: :VegetableTopping (:VegetableTopping)
	
	SubClassOf(:VegetableTopping :Topping)
	
	# Class: :VegetarianPizza (:VegetarianPizza)
	
	EquivalentClasses(:VegetarianPizza ObjectIntersectionOf(:Pizza ObjectAllValuesFrom(:hasTopping ObjectUnionOf(:CheeseTopping :VegetableTopping))))
	
	
	############################
	#   Named Individuals
	############################
	
	# Individual: :MeineKäseEiPizza (:MeineKäseEiPizza)
	
	ClassAssertion(:Pizza :MeineKäseEiPizza)
	DataPropertyAssertion(:hasCaloricContentValue :MeineKäseEiPizza "398"^^xsd:integer)
	
	# Individual: :MeineMargherita (:MeineMargherita)
	
	ClassAssertion(:MargheritaPizza :MeineMargherita)
	ClassAssertion(:VegetarianPizza :MeineMargherita)
	DataPropertyAssertion(:hasCaloricContentValue :MeineMargherita "263"^^xsd:integer)
	
	# Individual: :MeineQuattroFormaggio (:MeineQuattroFormaggio)
	
	ClassAssertion(:CheesePizza :MeineQuattroFormaggio)
	DataPropertyAssertion(:hasCaloricContentValue :MeineQuattroFormaggio "723"^^xsd:integer)
	
	# Individual: :MeineSauerkrautpizza (:MeineSauerkrautpizza)
	
	ClassAssertion(:Pizza :MeineSauerkrautpizza)
	DataPropertyAssertion(:hasCaloricContentValue :MeineSauerkrautpizza "10000"^^xsd:integer)
	
	
	DisjointClasses(:AmericanHotPizza :AmericanaPizza :Cheese :MargheritaPizza :PizzaBase :SohoPizza :Topping)
	DisjointClasses(:AmericanHotPizza :AmericanaPizza :CheesePizza :MargheritaPizza :SohoPizza)
	DisjointClasses(:AnchovyTopping :PrawnTopping :TunaTopping)
	DisjointClasses(:CaperTopping :MushroomTopping :OliveTopping :OnionTopping :PepperTopping :TomatoTopping)
	DisjointClasses(:CheeseTopping :MeatTopping :SeafoodTopping :VegetableTopping)
	DisjointClasses(:GreenPepperTopping :JalapenoPepperTopping :RedPepperTopping)
	DisjointClasses(:HamTopping :PepperoniTopping :SalamiTopping :SpicyBeefTopping)
	DisjointClasses(:Hot :Medium :Mild)		
)
`
