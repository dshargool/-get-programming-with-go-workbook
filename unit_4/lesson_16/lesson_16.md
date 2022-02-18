# Arrays
- Declare and initialize an array 
- Assign and access elements of array 
- Iterate through arrays 

## Declaring an Array 
```go
var var_name [array_size]type
// Can be accessed with where # is the element number starting at 0
var_name[#]
// Default value is initial value for that type so for example:
var int_arr [3]int 
// Initialized as: [0, 0 ,0]
```

If we attempt to access index outside of array compiler may throw an error.  If compiler does not detect the error and it is possible for the error to occur at runtime as a *panic*.

## Initialization
### Composite Literals
A way to initialize an array without doing it one by one. This way we can do it in a single step.
```go
dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
```

Can also use multiple lines to initialize for more readability.  Also don't have to define the length and can let the compile count with `...` in the declaration.
```go
planets := [...]string{
   "Mercury",
   "Venus",
   "Earth",
   "Mars",
   "Jupiter",
   "Saturn",
   "Uranus",
   "Neptune",
   }
```
The trailing comma is required when formatted this way.

## Iteration
### Using len 
```go
for i:=0; i< len(arr); i++ {
	element := arr[i]
}
```
### Using Range
The `range` function returns both the index and value:
```go
for index, value := range arr {
	fmt.Println(i, value
}
```

Note: Can use a blank identifier if we don't need the index variable provided by range
`_, value := range arr`

## Copying Arrays
When an array is assigned to a new variable or passing it to a function a complete copy of the array is made (pass by value).

Arrays of different sizes are considered different types.  Meaning we can't have a parameter in a function of type `[8]string` and pass it a value of `[5]string`.  This can be worked around by using slices instead of arrays.

## Arrays of Arrays
Arrays can have multiple dimensions defined as:
`var grid [8][8]string`
