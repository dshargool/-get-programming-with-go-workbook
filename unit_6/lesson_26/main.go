package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y++
}
func (t *turtle) down() {
	t.y--
}
func (t *turtle) right() {
	t.x++
}
func (t *turtle) left() {
	t.x--
}

func main() {
	answer := 42
	fmt.Println(&answer)  //Prints address of the answer variable
	address := &answer    // Address holds the address of answer variable
	fmt.Println(*address) //Prints 42 which is the value of the variable at address
	fmt.Printf("answer is a %T\n", address)

	canada := "Canada"
	var home *string
	fmt.Printf("home is a %T\n", home)
	home = &canada
	fmt.Println(*home)

	var administrator *string
	scolese := "Chris Scolese"
	administrator = &scolese
	fmt.Println(*administrator)
	bolden := "Charles Bolden"
	administrator = &bolden
	fmt.Println(*administrator) // Prints Charles Bolden
	bolden = "Charles Frank Bolden"
	fmt.Println(*administrator) // Prints Charles Frank Bolden
	*administrator = "Charlie Bolden"
	fmt.Println(bolden) //Prints Charlie Bolden

	major := administrator // Points to the same address as administrator
	*major = "Major Charles Bolden"
	fmt.Println(bolden) // Prints Major Charles Bolden

	lightfoot := "Robert Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator == major) // Major points to address bolden.  Adminsitrator is address lightfoot

	charles := *major // Makes a copy of the dereferenced value of major
	*major = "Charles Bolden"
	fmt.Println(charles) // Prints Major Charles Bolden
	fmt.Println(bolden)  // Prints Charles Bolden

	// Prefixing structure with address operator
	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	// Unnecessary to dereference structures to access their fields
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)

	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

	fmt.Println(timmy.age)
	birthday(timmy)
	fmt.Println(timmy.age)

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)
	fmt.Println(rebecca.age)

	timmy.birthday()
	fmt.Println(timmy.age)
	rebecca.birthday() // Don't include the & as Go will automatically determine address when calling methods with dot notation.
	fmt.Println(rebecca.age)

	player := character{name: "Matthias"}
	levelUp(&player.stats) // Character structure doesn't have pointers but can take memory address of any field when need
	fmt.Printf("%+v\n", player.stats)

	// turtle.go
	var turt turtle
	turt.up()
	turt.up()
	turt.up()
	turt.up()
	turt.up()
	turt.down()
	turt.right()
	fmt.Println(turt)

}
