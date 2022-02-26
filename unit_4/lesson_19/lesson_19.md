# Maps
- Use maps as collections for unstructured data
- Declare, access and iterate over maps
- Explore some uses of the versatile map type

## Summary
- Maps are versatile collections for unstructured data
- Composite literals provide convenient means to initialize maps
- `range` can iterate over maps
- Maps share same underlying data when assigned or passed to functions
- Collections can be combined with each other

## Declaration
```go
map[string]int
    ^ key   ^ value
```

```go
temperature := map[string]int{
   "Earth": 15,
   "Mars": 65,
}

// Change a value in the map
temperature["Earth"] = 16

// Add a vlue to the map
temperature["Venus"] = 464
```

Go provides the *comma, ok* syntax which can use to distinguish between key existing in the map vs. being not present.  In this example if the key is present ok is `true` otherwise it is `false`.
```go
if moon, ok := temperature["Moon"]; ok {
	fmt.Printf("On average the moon is %v C\n", moon)
} else {
	fmt.Println("Couldn't find the moon")
}
```

## Maps aren't copied
When a map is assigned to a new variable or passed to a function it is not copied and any changes to the map will affect the underlying data.

```go
planets := map[string]string {
	"Earth": "Earthling",
	"Mars": "Martian",
}

planets_2 := planets // These are both the same data
planets_2["Mars"] = "Marstian" // Changes both planets_2 and planets
fmt.Println(planets)
fmt.Println(planets_2)
```

This stands for deleting and adding elements to a map.

## Preallocation
Can preallocate a map with `make`.
```go
temperature := make(map[float64]int, 16)
```

## Iterating with maps
If *frequency* were a slice the keys would need to be integers and the underlying array would need to reserve space to count temperatures that never actually occured.

```go
temperatures := []float64{
	-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
}
frequency := make(map[float64]int)
for _, t := range temperatures {
	frequency[t]++
}
for t, num := range frequency {
	fmt.Printf("%+.2f occurs %d times\n", t, num)
}
```
When using `range` with maps we get back `key` and `value`.
Go does not guarantee the order of the keys so it may change between executions.

## Grouping with maps and slices
We can use `math.Trunc` to round values and assign them into a group.
```go 
	groups := make(map[float64][]float64) // A map with float64 keys and []float64 values

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
```
## Using maps as sets
A *set* is similar to an array except each element is guaranteed to only occur once.  We can make something similar with a map and disregarding the value (only use the key).

```go 
set := make(map[float64]bool)
for _, t := range temperatures {
	set[t] = true
}
```

But since maps are unordered if we want to order out set we need to convert it back to a slice. 

```go 
unique := make([]float64, 0, len(set))

for t := range set {
	unique = append(unique, t)
}

sort.Float64s(unique)
fmt.Println(unique)
```
