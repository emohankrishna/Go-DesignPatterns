package main

import "fmt"

// Implementation of AbstractFactoryPattern
type IFurnitureFactory interface {
	makeChair() IChair
	makeSofa() ISofa
	makeTable() ITable
}

func NewFurnitureFactory(s string) IFurnitureFactory {
	if s == "art" {
		return &Art{}
	} else if s == "victoria" {
		return &Victorian{}
	} else if s == "modern" {
		return &Modern{}
	}
	return nil
}

type IChair interface {
	setType(s string)
	getType() string
	setName(s string)
	getName() string
}
type Chair struct {
	Type string
	Name string
}

func (c *Chair) setType(s string) {
	c.Type = s
}
func (c *Chair) getType() string {
	return c.Type
}
func (c *Chair) setName(s string) {
	c.Name = s
}
func (c *Chair) getName() string {
	return c.Name
}

type ISofa interface {
	setType(s string)
	getType() string
	setName(s string)
	getName() string
}
type Sofa struct {
	Type string
	Name string
}

func (s *Sofa) setType(str string) {
	s.Type = str
}
func (s *Sofa) getType() string {
	return s.Type
}
func (s *Sofa) setName(str string) {
	s.Name = str
}
func (s *Sofa) getName() string {
	return s.Name
}

type ITable interface {
	setType(s string)
	getType() string
	setName(s string)
	getName() string
}
type Table struct {
	Type string
	Name string
}

func (t *Table) setType(s string) {
	t.Type = s
}
func (t *Table) getType() string {
	return t.Type
}
func (t *Table) setName(s string) {
	t.Name = s
}
func (t *Table) getName() string {
	return t.Name
}

// Art
type ArtChair struct {
	Chair
}
type ArtSofa struct {
	Sofa
}
type ArtTable struct {
	Table
}

// Victoria
type VictorianChair struct {
	Chair
}
type VictorianSofa struct {
	Sofa
}
type VictorianTable struct {
	Table
}

// Modern
type ModernChair struct {
	Chair
}
type ModernSofa struct {
	Sofa
}
type ModernTable struct {
	Table
}

type Art struct {
}

func (a *Art) makeChair() IChair {
	return &ArtChair{
		Chair: Chair{
			Type: "Art",
			Name: "Chair",
		},
	}
}
func (a *Art) makeSofa() ISofa {
	return &ArtSofa{
		Sofa: Sofa{
			Type: "Art",
			Name: "Sofa",
		},
	}
}
func (a *Art) makeTable() ITable {
	return &ArtTable{
		Table: Table{
			Type: "Art",
			Name: "Table",
		},
	}
}

// Victorian
type Victorian struct {
}

func (a *Victorian) makeChair() IChair {
	return &VictorianChair{
		Chair: Chair{
			Type: "Victorian",
			Name: "Chair",
		},
	}
}
func (a *Victorian) makeSofa() ISofa {
	return &VictorianSofa{
		Sofa: Sofa{
			Type: "Victorian",
			Name: "Sofa",
		},
	}
}
func (a *Victorian) makeTable() ITable {
	return &VictorianTable{
		Table: Table{
			Type: "Victorian",
			Name: "Table",
		},
	}
}

// Modern implementation
type Modern struct {
}

func (a *Modern) makeChair() IChair {
	return &ModernChair{
		Chair: Chair{
			Type: "Modern",
			Name: "Chair",
		},
	}
}
func (a *Modern) makeSofa() ISofa {
	return &ModernSofa{
		Sofa: Sofa{
			Type: "Modern",
			Name: "Sofa",
		},
	}
}
func (a *Modern) makeTable() ITable {
	return &ModernTable{
		Table: Table{
			Type: "Modern",
			Name: "Table",
		},
	}
}

func PrintChair(c IChair) {
	fmt.Printf(" Type of %s : %s \n", c.getName(), c.getType())
}
func PrintSofa(c ISofa) {
	fmt.Printf(" Type of %s : %s \n", c.getName(), c.getType())
}
func PrintTable(c ITable) {
	fmt.Printf(" Type of %s : %s \n", c.getName(), c.getType())
}

func main() {
	artFactory := NewFurnitureFactory("art")
	artChair := artFactory.makeChair()
	artTable := artFactory.makeTable()
	artSofa := artFactory.makeSofa()

	fmt.Println("====== Art Decor Furniture =======")
	PrintChair(artChair)
	PrintSofa(artSofa)
	PrintTable(artTable)

	victorianFactory := NewFurnitureFactory("victoria")
	v_Chair := victorianFactory.makeChair()
	v_Table := victorianFactory.makeTable()
	v_Sofa := victorianFactory.makeSofa()

	fmt.Println("====== Victoria Furniture =======")
	PrintChair(v_Chair)
	PrintSofa(v_Sofa)
	PrintTable(v_Table)

	modernFactory := NewFurnitureFactory("modern")
	m_Chair := modernFactory.makeChair()
	m_Sofa := modernFactory.makeSofa()
	m_Table := modernFactory.makeTable()

	fmt.Println("====== Modern Furniture =======")
	PrintChair(m_Chair)
	PrintSofa(m_Sofa)
	PrintTable(m_Table)

}
