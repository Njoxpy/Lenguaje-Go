## Packages in Go

Every Go program is composed of **packages**.
Execution always begins in the **`main`** package.

This example uses the packages with import paths **`"fmt"`** and **`"math/rand"`**.

By convention, the **package name** matches the **last element** of its import path.
For example, files in the `"math/rand"` package start with the declaration:

```go
package rand
```

### Example:

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}
```

---

## Imports

In Go, imports can be grouped into a **factored import statement**, enclosed in parentheses.
This is the preferred style for readability and organization.

### Example (factored import):

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

**Output:**

```
Now you have 2.6457513110645907 problems.
Program exited.
```

Alternatively, you can use multiple single-line import statements, though it’s less common:

```go
import "fmt"
import "math"
```

---

## Exported Names

In Go, **a name is exported** if it begins with a **capital letter**.

For example:

- `Pizza` — exported
- `Pi` — exported from the `math` package
- `pizza` and `pi` — _not exported_ (private to the package)

When you import a package, you can access only its **exported names**.
Unexported names (those starting with a lowercase letter) are **not accessible** from outside the package.

### Example (with error):

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.pi) // Error: cannot refer to unexported name math.pi
}
```

### Fixed version:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi) // Works correctly
}
```

# **Go Language Basics**

## **Functions**

A **function** in Go can take zero or more arguments.

In the following example, the `add` function takes two parameters of type `int`.
Notice that the type is written **after** the variable name.

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

---

## **Functions (continued)**

When two or more consecutive parameters share a type, you can omit the type from all but the last.
For example, instead of writing:

```go
x int, y int
```

you can write:

```go
x, y int
```

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

---

## **Multiple Results**

A function can return **multiple values**.
Here, the `swap` function returns two strings.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

---

## **Named Return Values**

Go allows you to name return values. These are treated as variables defined at the top of the function.
A **naked return** statement (without arguments) returns these named values automatically.

**Note:** Naked returns should be used only in short, simple functions to preserve readability.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

---

## **Variables**

The `var` statement declares a list of variables, with the **type coming last**.

A `var` declaration can appear at both package and function level.

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

---

## **Variables with Initializers**

A `var` declaration can include **initializers**.
If an initializer is provided, the type can be omitted since it will be inferred from the value.

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

---

## **Short Variable Declarations**

Inside a function, you can use the shorthand `:=` to declare and initialize variables.

Outside functions, `:=` cannot be used because all statements must start with a keyword (like `var` or `func`).

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

---

## **Basic Types**

Go’s basic data types include:

- `bool`
- `string`
- Numeric types:

  - Signed integers: `int`, `int8`, `int16`, `int32`, `int64`
  - Unsigned integers: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
  - Aliases: `byte` (for `uint8`), `rune` (for `int32`, representing Unicode code points)

- Floating point: `float32`, `float64`
- Complex numbers: `complex64`, `complex128`

Typically, `int`, `uint`, and `uintptr` are 32 bits on 32-bit systems and 64 bits on 64-bit systems.
Use `int` unless you have a specific reason to choose another type.

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

---

## **Zero Values**

Variables declared without an explicit initializer are assigned their **zero value**:

- `0` for numeric types
- `false` for booleans
- `""` (empty string) for strings

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

---

## **Type Conversions**

Use `T(v)` to convert the value `v` to type `T`.

Example:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, more simply:

```go
i := 42
f := float64(i)
u := uint(f)
```

Unlike C, Go **requires explicit type conversions**.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```

---

## **Type Inference**

If you declare a variable without specifying its type (using `:=` or `var =`), Go infers the type from the value.

- When the value has a type, the new variable takes that type.
- When the value is an **untyped constant**, Go assigns a suitable default type.

Examples:

```go
var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

```go
package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
```

---

## **Constants**

Constants are declared like variables, but with the `const` keyword.
They can be character, string, boolean, or numeric values.
You **cannot** use `:=` for constants.

```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

---

## **Numeric Constants**

Numeric constants are **high-precision** values.
An untyped constant takes on the type required by its context.

```go
package main

import "fmt"

const (
	// Create a huge number by shifting 1 left 100 places (1 followed by 100 zeroes)
	Big = 1 << 100
	// Shift it right again by 99, resulting in 1<<1 or 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```
