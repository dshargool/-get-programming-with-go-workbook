# Slices
- Use Slices
- Alphabetize slices with stdlib

Slicing is expressed with a half-open range.  For example `planets[0:4]` begins with index 0 and continues up to but NOT INCLUDING index 4.

## Indexing
You can still index slices like you would arrays.
```go
var big_arr [100]int
var slice := big_arr[0:20]
slice[0]  // This is valid
```

## Slicing Slice 
You can also slice slices
```go
var big_arr [100]int
var slice := big_arr[0:20]
var small_slice := slice[0:5]
```

## Modifying Slice
Modifying a slice modifies the underlying array
```go
var big_arr [100]int
var slice := big_arr[0:20]
var small_slice := slice[0:5]
small_slice[0] = 10
fmt.Println(big_arr[0] == 10) // This is TRUE
```

## Default indices when Slicing
Omitting first index defaults to beginning of array
```go
begin_slice := arr[:5]
```
Omitting last index defaults to length of array
```go
end_slice := arr[3:]
```
Omitting both gives whole array
```go
whole_slice := arr[:]
```

### Slicing strings 
Slicing also works on strings the same except for the fact that when the underlying array is changed it won't update the slice.

## Declaring a slice
Slices can be declared directly by giving it the type `[]type` with no value between the brackets.
```go
dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
```
An underlying array of size 5 still exists and then we get this slice view.

## Slice Uses 
### Function Parameters
We can use slices as function parameters since they don't have a length associated with their type.
```go
func my_func(my_param []string) // This is ok and will accept a slice of any length
```

## Slice Methods
By using a custom type we can attach methods to our slice.
```go
type StringSlice []string

func (p StringSlice) Sort()
```

