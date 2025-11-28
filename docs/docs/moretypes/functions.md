# 1. Function Values

In Go, **functions are first-class citizens**, meaning you can:

- Assign functions to variables
- Pass functions as arguments
- Return functions from other functions

---

### Example 1 — Assigning a function to a variable

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    var f func(int, int) int  // declare a variable of function type
    f = add                    // assign function to variable
    result := f(3, 4)
    fmt.Println(result)        // 7
}
```

**Use case:** Dynamic behavior selection, callbacks, or strategy patterns.

---

### Example 2 — Passing a function as argument

```go
func apply(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    sum := func(x, y int) int { return x + y }
    product := func(x, y int) int { return x * y }

    fmt.Println(apply(3, 4, sum))     // 7
    fmt.Println(apply(3, 4, product)) // 12
}
```

**Use case:** Higher-order functions, reusable computation logic.

---

# 2. Function Closures

A **closure** is a function that **captures variables from its surrounding scope**.

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    posSum := adder()
    fmt.Println(posSum(1)) // 1
    fmt.Println(posSum(2)) // 3
    fmt.Println(posSum(3)) // 6
}
```

- The returned function **remembers `sum`** even after `adder()` has finished.
- Closures are useful for maintaining state between calls without global variables.

**Use case:** Counters, memoization, event handlers.

---

# Step-by-step quizzes

---

## Beginner

### Quiz 1 — Function value assignment

1. Write a function `greet(name string) string` that returns `"Hello, <name>"`.
2. Assign it to a variable `f` and call `f("Alice")`.
3. Print the result.

**Use case:** Simple dynamic greetings or callback examples.

---

### Quiz 2 — Function as argument

1. Write two functions: `square(n int) int` and `cube(n int) int`.
2. Write a higher-order function `compute(n int, fn func(int) int) int` that applies `fn` to `n`.
3. Test `compute(3, square)` and `compute(2, cube)`.

**Use case:** Dynamic computation pipelines, reusable processing.

---

## Intermediate

### Quiz 3 — Closure as counter

1. Write a function `counter()` that returns a closure `func() int` which increments and returns a number every time it’s called.
2. Create two counters `c1` and `c2` and show they maintain separate counts.

**Use case:** Unique ID generators, counters in a package.

---

### Quiz 4 — Accumulator

1. Write a function `accumulator(start int)` that returns a closure `func(int) int`.
2. Each call of the closure adds the argument to a running total starting from `start`.
3. Test with a few numbers.

**Use case:** Financial transactions or running totals in a program.

---

## Advanced

### Quiz 5 — Memoization with closure

1. Write a function `FibMemo()` that returns a closure `func(n int) int`.
2. The closure calculates Fibonacci numbers and stores previously calculated results in a map.
3. Test `FibMemo()(10)` and observe speed improvements on repeated calls.

**Use case:** Optimize expensive recursive calculations dynamically.

---

### Quiz 6 — Function returning different operations

1. Write a function `operation(op string) func(int, int) int` which returns a closure performing addition, subtraction, multiplication, or division depending on `op`.
2. Test `operation("add")(3, 5)` and `operation("mul")(2, 4)`.

**Use case:** Dynamic command processing or calculator design.

---

### Quiz 7 — Stateful event handlers

1. Write a closure `logger(prefix string)` that returns `func(msg string)` and prints `"[prefix] msg"` every time it’s called.
2. Create two loggers `infoLogger` and `errorLogger` with different prefixes.
3. Call them with messages and observe how each keeps its prefix.

**Use case:** Logging system with different contexts.

---

# Key notes

1. Function values can be stored, passed, or returned like any other variable.
2. Closures **capture the surrounding environment**.
3. Use closures for **stateful functions** without global variables.
4. Higher-order functions make your code flexible and reusable.
