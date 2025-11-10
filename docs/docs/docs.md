### Basic Types in Go

Go provides a set of **basic types** that are commonly used in programs:

- **Boolean:**
  `bool`

- **Strings:**
  `string`

- **Signed integers:**
  `int`, `int8`, `int16`, `int32`, `int64`

- **Unsigned integers:**
  `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`

- **Aliases:**
  `byte` → alias for `uint8`
  `rune` → alias for `int32` (represents a Unicode code point)

- **Floating-point numbers:**
  `float32`, `float64`

- **Complex numbers:**
  `complex64`, `complex128`

---

The following example demonstrates variables of various types. It also shows how variable declarations can be **grouped together** in a block, similar to how imports are grouped.

> **Note:** The types `int`, `uint`, and `uintptr` are typically **32 bits** on 32-bit systems and **64 bits** on 64-bit systems. Use `int` for general-purpose integers unless you have a specific reason to choose a fixed-size or unsigned integer.

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

**Output example:**

```
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
```

### Zero Values in Go

In Go, **variables declared without an explicit initial value** are automatically assigned a **zero value**. The zero value depends on the type of the variable:

- **Numeric types** → `0`
- **Boolean type** → `false`
- **String type** → `""` (empty string)

---

**Example:**

```go
package main

import "fmt"

func main() {
    var i int      // zero value is 0
    var f float64  // zero value is 0
    var b bool     // zero value is false
    var s string   // zero value is ""

    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

**Output:**

```
0 0 false ""
```

> This feature ensures that all variables in Go always have a **well-defined initial value**, even if you don’t explicitly assign one.
