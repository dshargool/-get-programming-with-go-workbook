# First Class Functions
Functions are 'first-class' meaning they work in all places that other types work.  Functions can return functions even!

## Assign function to variable
```go
my_func_var := bigfunc // Assign function to var (not call as no ())
fmt.Println(my_func_var()) // Output the returned value of the function
```

Doing this allows us to further segment our code into cleaner pieces.
```go
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vº K\n", k)
		time.Sleep(time.Second)
	}
}
```
## Declaring function types
Can declare new type for a function to condence/clarify code that refers to it.
```go
type sensor func() kelvin
```
Then we can declare a function that takes a sensor as:
```go
func measureTemperature(samples int, s sensor)
```

## Function literals
Also known as anonymous functions is a function without a name.  This is useful when we want to create a function on the fly; such as returning a function.
```go
func main() {
	func() {
	fmt.Println("function anon")
	}
}
```

### Closures
Anonymous function references the variables accepted as parameters.  Even after the function returns the variables captured by the closure survive so they are still accessible.
The anonymous function 'encloses' the variables in scope hence the term 'closure'.
