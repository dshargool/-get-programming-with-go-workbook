# Bigger Slices
- Append more elements to Slices
- Investigate how length and capacity work

- Slices have length and capacity
- When insufficient capacity, append function will allocate new underlying array
- Can use `make` to pre-allocate slice
- Variadic functions accept multiple arguments, which are placed in a slice

By combining slices (which are views into arrays) and the `append` function we can have variable length arrays.

## The `append` function
`append` is variadic, like `Println` so it can take multiple elements and append them in one go.

```go
dwarfs := []string{"Ceres","Pluto","Haumea","Makemake","Eris"}
dwarfs = append(dwarfs, "Orcus")
```

## length and capacity
Number of elements visible through a slice determine its length.  If a slice has an underlying array that is larger, the slice may still have capacity to grow.

```go 
dwarfs := []string{"One", "Two", "Three", "Four", "Five"}
len(dwarfs) // Outputs 5
cap(dwarfs) // Outputs 5
len(dwarfs[1:2]) // Outputs 1
cap(dwarfs[1:2]) // Outputs 4
```

If the array backing a slice doesn't have enough room to append a new element it will create and copy the array to a new array with twice the capacity.

## Three index slicing
This way we can specify the capacity of our slice.  See below:
```go 
planets := []string{"Mercury","Venus","Earth","Mars","Jupiter","Saturn","Uranus","Neptune"}
terrestrial := planets[0:4:4] // Length 4, Capacity 4
worlds := append(terrestrial, "Ceres")
```
If the third index isn't specified, terrestrial will have an original capacity of 8 (the underlying array) and not create a new copy.  Instead appending to the slice will overwrite values in the underlying array.

Should use three slice array when we don't want to overwrite underlying array (which is probably fairly often!)

## Pre-allocate with `make`
To avoid extra allocations with Go trying to append with insufficient capacity we can use `make` to pre-allocate a slice.
```go 
dwarfs := make([]string, 0, 10)  // Make a slice with length 0 and capacity 10. Append makes length 1
dwarfs = make([]string, 10)	 // Make a slice with length 10 and capacity 10. Append makes length 11
```

## Declaring variadic functions
Functions like `Println` and `append` are variadic because they accept a variable number of arguments.  To declare variadic function, use ellipsis with the last parameter.
```go 
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds)) // Makes a new slice instead of modifying directly
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
```

The worlds parameter is a slice of strings that contains zero or more arguments passed to the function.
To pass a slice instead of multiple arguments, expand the slice with ellipsis when calling
```go 
planets := []string("Venus", "Mars", "Jupiter")
newPlanets := terraform("New", planets...)  // Notice how to pass the slice as multiple arguments
fmt.Println(newPlanets)
