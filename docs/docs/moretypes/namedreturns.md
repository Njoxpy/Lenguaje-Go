### Named Return Values in Go: Clear and Balanced Notes

#### What Are Named Return Values?

You can optionally **name** the return parameters in a function's signature. These names act like variables that are:

- Automatically declared at the start of the function.
- Initialized to their **zero values**.

A plain `return` statement (with no arguments) is called a **naked return** — it returns the current values of these named variables.

**Simple example:**

```go
func add(x, y int) (result int) {
    result = x + y
    return  // naked return
}
```

**Tour example:**

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // split(17) → 7 10
}
```

#### When Named Returns Are Helpful

- **Self-documenting**: Especially useful when returning multiple values of the same type.
  ```go
  func coords() (lat, lng float64)  // clearer than (float64, float64)
  ```
- **Works well with `defer`**: Common pattern for error handling or cleanup without repeating code.
  ```go
  func process() (result string, err error) {
      defer func() {
          if err != nil {
              log.Println("error:", err)
          }
      }()
      // ... main logic
      return
  }
  ```
- Used in the standard library and considered idiomatic in specific cases.

#### Common Advice and Criticism

The main controversy is around **naked returns**, not the naming itself:

- Naked `return` can reduce readability in longer functions (you have to scroll up to see what’s returned).
- Many teams and linters (e.g., `nakedret`) discourage or ban naked returns.
- Effective Go warns against them in longer functions.

**Named returns alone** (without naked returns) are generally viewed as harmless and sometimes beneficial for clarity.

#### Community Consensus

- Avoid naked returns in most real-world code.
- Use named returns sparingly — mainly when they improve documentation or enable useful `defer` patterns.
- There is no serious movement to remove the feature; it has been part of Go since 1.0 and remains stable.

#### Recommendation for Learners

- Follow the Tour examples to understand the concept.
- In your own code, prefer explicit returns (`return value, err`) unless naming adds real clarity or enables a clean `defer`.
- Keep functions short and readable — that solves most issues.

That’s the balanced view: named returns are a tool with specific good uses, but naked returns are what most experienced developers caution against.
