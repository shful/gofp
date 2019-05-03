package ontologies

import (
	"fmt"
	"testing"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/builtindatatypes"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/facets"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/owlfunctional/properties"
	"github.com/shful/gofp/storedefaults"
)

func helperTestExplicitDecl(ontologyTestString string, explicitDecls bool) (*Ontology,*parser.Parser,*storedefaults.DefaultK) {
	var p *parser.Parser

	k := storedefaults.NewDefaultK()
	rc := StoreConfig{
		AxiomStore: k,
		Decls:      k,
		DeclStore:  k,
	}
	o := NewOntology(
		map[string]string{"": "localprefix#", "hello": "hello.de#", "xsd": builtindatatypes.PRE_XSD, "rdfs": builtindatatypes.PRE_RDFS, "owl": builtindatatypes.PRE_OWL},
		rc,
	)
	o.K = k

	// Explicit mode and ontology has only explicit Decls
	k.ExplicitDecls = explicitDecls
	p = mock.NewTestParser(ontologyTestString)
	parser.TokenLog = false

	return o,p,k
}

func TestParseExplicitDeclMode(t *testing.T) {
	var err error
var o *Ontology
var p *parser.Parser
var k *storedefaults.DefaultK


	// Explicit mode and ontology has no implicit Decls
	o,p,k = helperTestExplicitDecl(ontologyMiniTestString1_ExplicitDecl, true)

	err = o.Parse(p)
	if err != nil {
		t.Fatal(err)
	}

	if len(o.K.AllClassDecls()) != 2 {
		t.Fatal(o.K.AllClassDecls())
	}
	if !k.ClassDeclExists("localprefix#AmericanHotPizza", false) {
		t.Fatal()
	}
	if !k.ClassDeclExists("hello.de#FishbonePizza", false) {
		t.Fatal()
	}
	if len(o.K.AllObjectPropertyDecls()) != 2 {
		t.Fatal(o.K.AllObjectPropertyDecls())
	}


	// Explicit mode and ontology has implicit Decls
	o,p,k = helperTestExplicitDecl(ontologyMiniTestString2_ImplicitDecl, true)

	err = o.Parse(p) //todo add machine readable error types, instead human-readable messages only
	if err == nil {
		t.Fatal("error expected because implicit decl")
	}

	// Implicit mode and ontology has implicit Decls
	o,p,k = helperTestExplicitDecl(ontologyMiniTestString2_ImplicitDecl, false)

	err = o.Parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(o.K.AllClassDecls()) != 6 {
		t.Fatal(o.K.AllClassDecls())
	}

	// also, one ObjectProperty came implicitly:
	if len(o.K.AllObjectPropertyDecls()) != 3 {
		t.Fatal(o.K.AllObjectPropertyDecls())
	}

	// existence of class decls:
	if !k.ClassDeclExists("localprefix#AmericanHotPizza", false) {
		t.Fatal()
	}
	if !k.ClassDeclExists("localprefix#BritishPizza", false) {
		t.Fatal()
	}
	if k.ClassDeclExists("hello.de#NotExistingPizza", false) {
		t.Fatal()
	}
	
	// this Class was declared implicitly. Check that declaration was implicit only:
	if k.ClassDeclExists("spdy://example.com/NewWorldPizza", false) {
		t.Fatal()
	}
	if !k.ClassDeclExists("spdy://example.com/NewWorldPizza", true) {
		t.Fatal()
	}

}

func TestParsePizzaOntology(t *testing.T) {
	var p *parser.Parser
	var err error
	p = mock.NewTestParser(ontologyTestString)

	k := storedefaults.NewDefaultK()
	rc := StoreConfig{
		AxiomStore: k,
		Decls:      k,
		DeclStore:  k,
	}
	o := NewOntology(
		map[string]string{"": "localprefix#", "hello": "hello.de#", "xsd": builtindatatypes.PRE_XSD, "rdfs": builtindatatypes.PRE_RDFS, "owl": builtindatatypes.PRE_OWL},
		rc,
	)

	o.K = k

	err = o.Parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if o.IRI != "<urn:absolute:test.de>" {
		t.Fatal(o.IRI)
	}
	if o.VERSIONIRI != "<http://test.de/1.0.777>" {
		t.Fatal(o.VERSIONIRI)
	}

	// === Decls
	if len(o.K.AllClassDecls()) != 51 {
		t.Fatal(o.K.AllClassDecls())
	}
	if !k.ClassDeclExists("localprefix#AmericanHotPizza", false) {
		t.Fatal()
	}
	if !k.ClassDeclExists("hello.de#FishbonePizza", false) {
		t.Fatal()
	}
	if k.ClassDeclExists("notexisting", false) {
		t.Fatal()
	}

	if len(o.K.AllDataPropertyDecls()) != 2 {
		t.Fatal(o.K.AllDataPropertyDecls())
	}
	if !k.DataPropertyDeclExists("localprefix#hasCaloricContentValue", false) {
		t.Fatal()
	}
	if k.DataPropertyDeclExists("localprefix#hasTopping", false) {
		t.Fatal()
	}
	if k.DataPropertyDeclExists("localprefix#HasTopping", false) { // case differs
		t.Fatal()
	}
	if len(o.K.AllObjectPropertyDecls()) != 7 {
		t.Fatal(o.K.AllObjectPropertyDecls())
	}
	if !k.ObjectPropertyDeclExists("localprefix#hasTopping", false) {
		t.Fatal()
	}

	if len(o.K.AllNamedIndividualDecls()) != 4 {
		t.Fatal(o.K.AllNamedIndividualDecls())
	}
	if !k.NamedIndividualDeclExists("localprefix#MyQuattroFormaggio", false) {
		t.Fatal()
	}

	// === SubDataPropertyOfs
	if len(o.K.AllSubDataPropertyOfs()) != 2 {
		for i, x := range o.K.AllSubDataPropertyOfs() {
			fmt.Println("  ", i, x.P1, x.P2)
		}
		t.Fatal(o.K.AllSubDataPropertyOfs())
	}
	{
		s := o.K.AllSubDataPropertyOfs()[0]
		// we assume an order of the list here - which we know is given, but not guaranteed
		if s.P1.(*declarations.DataPropertyDecl).IRI != "localprefix#hasCaloricContentValue" {
			t.Fatal(s.P1)
		}
		switch s.P2.(type) {
		case *properties.OWLTopDataProperty:
		default:
			t.Fatal(s.P2)
		}
		s = o.K.AllSubDataPropertyOfs()[1]
		// we assume an order of the list here - which we know is given, but not guaranteed
		if s.P1.(*declarations.DataPropertyDecl).IRI != "localprefix#hasSuperhighCaloricContentValue" {
			t.Fatal(s.P1)
		}
	}

	// === SubObjectPropertyOfs
	if len(o.K.AllSubObjectPropertyOfs()) != 2 {
		t.Fatal(o.K.AllSubObjectPropertyOfs())
	}
	{
		s := o.K.AllSubObjectPropertyOfs()[0]
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
	if len(o.K.AllInverseObjectProperties()) != 3 {
		t.Fatal(o.K.AllInverseObjectProperties())
	}
	{
		s := o.K.AllInverseObjectProperties()[0] // the slice preserves the statement order
		x := s.P1.(*declarations.ObjectPropertyDecl)
		if x.IRI != "localprefix#hasBase" {
			t.Fatal(x)
		}
	}

	// ObjectPropertyDomain
	if len(o.K.AllObjectPropertyDomains()) != 2 {
		t.Fatal(o.K.AllObjectPropertyDomains())
	}
	{
		s := o.K.AllObjectPropertyDomains()[0]
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
	if len(o.K.AllObjectPropertyRanges()) != 3 {
		t.Fatal(o.K.AllObjectPropertyRanges())
	}
	{
		s := o.K.AllObjectPropertyRanges()[0]
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
	if len(o.K.AllDataPropertyRanges()) != 1 {
		t.Fatal(o.K.AllDataPropertyRanges())
	}
	{
		s := o.K.AllDataPropertyRanges()[0]
		x := s.R.(*declarations.DataPropertyDecl)
		if x.IRI != "localprefix#hasCaloricContentValue" {
			t.Fatal(x)
		}
		y := s.D.(*facets.BuiltinDatatype)
		if y.DatatypeIRI != builtindatatypes.PRE_XSD+"integer" {
			t.Fatal(y)
		}
	}

	// SubClassOf
	if len(o.K.AllSubClassOfs()) != 67 {
		t.Fatal(len(o.K.AllSubClassOfs()))
	}
	{
		s := o.K.AllSubClassOfs()[0]
		x := s.C1.(*declarations.ClassDecl)
		if x.IRI != "hello.de#FishbonePizza" {
			t.Fatal(x)
		}
	}
}

func TestParseEquivalentClasses(t *testing.T) {
	var p *parser.Parser
	var err error
	k := storedefaults.NewDefaultK()
	rc := StoreConfig{
		AxiomStore: k,
		Decls:      k,
		DeclStore:  k,
	}

	var o *Ontology = NewOntology(
		map[string]string{},
		rc,
	)
	o.K = k

	o.Prefixes[""] = "localprefix#"
	// decls := o.Decls.(*storedefaults.DeclStore)
	decls := o.DeclStore
	decls.StoreClassDecl("localprefix#Pizza")
	decls.StoreClassDecl("localprefix#InterestingPizza")
	decls.StoreObjectPropertyDecl("localprefix#hasTopping")

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

	if len(o.K.AllEquivalentClasses()) != 1 {
		t.Fatal(o.K.AllEquivalentClasses())
	}

	var expr axioms.EquivalentClasses
	expr = o.K.AllEquivalentClasses()[0]
	if len(expr.EquivalentClasses) != 2 {
		t.Fatal(o.K.AllEquivalentClasses())
	}
}

func TestParseAnnotationAssertion(t *testing.T) {

	var p *parser.Parser
	var err error
	var expr annotations.AnnotationAssertion

	k := storedefaults.NewDefaultK()
	rc := StoreConfig{
		AxiomStore: k,
		Decls:      k,
		DeclStore:  k,
	}

	var o *Ontology = NewOntology(
		map[string]string{},
		rc,
	)
	o.K = k

	o.Prefixes[""] = "The local ns/"
	o.Prefixes["xsd"] = "The xsd-ns#"
	o.Prefixes["rdfs"] = "The rdfs-ns#"
	o.Prefixes["pizza"] = "The-Pizza-Namespace"
	decls := o.DeclStore
	decls.StoreClassDecl("The local ns#MargheritaPizza")

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
	if len(o.K.AllAnnotationAssertions()) != 1 {
		t.Fatal(o.K.AllAnnotationAssertions())
	}
	expr = o.K.AllAnnotationAssertions()[0]
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

	expr = o.K.AllAnnotationAssertions()[1]
	if expr.S != `The-Pizza-NamespacePizza` {
		t.Fatal(expr.S)
	}
	if expr.T != `https://en.wikipedia.org/wiki/Pizza` {
		t.Fatal(expr.T)
	}
}

const ontologyTestString = `
Ontology(<urn:absolute:test.de><http://test.de/1.0.777>

	Declaration(Class(hello:FishbonePizza))
	Declaration(Class(<http://hello.com#SomePizzaWithFull-IRIAnd*SpeciälCharß>))
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
	Declaration(DataProperty(:hasSuperhighCaloricContentValue))
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
	SubDataPropertyOf(:hasCaloricContentValue owl:topDataProperty)
	SubDataPropertyOf(:hasSuperhighCaloricContentValue :hasCaloricContentValue)

	
	
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

const ontologyMiniTestString1_ExplicitDecl = `
Ontology(
	<urn:absolute:test.de>
	<http://test.de/1.0.777>

	Declaration(Class(hello:FishbonePizza))
	Declaration(Class(:AmericanHotPizza))
	Declaration(ObjectProperty(:hasBase))
	Declaration(ObjectProperty(:hasIngredient))

	# Object Property: :hasBase (:hasBase)
	SubObjectPropertyOf(:hasBase :hasIngredient)
)
`

// ontologyMiniTestString2_ImplicitDecl declares :hasBase implicitly.
const ontologyMiniTestString2_ImplicitDecl = `
Ontology(
	<urn:absolute:test.de>
	<http://test.de/1.0.777>

	Declaration(Class(hello:FishbonePizza))
	Declaration(Class(:AmericanHotPizza))
	Declaration(ObjectProperty(:hasIngredient))

	# ObjectProperty shortened implicit:
	SubObjectPropertyOf(:hasBase :hasIngredient)

	# ObjectProperty with full IRI implicit:
	SubObjectPropertyOf(<https://example.com/hasAnything> :hasIngredient)

	# and a 3rd Class with full IRI implicit:
	SubClassOf(:AmericanHotPizza <spdy://example.com/NewWorldPizza>)

	# and Classes 4 and 5 shortened implicit:
	SubClassOf(:JamAndKidneyPizza :BritishPizza)

	# and an already implicit declared Class 6 explicit now:
	Declaration(Class(:BritishPizza))
)
`

