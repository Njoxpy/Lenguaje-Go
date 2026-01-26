# Practice Exercises with Defer

## **Exercise 1: Basic Defer Order**

```go
package main
import "fmt"

func exercise1() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")

    fmt.Println("Start")
}

// What prints? Try it!
```

## **Exercise 2: Defer with Variables**

```go
func exercise2() {
    x := 10
    defer fmt.Println("Defer x:", x) // What prints?

    x = 20
    fmt.Println("Current x:", x)
}
```

## **Exercise 3: Simple Payment Logger**

```go
type Payment struct {
    ID     string
    Amount float64
}

func processPayment(p Payment) string {
    defer fmt.Println("Payment processed!")

    if p.Amount <= 0 {
        return "FAILED"
    }

    return "SUCCESS"
}

// Call: processPayment(Payment{"TX1", 100})
```

## **Exercise 4: Return Value Modification**

```go
func addFee(amount float64) (total float64) {
    // Add $5 fee using defer
    defer func() {
        // TODO: Add fee here
    }()

    total = amount
    return
}

// Should return 105 for input 100
```

## **Exercise 5: Error Recovery**

```go
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            result = 0
            fmt.Println("Recovered from:", r)
        }
    }()

    // This might panic if b == 0
    result = a / b
    return
}

// Test: safeDivide(10, 0) should print recovery message
```

## **Exercise 6: Payment Cleanup Pattern**

```go
func processTransaction(amount float64) (status string) {
    // TODO: Add defer to print "Cleanup done"
    // TODO: Add defer to print final status

    if amount > 1000 {
        status = "REJECTED"
        return
    }

    status = "APPROVED"
    return
}

// Make it print:
// Cleanup done
// Status: APPROVED
```

## **Exercise 7: Timing a Payment**

```go
func timePayment() {
    // TODO: Use defer to print how long function took
    // Hint: time.Now() and time.Since()

    // Simulate processing
    for i := 0; i < 1000000; i++ {
        // busy work
    }
}
```

## **Solutions (Try First!)**

<details>
<summary>Solution 1</summary>

```go
// Output:
// Start
// 3
// 2
// 1
```

LIFO: Last In, First Out

</details>

<details>
<summary>Solution 2</summary>

```go
// Output:
// Current x: 20
// Defer x: 10
```

Arguments evaluated when defer is called, not when it runs.

</details>

<details>
<summary>Solution 4</summary>

```go
func addFee(amount float64) (total float64) {
    defer func() {
        total += 5.0  // Modify return value
    }()

    total = amount
    return
}
```

</details>

<details>
<summary>Solution 6</summary>

```go
func processTransaction(amount float64) (status string) {
    defer fmt.Println("Cleanup done")
    defer func() { fmt.Printf("Status: %s\n", status) }()

    if amount > 1000 {
        status = "REJECTED"
        return
    }

    status = "APPROVED"
    return
}
```

Note: Second defer must be a closure to capture final `status`.

</details>

## **Practical Challenge: Build a Payment Processor**

```go
package main

import (
    "fmt"
    "time"
)

type Payment struct {
    ID        string
    Amount    float64
    Fee       float64
    Status    string
    Timestamp time.Time
}

// TODO: Implement with proper defer usage
func (p *Payment) Process() (receipt string) {
    // 1. Defer to calculate processing time
    // 2. Defer to update status in log
    // 3. Defer to add fees
    // 4. Defer for cleanup message

    // Business logic
    if p.Amount <= 0 {
        p.Status = "FAILED"
        return "Invalid amount"
    }

    // Simulate processing
    time.Sleep(50 * time.Millisecond)
    p.Status = "COMPLETED"

    return fmt.Sprintf("Paid $%.2f", p.Amount)
}

func main() {
    p := &Payment{
        ID:     "TX001",
        Amount: 150.0,
        Fee:    2.5,
    }

    fmt.Println(p.Process())
}
```

## **Tips for Practice:**

1. **Start simple** - Just print with defer
2. **Add one defer at a time**
3. **Test edge cases** - What if function returns early?
4. **Use go playground** - Quick experiments: https://go.dev/play/
5. **Think "cleanup"** - What always needs to run?

## **Common Patterns to Memorize:**

```go
// Pattern 1: Cleanup
defer resource.Close()

// Pattern 2: Log timing
start := time.Now()
defer func() { fmt.Println(time.Since(start)) }()

// Pattern 3: Modify return
defer func() { result += bonus }()

// Pattern 4: Panic recovery
defer func() { recover() }()
```
