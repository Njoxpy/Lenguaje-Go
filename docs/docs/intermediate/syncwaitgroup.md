#  `sync.WaitGroup` — The Missing Coordinator

Up to now you know:

* Goroutines start independently
* `main` does **not** wait
* Channels are for **communication**

But sometimes you don’t want to pass data.
You just want to say:

> “Wait until all goroutines finish.”

That’s exactly what `WaitGroup` is for.


##  What is a `WaitGroup`?

> A `WaitGroup` waits for a **collection of goroutines** to finish.

It answers **one question only**:

* *“Are all workers done?”*

No data.
No messages.
Just completion.


##  Core Methods (memorize these)

```go
wg.Add(n)   // number of goroutines to wait for
wg.Done()  // called when one goroutine finishes
wg.Wait()  // blocks until counter reaches zero
```

That’s the whole API.


##  Minimal Example

```go
var wg sync.WaitGroup

func worker(id int) {
    defer wg.Done()
    fmt.Println("Worker", id, "done")
}

func main() {
    wg.Add(3)

    go worker(1)
    go worker(2)
    go worker(3)

    wg.Wait()
    fmt.Println("All workers finished")
}
```


##  What’s happening internally?

* `Add(3)` → counter = 3
* Each `Done()` → counter − 1
* `Wait()` blocks until counter = 0

Simple. Deterministic. Safe.


##  VERY IMPORTANT RULES (don’t skip)

###  `Add()` must happen **before** starting goroutines

 Wrong:

```go
go worker()
wg.Add(1)
```

 Correct:

```go
wg.Add(1)
go worker()
```

Why?

* Race condition
* `Wait()` might run first

###  Always call `Done()` exactly once

Best practice:

```go
defer wg.Done()
```

This guarantees it runs even on errors.


###  `WaitGroup` ≠ communication

* Do NOT use it to pass data
* Use channels for that

Think:

* **WaitGroup** → “Are you done?”
* **Channel** → “Here is data”


##  Mental Model

* `WaitGroup` = checklist
* Each goroutine checks itself off
* `Wait()` blocks until checklist is empty


##  Common Mistakes

* Forgetting `Done()` → deadlock
* Calling `Add()` inside goroutine → race
* Using `WaitGroup` for data passing → wrong tool


##  Checkpoint (answer briefly)

1. What problem does `WaitGroup` solve?
2. Why must `Add()` be called before starting goroutines?
3. Why is `defer wg.Done()` recommended?
4. When should you prefer channels over `WaitGroup`?

