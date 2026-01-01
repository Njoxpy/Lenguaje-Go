# 1. Go’s Concurrency Philosophy (Mental Model First)

### Concurrency ≠ Parallelism

- **Concurrency**: dealing with _many things at once_
- **Parallelism**: _doing_ many things at the same time (on multiple cores)

Go focuses on **concurrency**, and lets the runtime decide _when_ and _how_ to run things in parallel.

> **“Don’t communicate by sharing memory; share memory by communicating.”**

This single sentence explains almost everything about Go’s concurrency design.

---

# 2. Goroutines: Lightweight Concurrent Execution

### What is a Goroutine?

A **goroutine** is a function executing concurrently with other goroutines in the same address space.

```go
go doWork()
```

### Why Goroutines Are Special

Compared to OS threads:

| OS Thread            | Goroutine                        |
| -------------------- | -------------------------------- |
| Heavy (~1–2MB stack) | Lightweight (~2KB initial stack) |
| Expensive to create  | Cheap to create                  |
| Managed by OS        | Managed by Go runtime            |
| Fixed stack          | Dynamically growing stack        |

You can run **hundreds of thousands** of goroutines.

---

### How Goroutines Are Scheduled (G-M-P Model)

Internally, Go uses:

- **G** → Goroutine
- **M** → OS Thread (Machine)
- **P** → Processor (logical resource)

Each `P`:

- Has a run queue of goroutines
- Is required to execute Go code

`GOMAXPROCS` controls how many `P`s exist.

```go
runtime.GOMAXPROCS(4)
```

This design:

- Decouples goroutines from OS threads
- Enables cheap context switching
- Allows work stealing

---

### Example: Basic Goroutine

```go
func say(msg string) {
    fmt.Println(msg)
}

func main() {
    go say("world")
    say("hello")
}
```

⚠️ The program may exit before `world` prints.

---

### Waiting for Goroutines (sync.WaitGroup)

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    say("world")
}()

wg.Wait()
```

**Rule:**
If `main` exits, all goroutines die.

---

# 3. Channels: Communication, Not Shared Memory

### What is a Channel?

A **channel** is a typed conduit for sending and receiving values.

```go
ch := make(chan int)
```

### Send / Receive

```go
ch <- 42      // send
x := <-ch     // receive
```

Channels **block by default**:

- Send blocks until someone receives
- Receive blocks until someone sends

---

### Why Channels Exist

Instead of this (shared memory + locks):

```go
mu.Lock()
data++
mu.Unlock()
```

Go prefers:

```go
ch <- 1
```

This reduces:

- Race conditions
- Complex locking logic
- Deadlocks (when used properly)

---

### Unbuffered Channels (Synchronization)

```go
ch := make(chan int)

go func() {
    ch <- 10
}()

fmt.Println(<-ch)
```

This guarantees:

- Sender waits for receiver
- Receiver waits for sender

Used for **synchronization**, not throughput.

---

### Buffered Channels (Throughput)

```go
ch := make(chan int, 3)

ch <- 1
ch <- 2
ch <- 3
```

Buffered channels:

- Allow sending without immediate receiving
- Act like a queue

Use when:

- You want to decouple producers and consumers
- You care about throughput

---

### Closing Channels

```go
close(ch)
```

- Only **sender** should close
- Receiving from a closed channel yields zero value

```go
v, ok := <-ch
if !ok {
    // channel closed
}
```

---

### Range Over Channels

```go
for v := range ch {
    fmt.Println(v)
}
```

Loop exits when channel is closed.

---

# 4. Select: Coordinating Multiple Channels

### select = switch for channels

```go
select {
case v := <-ch1:
    fmt.Println(v)
case ch2 <- 10:
    fmt.Println("sent")
default:
    fmt.Println("nothing ready")
}
```

### Key Properties

- Blocks until one case is ready
- Randomly chooses if multiple are ready
- `default` makes it non-blocking

---

### Timeout Pattern

```go
select {
case res := <-ch:
    fmt.Println(res)
case <-time.After(2 * time.Second):
    fmt.Println("timeout")
}
```

---

# 5. Synchronization Primitives (When Channels Aren’t Ideal)

Channels are great—but **not always the right tool**.

---

## sync.Mutex

Use when:

- Protecting shared state
- Low-level internal structures

```go
var mu sync.Mutex
var count int

func inc() {
    mu.Lock()
    count++
    mu.Unlock()
}
```

### Rule:

- Keep lock scope small
- Never lock in one goroutine and unlock in another

---

## sync.RWMutex

Allows:

- Multiple readers
- One writer

```go
var mu sync.RWMutex

mu.RLock()
read()
mu.RUnlock()

mu.Lock()
write()
mu.Unlock()
```

Good for read-heavy workloads.

---

## sync.Once

Run initialization exactly once.

```go
var once sync.Once

once.Do(initConfig)
```

Thread-safe lazy initialization.

---

## sync.Cond (Advanced)

Used for complex coordination.
Rarely needed if channels suffice.

---

# 6. Common Concurrency Patterns (This Is Where Power Is)

---

## 1. Worker Pool

```go
jobs := make(chan int)
results := make(chan int)

worker := func() {
    for j := range jobs {
        results <- j * 2
    }
}

for i := 0; i < 3; i++ {
    go worker()
}

go func() {
    for i := 0; i < 5; i++ {
        jobs <- i
    }
    close(jobs)
}()

for i := 0; i < 5; i++ {
    fmt.Println(<-results)
}
```

Used in:

- Background processing
- APIs
- Task queues

---

## 2. Fan-Out / Fan-In

Fan-Out: distribute work
Fan-In: collect results

```go
func worker(id int, jobs <-chan int, out chan<- int) {
    for j := range jobs {
        out <- j * 2
    }
}
```

---

## 3. Pipeline

Each stage is a goroutine.

```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
```

---

## 4. Cancellation (context.Context)

**Essential for real systems.**

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    defer cancel()
    doWork(ctx)
}()
```

```go
select {
case <-ctx.Done():
    return
default:
    work()
}
```

Used in:

- HTTP servers
- Background jobs
- Timeouts
- Graceful shutdown

---

# 7. Real-World Rules & Best Practices

### 1. Start with Simplicity

- Don’t add concurrency until needed
- Sequential code is easier to reason about

---

### 2. Prefer Channels for Ownership Transfer

If one goroutine **owns** data → use channels.

---

### 3. Prefer Mutexes for Shared State

If many goroutines **mutate shared data** → use mutex.

---

### 4. Always Think About Lifetime

Ask:

- Who creates this goroutine?
- Who stops it?
- What happens on shutdown?

---

### 5. Avoid Goroutine Leaks

Bad:

```go
go func() {
    ch <- work() // blocks forever if no receiver
}()
```

Good:

- Use buffering
- Use context cancellation
- Ensure receivers exist

---

### 6. Race Detector Is Non-Optional

```bash
go test -race
go run -race main.go
```

Use it early and often.

---

# 8. Why Go Concurrency Works So Well

Go’s design choices:

- Cheap goroutines → fine-grained concurrency
- Channels → structured communication
- Runtime scheduler → efficient CPU use
- Context → composable cancellation
- Simple primitives → fewer footguns

This is why Go dominates:

- Backend services
- Distributed systems
- Networking
- Cloud infrastructure

---

# 9. Mental Checklist Before Writing Concurrent Code

Ask yourself:

1. Can this be sequential?
2. What data is shared?
3. Who owns the data?
4. How do goroutines stop?
5. What happens on failure?

If you can answer these, your concurrency will be solid.

---

## Reference

- [golang concurrency](https://www.concurrency.rocks/)