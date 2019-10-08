package gofp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional"
	"github.com/shful/gofp/owlfunctional/parser"
)

func TestParsePrefixTo(t *testing.T) {
	var err error
	// parser.TokenLog = true
	var p *parser.Parser
	var prefixes map[string]string = map[string]string{}

	p = mock.NewTestParser(`Prefix(terms:=<http://purl.org/dc/terms/>)`)
	err = parsePrefixTo(prefixes, p)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParsePosition1(t *testing.T) {
	var err error
	var o *ontologies.Ontology
	parser.TokenLog = true

	o, err = OntologyFromReader(strings.NewReader(`X`), "Testsource")
	fmt.Println(err, o)
	if err == nil {
		t.Fatal()
	}
	pos := err.(*parser.PErr).AfterPos
	if pos.LineNo1() != 1 {
		t.Fatal(pos)
	}
	if pos.GetCurrentLineHead() != `` {
		t.Fatal("linehead=" + pos.GetCurrentLineHead() + ".")
	}
	if pos.ColNo1() != 1 {
		t.Fatal(pos)
	}
}

func TestParsePosition2(t *testing.T) {
	var err error
	var o *owlfunctional.Ontology

	parser.TokenLog = true
	o, err = OntologyFromReader(strings.NewReader(`
Prefix(:=<urn:absolute:similix.de/similixadmin#>)X
`), "Testsource")
	fmt.Println(err, o)
	if err == nil {
		t.Fatal()
	}
	pos := err.(*parser.PErr).AfterPos
	if pos.LineNo1() != 2 {
		t.Fatal(pos)
	}
	if pos.GetCurrentLineHead() != `Prefix(:=<urn:absolute:similix.de/similixadmin#>)` {
		t.Fatal("linehead=" + pos.GetCurrentLineHead() + ".")
	}
	if pos.ColNo1() != 50 {
		t.Fatal(pos)
	}
}
func TestParsePosition3(t *testing.T) {
	var err error
	var o *owlfunctional.Ontology

	// parser.TokenLog = true

	// Data with unknown prefix in line 144 (counting from 1)
	// and col 109, provided the leading tab counts as 1 column
	o, err = OntologyFromReader(strings.NewReader(`
Prefix(:=<urn:absolute:similix.de/similixadmin#>)
Prefix(hello:=<urn:absolute:similix.de/similixadmin#>)
Prefix(xsd:=<http://www.w3.org/2001/XMLSchema#>)
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
	SubClassOf(:AmericanaPizza ObjectAllValuesFrom(:hasTopping ObjectUnionOf(:ParmesanTopping :PepperoniTopping wrongPrefix:TomatoTopping)))
`),
		"Testsource")
	fmt.Println(err, o.About())
	if err == nil {
		t.Fatal()
	}
	pos := err.(*parser.PErr).AfterPos
	if pos.LineNo1() != 145 {
		t.Fatal(pos, err)
	}
	if pos.GetCurrentLineHead() != `	SubClassOf(:AmericanaPizza ObjectAllValuesFrom(:hasTopping ObjectUnionOf(:ParmesanTopping :PepperoniTopping ` {
		t.Fatal("linehead=" + pos.GetCurrentLineHead() + "<<")
	}
	if pos.ColNo1() != 110 { // count tab as 1 col
		t.Fatal(pos)
	}
}
