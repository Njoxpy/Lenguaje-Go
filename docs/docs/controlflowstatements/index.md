## **For Loops in Go**

Go has only **one looping construct**: the `for` loop.

The basic `for` loop consists of **three components** separated by semicolons:

1. **Initialization statement** – executed before the first iteration
2. **Condition expression** – evaluated before every iteration
3. **Post statement** – executed at the end of each iteration

---

### **How it Works**

- The **init statement** is commonly a short variable declaration.
  Variables declared here are visible **only within the scope of the for loop**.
- The loop continues to execute **as long as the condition is true**.
- Once the condition evaluates to `false`, the loop stops.
- Unlike C, Java, or JavaScript:

  - There are **no parentheses** around the components.
  - The **curly braces `{}` are mandatory**, even for single-line loops.

---

### **Example**

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

**Explanation:**

- `i := 0` initializes the counter.
- The loop runs while `i < 10`.
- After each iteration, `i++` increments `i` by 1.
- The loop accumulates the sum of numbers from 0 to 9.

**Output:**

```
45
```

## **For (continued)**

In Go, the **init** and **post** statements in a `for` loop are **optional**.
You can omit one or both, depending on the logic you need.

---

### **Example**

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

---

### **Explanation**

- The **init** and **post** sections are left empty.
- Only the **condition** (`sum < 10`) remains between the semicolons.
- This is perfectly valid Go syntax and works just like a **while loop** in other languages.

The above example doubles the value of `sum` until it is no longer less than 10.

---

### **Output**

```
16
```

---

### **Tip**

If both the **init** and **post** statements are omitted, the semicolons themselves can also be removed — turning it into Go’s **while-style** loop:

```go
for sum < 10 {
	sum += sum
}
```

## **For (continued)**

The **init** and **post** statements in a Go `for` loop are **optional**.
If you omit them, the `for` loop still works — it simply becomes more flexible, resembling a `while` loop in other languages.

---

### **Example**

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

---

### **Explanation**

- The **init** and **post** statements are left empty (`;` marks their positions).
- This means:

  - There is **no initialization** before the loop starts.
  - There is **no automatic update** after each iteration.

- The loop runs **as long as** the condition `sum < 10` is `true`.

Here’s what happens step by step:

| Step | Value of `sum` | Condition (`sum < 10`) | Action            |
| ---- | -------------- | ---------------------- | ----------------- |
| 1    | 1              | true                   | `sum += sum → 2`  |
| 2    | 2              | true                   | `sum += sum → 4`  |
| 3    | 4              | true                   | `sum += sum → 8`  |
| 4    | 8              | true                   | `sum += sum → 16` |
| 5    | 16             | false                  | Loop ends         |

---

### **Output**

```
16
```

---

### **Key Insight**

When you omit both the **init** and **post** parts, the `for` loop in Go acts like a **while loop** in other languages.

Equivalent **while-style syntax** in pseudocode:

```go
for sum < 10 {
	sum += sum
}
```

This demonstrates Go’s flexibility — you can use `for` in multiple forms, including the traditional C-style loop or simplified while-like loops.

## **For as Go's "While"**

In Go, there is **no separate `while` loop**.
You can use a `for` loop **without semicolons** to achieve the same behavior.

This makes the `for` loop Go’s equivalent of `while` in other languages like C or Java.

---

### **Example**

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

---

### **Explanation**

- The **condition** `sum < 1000` is checked before each iteration.
- There are **no init or post statements**; the semicolons are omitted.
- The loop continues **as long as the condition is true**.
- Each iteration doubles the value of `sum`.

---

### **Output**

```
1024
```

---

### **Key Insight**

- Go’s `for` loop is **versatile**:

  - Traditional `for` with `init; condition; post`
  - `for` as a **while loop** (`for condition { ... }`)
  - Infinite loop (`for { ... }`)

This flexibility makes Go loops **simple and powerful**.

Forever
If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```go
package main

func main() {

    for {

    }

}
```

## **If Statements**

In Go, **`if` statements** are similar to `for` loops in syntax:

- The **condition does not require parentheses** `()`.
- **Braces `{}` are required**, even for single-line blocks.

---

### **Example**

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

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

---

### **Explanation**

1. The function `sqrt` computes the square root of `x`.
2. The `if` statement checks if `x` is negative:

   - If yes, it recursively calls `sqrt` with `-x` and appends `"i"` to indicate an imaginary number.
   - If not, it returns the square root of `x` using `math.Sqrt`.

3. The `fmt.Sprint` function converts the float result to a string for consistent return type.

---

### **Output**

```
1.4142135623730951 2i
```

---

### **Key Points**

- Go allows **inline `if` conditions** without parentheses:

  ```go
  if x < 0 { ... }
  ```

- Braces are **mandatory** and define the scope of the conditional block.
- Go supports **recursive function calls** as shown in the example.

## **If Statement with a Short Statement**

In Go, an `if` statement can include a **short statement** that runs **before the condition**.

- Variables declared in this short statement are **only in scope within the `if` block**.
- This is useful for computations that are only needed for the conditional check.

---

### **Example**

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

---

### **Explanation**

1. The `if` statement starts with a short statement:

   ```go
   v := math.Pow(x, n)
   ```

   This computes `x` raised to the power `n` and assigns it to `v`.

2. The condition `v < lim` is evaluated immediately after the short statement.
3. `v` is **only visible inside the `if` statement**.
4. If `v < lim`, the function returns `v`; otherwise, it returns `lim`.

---

### **Output**

```
9 20
```

---

### **Key Points**

- Short statements in `if` allow **clean, inline computation** before the condition.
- Scope of variables declared this way is **limited to the `if` block**, which helps avoid polluting the outer scope.

Example of limited scope:

```go
if v := computeSomething(); v > 10 {
    fmt.Println(v) // OK
}
// fmt.Println(v) // Error: v is undefined here
```

## **Switch Statement**

A **`switch` statement** in Go is a concise way to write a sequence of `if-else` statements.

- It **runs the first case** whose value matches the switch expression.
- Unlike C, C++, Java, or JavaScript:

  - Go **automatically breaks** after the matched case. No explicit `break` is needed.
  - Switch **cases do not have to be constants**, and values do not need to be integers.

---

### **Example**

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

---

### **Explanation**

1. `os := runtime.GOOS` initializes a variable with the operating system name.
2. The `switch` evaluates `os` and executes **only the first matching case**:

   - `"darwin"` → prints `"macOS."`
   - `"linux"` → prints `"Linux."`
   - `default` → prints any other OS.

3. No `break` statements are required — Go stops automatically after executing the matching case.

---

### **Output** (Example on Linux)

```
Go runs on Linux.
```

---

### **Key Points**

- Go’s `switch` is **more flexible** than many other languages:

  - Cases can be any **expression**, not just constants.
  - Only the **matching case executes**; no fallthrough unless explicitly requested with the `fallthrough` keyword.

- Variables declared in the **switch statement** (like `os` above) are scoped **only within the switch**.

---

## **Switch Evaluation Order**

In Go, **switch cases are evaluated from top to bottom**.
The switch **stops at the first case that succeeds** — later cases are **not evaluated**.

> Example:
>
> ```go
> switch i {
> case 0:
> case f():
> }
> ```
>
> If `i == 0`, the function `f()` is **not called**, because the first case matched.

---

### **Example: Checking Days Until Saturday**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

---

### **Explanation**

1. `today := time.Now().Weekday()` gets the current day of the week.
2. The `switch` statement checks `time.Saturday` against relative values:

   - `today + 0` → today
   - `today + 1` → tomorrow
   - `today + 2` → in two days

3. The first case that matches executes, and the switch **exits immediately**.
4. The `default` case runs if none of the other cases match.

---

### **Output** (Example)

```
When's Saturday?
Tomorrow.
```

> Note: In the Go Playground, `time.Now()` starts at `2009-11-10 23:00:00 UTC`.

---

### **Key Points**

- Go evaluates **cases sequentially** from top to bottom.
- Only the **first successful case executes**.
- This allows you to include **complex expressions** in later cases without worrying about unnecessary evaluation.

## **Switch with No Condition**

A **switch without a condition** in Go is equivalent to `switch true`.
This allows you to write **clean, readable if-then-else chains** without nesting multiple `if` statements.

---

### **Example**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

---

### **Explanation**

1. `switch { ... }` without a condition automatically evaluates **each case as a boolean expression**.
2. The first case that evaluates to `true` is executed.
3. The `default` case runs if **no other case matches**.

This pattern is ideal for **multiple conditional branches** where only one branch should run.

---

### **Output** (Example)

```
Good afternoon.
```

> Output will vary depending on the current system time.

---

### **Key Points**

- `switch` without a condition behaves like `switch true`.
- Each `case` is a boolean expression.
- Only the **first true case** executes.
- It improves readability over long chains of `if-else if-else` statements.

## **Defer Statement**

In Go, the **`defer`** statement postpones the execution of a function until the **surrounding function returns**.

- The **arguments** of the deferred function are evaluated immediately.
- The **function call itself** is executed later, just before the function exits.

This is often used for **cleanup tasks**, such as closing files or releasing resources.

---

### **Example**

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

---

### **Explanation**

1. `defer fmt.Println("world")` schedules the `Println` call to run **after `main` returns**.
2. `fmt.Println("hello")` runs immediately.
3. When `main` is about to exit, the deferred call executes, printing `"world"`.

---

### **Output**

```
hello
world
```

---

### **Key Points**

- Multiple `defer` statements are **executed in LIFO order** (last-in, first-out).
- Useful for **resource cleanup**:

```go
f, _ := os.Open("file.txt")
defer f.Close() // ensures file is closed at the end of the function
```

- Arguments are evaluated **at the point of defer**, but the call happens **later**.

## **Stacking Defers**

In Go, **deferred function calls are pushed onto a stack**.
When the surrounding function returns, the deferred calls are executed in **last-in, first-out (LIFO) order**.

This allows you to schedule multiple cleanup or logging tasks and ensures they run in reverse order of their deferral.

---

### **Example**

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

---

### **Explanation**

1. `fmt.Println("counting")` prints immediately.
2. The `for` loop defers printing numbers `0` to `9`. Each deferred call is **pushed onto the stack**.
3. `fmt.Println("done")` prints next.
4. When `main` returns, deferred calls execute in **reverse order**, printing `9` down to `0`.

---

### **Output**

```
counting
done
9
8
7
6
5
4
3
2
1
0
```

---

### **Key Points**

- Deferred calls are executed **after the surrounding function finishes**.
- They run in **LIFO order**, so the **last deferred function executes first**.
- Useful for:

  - Reversing operations
  - Cleaning up resources in reverse order of acquisition

> ✅ You have completed the **Defer** lesson!
> You can now move on to the next lesson or explore other Go topics.

## **If and Else**

In Go, variables declared inside an **`if` short statement** are also available inside any associated **`else`** block.
However, these variables are **scoped only within the `if` and `else` statements** — they cannot be accessed outside.

---

### **Example**

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
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

---

### **Explanation**

1. The variable `v` is declared inside the `if` statement:

   ```go
   if v := math.Pow(x, n); v < lim {
   ```

   It exists only within the `if` and `else` blocks.

2. If `v` is less than `lim`, the function returns `v`.
   Otherwise, the `else` block runs and prints a message.

3. After the `if-else`, `v` is **out of scope** — trying to use it there would cause a compiler error.

---

### **Output**

```
27 >= 20
9 20
```

---

### **Key Points**

- Variables declared in the `if` short statement (`if v := ...; condition`)
  are available **only within** that `if` and its corresponding `else`.
- The `else` block runs **only if** the `if` condition evaluates to `false`.
- This makes Go’s `if` syntax concise and allows safe, scoped variable declarations.

---

✅ **Summary:**
Go’s `if` and `else` statements can include short variable declarations, which remain local to those blocks — clean and efficient for conditional logic.

## **Switch Statement**

A **`switch`** statement in Go provides a concise way to handle multiple conditional branches.
It’s often cleaner and more readable than writing multiple `if–else` chains.

---

### **How It Works**

- A `switch` runs the **first case** that matches the condition.
- Unlike C, C++, or Java, **Go does not require a `break` statement** — it stops automatically after a matching case.
- Cases **don’t have to be constants** or integers — they can be strings, variables, or expressions.

---

### **Example**

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

---

### **Explanation**

1. The **short statement** `os := runtime.GOOS` assigns the operating system name.
   It’s scoped only within the `switch`.

2. Go compares `os` against each `case`:

   - `"darwin"` → macOS
   - `"linux"` → Linux
   - `default` → all other OS names

3. The first matching case runs, and Go **automatically breaks** afterward.

---

### **Possible Outputs**

Depending on your system:

```
Go runs on Linux.
```

or

```
Go runs on macOS.
```

or

```
Go runs on windows.
```

---

### **Key Points**

- No need for manual `break` — it’s implicit.
- Cases can evaluate **any type of expression**.
- Short variable declarations (like `os := ...`) are scoped to the `switch` only.
- `default` handles all unmatched cases, similar to the final `else`.

---

✅ **Tip:**
Use `switch` when checking one value against several possibilities — it makes code **cleaner and easier to read** than long `if-else` chains.

### 🌀 **Switch Evaluation Order in Go**

In Go, **switch cases are checked from top to bottom**, and evaluation **stops immediately** when a matching case is found — no further cases are evaluated.

---

### **Example Code**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

---

### **Explanation**

1. `today := time.Now().Weekday()`
   → gets the current day of the week (like `Monday`, `Tuesday`, etc.).

2. The `switch` compares `time.Saturday` to each case **in order**:

   - `today + 0` → means _today is Saturday_.
   - `today + 1` → means _Saturday is tomorrow_.
   - `today + 2` → means _Saturday is in two days_.
   - Otherwise → `"Too far away."`

3. Go **stops evaluating** after the first matching case — it does **not** check further.

---

### **Example Outputs**

If today is **Friday**:

```
When's Saturday?
Tomorrow.
```

If today is **Saturday**:

```
When's Saturday?
Today.
```

If today is **Tuesday**:

```
When's Saturday?
Too far away.
```

---

### ⚙️ **Key Concepts**

- Switches evaluate **top to bottom**.
- Go **stops** after the **first successful** match (no “fallthrough” unless explicitly stated).
- Each `case` can be **an expression**, not just a constant.
- The **`default`** clause runs if **no other case** matches.

---

💡 **Trivia:**
In the Go Playground, `time.Now()` is **always** `2009-11-10 23:00:00 UTC` — it’s a fixed date for deterministic program output.
So in that environment, the day is **Tuesday**, making the output always:

```
When's Saturday?
Too far away.
```

### 🌇 **Switch with No Condition in Go**

A **switch without a condition** in Go is equivalent to writing:

```go
switch true {
```

This allows each `case` to act as a boolean expression — similar to a clean, readable `if-else if-else` chain.

---

### **Example Code**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

---

### **Explanation**

1. `t := time.Now()`
   → captures the current system time.

2. The `switch` **has no condition**, so Go automatically treats it as `switch true`.

3. Each `case` is evaluated **in order**:

   - If `t.Hour() < 12`, prints `"Good morning!"`
   - Else if `t.Hour() < 17`, prints `"Good afternoon."`
   - Otherwise, `"Good evening."`

4. Once a matching case is found, Go **exits the switch** — no further cases are checked.

---

### **Example Outputs**

If the current time is:

| Time (24h) | Output          |
| ---------- | --------------- |
| 09:45      | Good morning!   |
| 14:30      | Good afternoon. |
| 20:10      | Good evening.   |

---

### ⚙️ **Key Points**

- `switch` without a condition is shorthand for `switch true`.
- Makes long `if-else` chains more **compact and readable**.
- Each case should be a **boolean expression**.
- Execution stops after the **first true** case.

---

💡 **Tip:**
You can use this style for multiple complex conditions — it keeps your Go code simple, expressive, and elegant.

### 🕒 **Defer in Go**

A **`defer` statement** postpones the execution of a function **until the surrounding function returns**.

When you call `defer`, the **arguments** are evaluated immediately, but the **function call itself** is executed **later** — right before the function exits.

---

### **Example Code**

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

---

### **Explanation**

1. `defer fmt.Println("world")`
   → schedules `"world"` to be printed **after** `main()` finishes.

2. `fmt.Println("hello")`
   → executes immediately and prints `"hello"`.

3. Once `main()` is about to exit, Go executes the deferred statement, printing `"world"`.

---

### **Output**

```
hello
world
```

---

### ⚙️ **Key Points**

- `defer` is executed **after** the surrounding function ends.
- **Arguments** are evaluated immediately when the `defer` is declared.
- Useful for tasks like:

  - Closing files (`defer file.Close()`)
  - Unlocking mutexes
  - Releasing network connections or database resources

---

💡 **Tip:**
Think of `defer` as a “**cleanup-on-exit**” mechanism — perfect for writing cleaner, safer Go code.

### 🔁 **Stacking `defer` in Go**

In Go, **deferred function calls** are **pushed onto a stack**.
When the surrounding function (`main`, in this case) returns, **deferred calls are executed in _last-in-first-out (LIFO)_ order** — like popping items from a stack.

---

### **Example Code**

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

---

### **Explanation**

1. The loop runs from `i = 0` to `i = 9`.
   Each time, it **defers** `fmt.Println(i)` — adding it to the stack.
2. When the loop finishes, all 10 deferred calls are stacked.
3. As `main()` returns:

   - Go pops and executes each deferred call **in reverse order**.

---

### **Execution Order**

| Step | Action                     | Output        |
| ---- | -------------------------- | ------------- |
| 1    | Print `"counting"`         | counting      |
| 2    | Loop defers 10 print calls | (nothing yet) |
| 3    | Print `"done"`             | done          |
| 4    | Execute defers (LIFO)      | 9 → 0         |

---

### **Final Output**

```
counting
done
9
8
7
6
5
4
3
2
1
0
```

---

### 🧠 **Key Takeaways**

- Deferred calls execute **after** the function finishes.
- **LIFO order:** last deferred → executed first.
- Commonly used for **cleanup in reverse order**, like:

  ```go
  defer conn.Close()
  defer file.Close()
  ```

---

💡 **Analogy:**
Think of `defer` as **“stacking up cleanup tasks”** — when you leave the room (function), you clean up everything you stacked, in reverse order.
