package designpattern

import "fmt"

type IBlacksmith interface {
	makeSword() ISword
}

type ISword interface {
	setSheath(sheath string)
	setCurve(curveness int)
	setHolder(holder string)
	status()
}

type modernBlacksmith struct{}

func (mb *modernBlacksmith) makeSword() ISword {
	return new(claymore)
}

type sword struct {
	sheath    string
	curveness int
	holder    string
}

type claymore struct {
	sword
}

func (c *claymore) setSheath(sheath string) {
	c.sheath = sheath
}

func (c *claymore) setCurve(curveness int) {
	c.curveness = curveness
}

func (c *claymore) setHolder(holder string) {
	c.holder = holder
}

func (c *claymore) status() {
	fmt.Printf("Claymore status: \nsheath: %s\ncurveness: %d\nholder: %s\n-------------------------\n", c.sheath, c.curveness, c.holder)
}

type classicBlacksmith struct{}

func (cb *classicBlacksmith) makeSword() ISword {
	return new(katana)
}

type katana struct {
	sword
}

func (k *katana) setSheath(sheath string) {
	k.sheath = sheath
}

func (k *katana) setCurve(curveness int) {
	k.curveness = curveness
}

func (k *katana) setHolder(holder string) {
	k.holder = holder
}

func (k *katana) status() {
	fmt.Printf("Katana status: \nsheath: %s\ncurveness: %d\nholder: %s\n-------------------------\n", k.sheath, k.curveness, k.holder)
}

func getBlacksmith(kind string) IBlacksmith{
	if kind == "modern"{
		return &modernBlacksmith{}
	}
	return &classicBlacksmith{}
}

func AbstractFactory() {
	modernBs := getBlacksmith("modern")
	classicBs := getBlacksmith("classic")

	claymore := modernBs.makeSword()
	katana := classicBs.makeSword()

	claymore.setCurve(15)
	claymore.setHolder("iron")
	claymore.setSheath("leather")

	katana.setCurve(35)
	katana.setHolder("wood")
	katana.setSheath("bamboo")

	printSwordDetail(claymore)
	printSwordDetail(katana)
}

func printSwordDetail(sword ISword){
	sword.status()
}