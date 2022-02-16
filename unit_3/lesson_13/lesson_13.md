# Methods
Functions that enhance types with additional behaviour.  Same as functions but allow different way to organize them.

```go
type celsius float64
var temperature celsius = 20
fmt.Println(temperature)
```

These types are unique and can not an alias; meaning they can not be combined with other types (for example the float64 that our type is based on.

```go
var warmUp float64 = 10
temperature += warmUp // This is invalid
```

## Using own types
Once declared can use own type anywhere a predeclared (int for example) Go type can be used.

Using own types can help prevent silly mistakes (like passing a celsius instead of a fahrenheit type)

## Adding methods to types
Can add methods to custom types (not predeclared types)
```go
func (receiver) methodname result {
}

func (kelvin k) celsius() celsius {
}
```
Can be called using dot notation `kelvin.celsius()`


