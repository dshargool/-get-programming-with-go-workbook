# State and Behaviour - Object Oriented Design
- variables represent state 
- functions and methods define behaviour 

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
