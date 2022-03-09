# Unit 6
## Pointers
- Declare and use pointers 
- Understand relationship between pointers and RAM 
- When to use (and not use) pointers

### Summary 
- Pointers store memory addresses 
- Address operator (&) provides memory address of variable 
- Dereference operator (\*) provides value that address points to 
- Pointers are types declared with preceding asterisk `*int` 
- Use pointers to mutate values across function and method bounderies 
- Pointers are most useful with structures and arrays 
- Maps and slices use pointers behind the scenes 
- Interior pointers can point at fields inside structures without declaring them as pointers 
- Don't overuse pointers as they can make code confusing

### Overview 
Pointer is a variable that points to the address of another variable.  A form of *indirection*.

Danging pointers are pointers that point to unallocated/the wrong memory address. 

### Ampersand and Asterisk 
Address Operator: &
- Gives the address of a variable in memory.  
```go 
answer := 42 
fmt.Println(&answer) //Prints address of the answer variable 
```

The reverse operation (from getting the address) is called *dereferencing* which is done with the Asterisk.

```go 
answer := 42 
fmt.Println(&answer) //Prints address of the answer variable 
address := &answer // Address holds the address of answer variable 
fmt.Println(*address) //Prints 42 which is the value of the variable at address
```

Note: Multiplication (\*) requires two values so they can't be confused.

#### Pointer Types 
Pointers store memory addresses.  In the case above the address is of type `*int`

```go 
canada := "Canada" 
var home *string 
fmt.Printf("home is a %T\n", home)
home = &canada 
fmt.Println(*home)
```

The `*string` type pointer can point to any other variable of type string but can't point to other types.

Pointer dereferencing a type is declaring the pointer type.  Pointer dereferencing a variable is used to get the value that variable points to.

#### Using pointers
```go
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
```

#### Pointing to structures 
Unlike strings and numbers composite literals can be prefixed with an address operator.
```go 
	type person struct {
		name, superpower string
		age              int
	}

	// Prefixing structure with address operator
	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	// Unnecessary to dereference structures to access their fields
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
```

#### Pointing to Arrays 
As with structures, composite literals for arrays can be prefixed with the address operator.  Arrays also provide automatic dereferencing (as they do with structures). 
```go
	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
```
NOTE: Composite literals for for slices and maps can also be prefixed with address operator but won't be automatically dereferenced.

### Enabling Mutation 

#### Pointer Parameters
Pointers as parameters.  Function and methods parameters are passed by value in Go.  Meaning functions always operator on a copy of passed arguments.  When a pointer is passed the function receives a copy of the address.  By dereferencing the address a function can mutate the value a pointer points to.

```go 
func birthday(p *person) {
	p.age++
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)
}
```
#### Pointer Receivers 
Method receivers can be used on pointers 
```go
func (p *person) birthday() {
	p.age++
}


	timmy.birthday()
	fmt.Println(timmy.age)
	rebecca.birthday()  // Don't include the & as Go will automatically determine address when calling methods with dot notation.
	fmt.Println(rebecca.age)
```

We don't have to specify the & as Go will automatically determine address when using dot notation.

NOTE: Should use pointer receivers consistently. If methods need pointer receivers, use pointer receivers for all methods of the type.  (Dot notation is the same either way)

#### Interior pointers 
Used to determine the memory address of a field inside of a structure.
```go 
	player := character{name: "Matthias"}
	// &player.stats is an interior pointer
	levelUp(&player.stats)  // Character structure doesn't have pointers but can take memory address of any field when need.
	fmt.Printf("%+v\n", player.stats)
```

#### Changing Arrays 
Slices tend to be preferred but can use arrays when size is fixed.
```go 
func reset(board *[8][8]rune) {
	board[0][0] = 'r'
// ...
}
func main() {
	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c", board[0][0])
}
```

#### Maps are pointers
```go 
func demolist(planets *map[string]string) //Unnecessary pointer
``` 

#### Slices point at arrays 
A slice is represented with three elements, a pointer to an array, the capacity and length of the slice.  Internal pointer allows underlying data to be mutated when a slice is passed directly to a function or method.

Explicit pointers to slices are only useful when modifying slice itself (length, capacity, or starting offset).  Arguably it would be cleaner to return a new slice.
```go 
// Modify the length of the slice
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}
```

#### Pointers and Interfaces
A pointer to a value that satisfies all the interfaces that the non-pointer version satisfies will satisfy the interface.


