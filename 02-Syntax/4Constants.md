# Constants
* Only exist at compilation
* not variables
* Constants can be 'of kind' OR 'of type'
* Constants can only be based on built-in types (since they only exist at compile time)
```go
// Untyped Constants ('of kind')
const ki = 12345 // kind: integer
const kf = 3.24 // kind: floating-point

// Type Constants have restricted type
const ti int = 12345 // type: int
const tf float64 = 3.14 // type: float64
```
* Constants 'of kind' can be implicitly converted by compiler, constants 'of type' cannot be implicitly converted
    - not limited by precision

* When it comes to compiler, constants are like high precisioin calculator
    - consts of kind are automatically at least 256 bits of precision
    - The max int (int64) that will compile: 9223372036854775807
    - Can get around this by using constant of kind
    ```go
    // Will compile
    const big = 9223372036854775807987654321
    
    // Will not compile
    const bigInt int64 = 223372036854775807987654321


## Kind Promotion
    - tells user how things promote
    - floats promote over ints
    - types promote over kind
### Answer
```go 
var answer = 3 * 0.333
```
- KindFloat(3) * KindFloat(0.333) 
- These are constants of a kind 
- kind Int will be 'promoted up' to be kind Float

### Third
```go
const third = 1 / 3.0 
```
- Integer promotes up to be KindFloat

### Zero
```go
zero = 1 / 3
```
- No Kind Promotion here
- Constant zero will be of kind integer 

### Promote from Kind to Type
```go
const one int8 = 1
const two = 2 * one // int8(2) * int8(1)
```
2 is originally a KindInt, but it gets promoted up to an int8

## Duration  
```go
import "time"
import "fmt"
// From time package
type Duration int64
const (
    Nanosecond Duration = 1
    Microsecond         = 1000 * Nanosecond
    Millisecond         = 1000 * Microsecond
    Second              = 1000 * Millisecond
    Minute              = 60 * Second
    Hour                = 60 * Minute
)

// Add returns the time t+d
func (t Time) Add(d Duration) Time {}

now = time.Now()

// Subtract 5 nanoseconds from now using a literal constant
// -5, an unnamed constant OF KIND, can be passed to function that will accept Duration
// This works
now.Add(-5) 


// Do the same using a declared constant
// This works
const timeout = 5 * time.Second // time.Duration(5) * time.Duration(Second)
const := now.Add(-timeout)

// What DOES NOT work: 
minusFive := int64(-5)
variable := now.Add(minusFive)
// Add only wants Duration, an int64 constant OF TYPE cannot be implicily converted into Duration
``` 
- Second way to declare Types: Duration is based on int64, it's not alias
- How type & kind work together Duration is Type, 1000 & 60 are Kinds (unnamed constant)
- Nanosecond is Type int64, 1000 is KindInteger which is promoted up to int64 to produce Microsecond, which is of Type
- Constants of a kind can be implicitly converted when there's compatibility

## IOTA
```go
const (
    A1 = iota // 0 : Starts at 0
    B1 = iota // 1 : Increment by 1
    C1 = iota // 2 : Increment by 1
)
```
EQUIVALENT TO:
```go
const (
    A2 = iota // 0 : Starts at 0
    B2 // 1 : Increment by 1
    C2 // 2 : Increment by 1
)
```
## Say we want to start at 1
```go
const (
    A3 = iota + 1 // 1: Start at 0 + 1
    B3 // 2 : Increment by 1
    C3 // 3 : Increment by 1
)


