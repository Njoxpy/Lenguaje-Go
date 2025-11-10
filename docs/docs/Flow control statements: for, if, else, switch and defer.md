# Flow Control Statements in Go

## The `for` Loop

Go features a single, versatile looping construct: the `for` loop. This simplicity in design makes Go's looping mechanism both powerful and easy to understand.

### Basic For Loop Structure

The basic `for` loop consists of three components, separated by semicolons:

1. **Init Statement**: Executed once before the first iteration

   - Often contains a short variable declaration
   - Variables declared here are only visible within the loop's scope

2. **Condition Expression**: Evaluated before each iteration

   - Loop continues as long as this condition is `true`
   - Loop terminates when condition becomes `false`

3. **Post Statement**: Executed at the end of each iteration

### Important Syntax Notes:

- Unlike C, Java, or JavaScript, parentheses `()` are not used around the three components
- Curly braces `{}` are mandatory for the loop body

### Basic For Loop Example

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```

### Flexible For Loop Syntax

The `for` loop in Go is flexible - both the initialization and post statements are optional:

```go
package main

import "fmt"

func main() {
    sum := 1
    for ; sum < 10; {
        sum += sum
    }
    fmt.Println(sum)
}
```

### While Loop Equivalent

In Go, the `for` loop can be used as a while loop by dropping the semicolons. This makes Go's `for` loop serve the same purpose as C's `while`:

```go
package main

import "fmt"

func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
```

### Infinite Loops

To create an infinite loop in Go, simply omit the condition. This creates a clean and concise way to express endless loops:

```go
package main

func main() {
    for {
        // This loop will run forever
    }
}
```

## Conditional Statements

### The `if` Statement

Go's `if` statements follow the same philosophy as its `for` loops:

- No parentheses `()` around the condition
- Curly braces `{}` are mandatory

Here's a basic example:

```go
package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}
```

### If with Initialization

Go allows you to include a short initialization statement before the condition:

- The initialized variables are only in scope for the `if` block
- This is useful for keeping code concise and variables scoped appropriately

Example of if with initialization:

```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
```

### If-Else Statements

In Go, variables declared in an `if` statement's initialization are also accessible in any corresponding `else` blocks:

```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // Note: v is not accessible here
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
```

## Practice Exercise: Implementing Square Root Function

### Problem Description

Create a function that computes the square root of a number using iterative approximation.

### The Challenge

Given a number `x`, find a number `z` where `z²` is as close as possible to `x`.

### Algorithm Overview

1. Use Newton's method for approximation
2. Formula: `z -= (z*z - x) / (2*z)`
3. Each iteration brings us closer to the actual square root

### Implementation Requirements

1. Start with the following template:

```go
package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    // Your implementation here
}

func main() {
    fmt.Println(Sqrt(2))
}
```

### Implementation Steps

1. **Initial Setup**

   - Start with z = 1.0 or z = float64(1)
   - You can also try other starting values like x or x/2

2. **Iteration Process**

   - Initially, run the calculation 10 times
   - Print each intermediate value of z
   - Later, modify to stop when changes become very small

3. **Optimization**
   - Experiment with different stopping conditions
   - Compare your results with math.Sqrt
   - Try different initial values for z

### Note on Newton's Method

The formula `z -= (z*z - x) / (2*z)` works because:

- `(z*z - x)` measures how far z² is from the target x
- Division by `2z` (the derivative of z²) scales the adjustment
- This method converges quickly for square roots

Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

1
package main
2
​
3
import (
4
"fmt"
5
"runtime"
6
)
7
​
8
func main() {
9
fmt.Print("Go runs on ")
10
switch os := runtime.GOOS; os {
11
case "darwin":
12
fmt.Println("macOS.")
13
case "linux":
14
fmt.Println("Linux.")
15
default:
16
// freebsd, openbsd,
17
// plan9, windows...
18
fmt.Printf("%s.\n", os)
19
}
20
}
21
​

Switch evaluation order
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

switch i {
case 0:
case f():
}
does not call f if i==0.)

Note: Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.

1
package main
2
​
3
import (
4
"fmt"
5
"time"
6
)
7
​
8
func main() {
9
fmt.Println("When's Saturday?")
10
today := time.Now().Weekday()
11
switch time.Saturday {
12
case today + 0:
13
fmt.Println("Today.")
14
case today + 1:
15
fmt.Println("Tomorrow.")
16
case today + 2:
17
fmt.Println("In two days.")
18
default:
19
fmt.Println("Too far away.")
20
}
21
}
22
​
