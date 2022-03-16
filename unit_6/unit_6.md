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

## NIL 
- Do something about nothing (NIL)
- Understand the trouble with nil 
- See how Go improves on nil 

`nil` is a zero value.  A pointer with nowhere to point has the value `nil`.  `nil` identifier is the zero value for slices, maps and interfaces.
nil is known as NULL, null or None in other languages.

#### Summary 
- Nil pointer dereferences will crash program 
- Methods can guard against receiving nil values 
- Default behaviour can be provided for functions pass as args 
- A nil slice often interchangeable but not the same as empty slice 
- Nil map can be read but not written 
- Interface must have type and value nil to be considered nil 

### nil causes panic 
If a pointer isn't pointing anywhere, attempting to dereference will cause a panic and the program will crash.
```go
	var nowhere *int
	fmt.Println(nowhere)
	fmt.Println(*nowhere) // Causes a panic
```

Easy to avoid by checking for nil pointer.
```go
	var nowhere *int
	fmt.Println(nowhere)
	if nowhere != nil {
		fmt.Println(*nowhere) // Not reached if it would cause a panic
	}
```

### Guarding Methods 
Since methods frequentyl receive pointers we know it could be nil.
Go will happily call methods even when the receiver has a nil value.  This means we can use a similar tactic to above to guard against the nil value.
```go 
func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}
```

### Nil function values 
When a variable is declared as a function type its value is nil by default.  
```go
var fn func(a, b int) int 
fmt.Println(fn == nil) // Prints true
```

It's possible to check whether a function value is nil and provide default behaviour

### nil slice 
Nil and empty slice aren't equivalent.  When accepting a slice ensure that nil and empty do have same behaviour.
Safe functions to use are `len`, `cap`, `append`, and `range`.  Directly accessing an element of an empty or nil slice will cause a panic.

### nil maps 
As with slices, a map declared without a composite literal or make build-in has a vlue of nil.  Maps can be read when nil but can not be written to (will cause panic on write).

### nil Interfaces 
When variable is declared to be an interface type without assignment, zero value is nil.
```go
var v interface{} 
fmt.Printf("%T %v %v\n, v, v ,v == nil) //Print <nil> <nil> true 
```
Both interface type and value must be nil for the variable to equal nil.
```go
fmt.Printf("%#v\n", v)
```

### Alternative to nil 
Use a small structure for holding an unset value.
```go 
type number struct {
	value int
	valid bool 
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
if !n.valid {
	return "not set"
}
return fmt.Sprintf("%d", n.value)
```

## Err 
- Write files and handle failure 
- Handle errors with creativity
- Make and identify specific errors 
- Not panic

#### Summary 
- Errors are values that interoperate with multiple return values and rest of Golang 
- Lots of flexibility if you get creative 
- Custom error types are possible by satisfying `error` interface 
- `defer` cleans up before a function returns 
- Type assertions can convert interface to a concrete type or another interface 
- Avoid use of `panic` and return an `error `instead.


### Handling errors
Go has multiple return values allowing a function to return `err` as one of those.
```go
	files, err := ioutil.ReadDir("..")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
```

### Adding elegance to err 
One way to reduce amount of error handling code is to isolate error-free subset of a program from an inherently error-prone one (ex. IO).

#### Writing error handling
It is good practice to consistently indent errors making it easier to scan through the code without reading all the error handling.

Don't try to exit the program if we can handle it gracefully, functions should return errors instead of exiting in case the calling program wants to retry.

### `defer`
TO ensure that an action takes place before a containing function returns we can use the defer statement.
Any method or function can be deferred and it's not specifically for error handling though it does remove that burden by handling clean up.

### Adding creativity to error handling 
This code snippet shows how to can be smart and store the `err` in a structure and handle the error on consecutive calls to this writer.  By doing this we apply the error handling to the structure and not to each individual line.
```go
type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // Don't write if the last time we tried something it err'd
	}

	_, sw.err = fmt.Fprintln(sw.w, s) //Write a line and store the error in the struct
}
```

### New errors 
New errors occur when something goes wrong and we need to tell the caller about the problem.

`errors` package contains a constructor function that accepts a string for an error message.
Make sure that error messages are useful.  They may be useful to users or other developers trying to use the program.

```go 
(g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("Out of bounds")
	}
	g[row][column] = digit
	return nil

}
```

### Errors in packages 
We can create custom error variables to include and return the variable instead of having to create a new error every time

```go 
var (
	ErrBounds = errors.New("Out of bounds")
	ErrDigit  = errors.New("Invalid digit")
)

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	g[row][column] = digit
	return nil

}
```

### Custom errors types 
`error` type is an interace meaning that any type implementing `Error()` will satisfy the interface.

```go 
type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}
```

#### Type assertions 
To access individual errors we can do so with type assertions.
Below we assert that `err` is of type `SudokuError` and then we can get the struct of type `SodukoError` and examine the slice that the type is.

### Panic
Go doesn't have `exceptions` like other languages but it does have `panic`.  When a panic occurs the program will crash (as with unhandled exceptions).

Exceptions tend to be opt-in whereas Error values in go are simple and are readily available.

Benefits: 
- Pushes developers to consider errors 
- Errors don't require specialized keywords making them simpler and more flexible

Panic should be rare but it should be noted that `panic` will run any deferred functions where `exit` will not.

To keep panic from crashing the program there is the `recover` function.

If a deferred function calls recover, the panic will stop and the program continues running.

```go 
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovery....!")
			fmt.Println(e)
		}
	}()

	panic("I forgot my towel")
```
