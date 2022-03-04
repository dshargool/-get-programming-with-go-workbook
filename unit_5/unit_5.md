# State and Behaviour - Object Oriented Design
- variables represent state 
- functions and methods define behaviour 

### Summary
- Combining Methods and Structure provide much of what other languages provide for classes
- Constructor functions are ordinary functions. 

## Lesson 21 - Structures
- Structures group values together into one unit 
- Structures are values that are copied when assigned or passed to functions 
- Composite literals provide a convenient means to initialize structures 
- Struct tags decorate exportaed fields with additional information that packages can use.
- The `json` package utilizes struct tags to control the output of field names

### Declaring a structure
Could have two independent variables but less error prone and less tedious if we define a structure:
```go
var curiousity struct {
    lat float64
    long float64
}
```
#### As a type
```go 
type location struct {
    lat float64 
    long float64 
}

var curiousity location 
curiousity.lat = 10.2
curiousity.long = 543.2
var spirit location  // Re-use struct as a type
spirit.lat = 12.2
spirit.long = -322.2
```

### Initializing with composite literals
```go
type location struct {
    lat, long float64 
}

opportunity := location{lat:33.3, long: -99.9}
```
You can initialize without using the field names as long as you define them in order.  If a field is not listed it gets the zero value.
```go 
opportunity = location{1.1, 2.2}
```

### Pass by Copy
Structures are copied between each other so changes to each one don't affect the other
```go
obj_one := location{-4.5, 137.4}
obj_two := obj_one
obj_two.lat += 0.111
fmt.Println(obj_one, obj_two)
``` 

### Slices of Structures 
```go
type named_location struct {
	name string
	lat  float64
	long float64
}

locations := []named_location{
	{name: "Brad", lat: 4.44, long: 6.66},
	{name: "Colum", lat: 1.22, long: 99.99},
	{name: "Chall", lat: -1.11, long: -0.22},
}
fmt.Println(locations)
```

### Encoding structures to JSON 
JSON is a subset of JavaScript commonly used for web APIs.  The `Marshall` function from the `json` package returns JSON as bytes which can be sent over the wire or converted to a string for display.
```go
type json_location struct {
	Lat, Long float64 // Must be capitalized for JSON!
}

curiousity_json := json_location{-4.58, 137.44}
bytes, _ := json.Marshal(curiousity_json)
fmt.Println(string(bytes))
```
#### Customizing JSON with struct tags 
`json` package requires fields exported (initial uppercase letter and CamelCase on multiword names).  If using snake_case fields of structure can be tagged with the field names we want the `json` package to use.

```go 
type json_tags_location struct {
	Lat  float64 `json:"latitude"` // Use raw string literal ` so we don't have to escape the quotation marks
	Long float64 `json:"longitude"`
}

curiousity_tags_json := json_tags_location{-4.58, 137.44}
bytes, _ = json.Marshal(curiousity_tags_json)
fmt.Println(string(bytes))
```

## Lesson 22 - No Class
- Write methods providing behaviour to structure data
- Apply principles of OO Design

### Attaching methods to structures
The same way that we attach methods to types of custom `int` or `float64` we can attach to custom structs.
```go 
type coordinate struct {
	d, m, s float64
	h rune
}

//Convert from d/m/s to decimal degrees
func (c coordinate) decimal float64 {
	sign := 1.0
	switch c.h {
	case 'S','W','s','w':
		sign = -1 
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
lat := coordinate{4, 35, 22.2, 'S'}
long := coordinate{137, 26, 30.12, 'E'}

fmt.Println(lat, long)
fmt.Println(lat.decimal(), long.decimal())
```

### Constructors
If you need a composite literal that's anything more than list of values, can write constructor function.
```go 
type location struct {
	lat, long float64
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}
```
Go does not have constructors as a special language feature and just uses ordinary functions with specific naming scheme.  Functions of the form `newType` or `NewType` are used to construct a value of said type.  It can then be used as any other function would be.

If we want to construct from different input types, just create appropriate functions to do so.

### Class alternative
Go doesn't have a `class` as in other languages.  Just structures with methods related to them which accomplishes the same task.

## Lesson 23 - Composition and Forwarding
- Compose structures with composition
- Forward methods to other methods 
- Forget about classical inheritance

### Summary 
- Composition is a technique of breaking large structures down into small structures and putting them together 
- Embedding gives access to the fields of inner structure in the outer structure 
- Methods are automatically forwarded when you embed types in a structure 
- Go will inform of name collisions but only if those methods are used

Go provides a special feature called *embedding* to forward methods when composing structures.

### Composing Structures 
Can have structures within a structure to help simplify it.
```go 
type report struct {
  sol int
  temperature temperature 
  location location 
}

type temperature struct {
  high, low celsius 
}

type location struct {
  lat, long float64 
}

type celsius float64
```

Can further organize code by hanging methods from each type (such as the average of two values).
```go 
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}
```

If we want to expose this function on the parent structure we can just write a function to forward the function.  This provides convenient access.
```go 
func (r report) average() celsius {
	return r.temperature.average()
}
```

### Forwarding Methods
It's inconvenient to add a lot of *boilerplate* code to forward these functions as it's just clutter.  Go gets around this by using *struct embedding*.  This is done by adding the type without a field name to the struct.
```go 
type report struct {
  sol int
  temperature 
  location 
}

fmt.Printf("average temp %+v C\n", report.average())
```

This doesn't only forward methods but also the fields of the inner structure.

### Name collisions
If two embedded types have a field/method of the same name the compiler is smart enough to through an error.  To resolve we simply specify which type to use or we can define the behaviour using a methodfor the struct itself.  This is called an *ambiguous selector* error.
