package main

import "fmt"

type Animal struct {
	Name     string
	Gender   string
	IsMammal bool
}

func (a *Animal) Speak() {
	text := fmt.Sprintf("I am %s. My gender is %s, I am not a mammal.\n", a.Name, a.Gender)
	if a.IsMammal {
		text = fmt.Sprintf("A am %s. My gender is %s, I am a mammal.\n", a.Name, a.Gender)
	}
	fmt.Println(text)
}

type Cat struct {
	Animal
	Breed string
}

func (c *Cat) Speak() {
	c.Animal.Speak()
	fmt.Printf("My breed is %s.\n", c.Breed)
}

type Whale struct {
	Animal
}

func (w *Whale) Speak() {
	fmt.Printf("Woooommmm\n")
}

func polymorphism() {
	fmt.Println("====Polymorphism====")
	animalDiona := Animal{
		Name:     "Diona",
		Gender:   "Female",
		IsMammal: true,
	}
	animalDiona.Speak()

	catKecing := Cat{
		Animal: Animal{
			Name:     "Kecing",
			Gender:   "Female",
			IsMammal: true,
		},
		Breed: "Anggora",
	}
	catKecing.Speak()

	whaleTartaglia := Whale{
		Animal: Animal{
			Name:     "Tartaglia",
			Gender:   "Male",
			IsMammal: true,
		},
	}
	whaleTartaglia.Speak()
}
