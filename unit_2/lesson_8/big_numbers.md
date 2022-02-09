# Big Numbers
- Specify exponents
- `big` package
- `big` constants and literal values


## Big Types
- `big.Int`
  - Can happily store and operate on 24 quintillion km
  - Two ways to set big.Int
    - ```my_var := big.NewInt(3000)```
        when NewInt doesn't overflow uint64
    - ```
	  my_var := new(big.Int)
	  my_var.SetString("3000", 10)
       ```
       when value does overflow uint64
- `big.Float`
- `big.Rat`

## Constants
Constants are untyped and do not need to be defined to use big as they do so in the background at compile time.  This means that calculations resulting in values that don't overflow int can be assigned to regular variables (non-big).  In the event that variables are too large and overflow an error will be thrown.


