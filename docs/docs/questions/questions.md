# 🧩 BEGINNER QUESTIONS

**1. Pointers Basics**

- What is the difference between `*` and `&` in Go?
- Given the code, what is the output? Explain why.

```go
a := 10
b := &a
*b = 20
fmt.Println(a, *b)
```

- Declare an integer variable, create a pointer to it, and modify its value through the pointer.

**2. Structs Basics**

- Define a struct `User` with fields `Name` and `Age`. Create an instance and print the fields.
- Given a struct, what happens when you assign it to another variable and modify the copy?
- Initialize a struct using both **ordered values** and **named fields**.

**3. Arrays Basics**

- Declare an array of 5 integers. Set the first and last elements, then print the array.
- How does Go handle uninitialized elements in an array?

**4. Slices Basics**

- Declare a slice of strings with three values. Append two more values and print the slice.
- Given a slice `s := []int{1,2,3,4}`, what is the output of `s[1:3]`?
- What is the difference between arrays and slices in Go?

---

# 🧩 INTERMEDIATE QUESTIONS

**1. Pointers to Structs**

- Define a struct `Account` with a field `Balance`. Write a function `Deposit` that increases balance via pointer.
- Explain what happens when you assign a pointer to a struct versus copying the struct.
- Create two pointers pointing to the same struct. Modify the struct through one pointer. What is visible through the other pointer?

**2. Structs and Methods**

- Write a method `IncrementAge` for a `User` struct that increases age by 1. Should you use a pointer receiver or value receiver? Explain.
- Create a nested struct `Employee` with an `Address` struct inside. Write a function that updates the city via pointer.

**3. Arrays & Slices**

- Given an array `[5]int{1,2,3,4,5}`, create a slice `[2:4]` and modify it. What happens to the original array?
- Write a function that takes a slice of integers and doubles every value. Will the original slice be modified? Explain.
- Explain the difference between `len()` and `cap()` for slices with examples.

---

# 🧩 ADVANCED QUESTIONS

**1. Structs + Pointers + Functions**

- Implement a `Bank` struct containing a slice of `Account` structs. Write a function that transfers money from one account to another using pointers.
- Explain what happens internally when a slice is passed to a function. Does modifying the slice affect the original?

**2. Memory & References**

- Given the code below, what will it print and why?

```go
type Node struct {
    Value int
    Next  *Node
}

n1 := Node{Value:1}
n2 := n1
n2.Next = &Node{Value:2}

fmt.Println(n1.Next, n2.Next)
```

- Explain the difference between shallow copy and deep copy with structs containing pointers.

**3. Practical Slice Manipulation**

- Implement a function that removes an element at a given index from a slice of integers without using extra slices.
- Given a slice of integers, write a function that returns a new slice containing only the even numbers. Do not modify the original slice.

**4. Combined Problem (Pointers + Slices + Structs)**

- Implement a `Product` struct with `Name` and `Stock`.
- Create a slice of `Product` structs.
- Write a function `Restock(products []*Product, n int)` that increases stock of all products by `n` using pointers.
- What would happen if you pass `[]Product` (not pointers) instead?

**5. Big Tech-style Challenge**

- Design a simple in-memory `UserManager` with:

  - Slice of `User` structs
  - Functions: `AddUser`, `UpdateEmail`, `DeleteUser`
  - Ensure `UpdateEmail` modifies the original struct efficiently.

- Explain why slices of pointers might be preferred over slices of structs in this scenario.

---

# Maps in Go — clear explanation + step-by-step quizzes & real uses

Good — let’s learn **maps**. I’ll give a short explanation, show common operations with code, then a set of step-by-step quizzes (beginner → advanced) you can code and test. I won’t include answers so you can practice.

---

# What is a map?

A **map** is an unordered collection of key → value pairs. Keys must be of a comparable type (strings, ints, pointers, structs with comparable fields, etc.). Maps are reference types and grow dynamically.

Basic form:

```go
var m map[string]int          // nil map, must be initialized before use
m = make(map[string]int)     // initialize with make
m["apple"] = 5               // set
v := m["apple"]              // get (zero value if absent)
v, ok := m["apple"]          // comma-ok form to check presence
delete(m, "apple")           // remove key
```

Common uses

- Frequency counters (word/char counts)
- Caches and memoization
- Lookups (user ID → profile)
- Indexes (value → positions)
- Adjacency lists for graphs

---

# Quick reference examples

Create and initialize:

```go
m := map[string]int{
    "alice": 30,
    "bob":   25,
}
```

Add / modify / read:

```go
m["carol"] = 28
age := m["alice"]
age, ok := m["dan"]  // ok == false if "dan" not present
```

Delete and length:

```go
delete(m, "bob")
fmt.Println(len(m))
```

Iterate (order not guaranteed):

```go
for k, v := range m {
    fmt.Println(k, v)
}
```

Nested maps:

```go
nested := make(map[string]map[string]int)
nested["group1"] = make(map[string]int)
nested["group1"]["x"] = 1
```

Map keys constraints:

- Keys must be comparable (`==` allowed). Slices, maps, and functions cannot be keys.

---

# Step-by-step quizzes (no answers included)

Each task includes code skeleton + steps. Run `go run` / `go test` to validate your work.

---

## Beginner

### Quiz 1 — Basic creation & access

1. Create a file `quiz1.go`.
2. Create and initialize a map `ages` with entries: `"Alice": 30`, `"Bob": 25`.
3. Print the age of `"Alice"`.
4. Read `"Charlie"` from the map using comma-ok and print whether the key exists.

Skeleton:

```go
package main

import "fmt"

func main() {
    // 1. create and initialize ages
    // 2. print ages["Alice"]
    // 3. read ages["Charlie"] with comma-ok and print existence
}
```

Use-case: simple lookup table for user age.

---

### Quiz 2 — Add / modify / delete

1. Start with an empty map `inventory := make(map[string]int)`.
2. Add `"apple": 10`, `"banana": 5`.
3. Increase apple by 5.
4. Remove `"banana"`.
5. Print final map and its length.

Skeleton:

```go
package main

import "fmt"

func main() {
    inventory := make(map[string]int)
    // steps...
}
```

Use-case: inventory counts in a small store.

---

## Intermediate

### Quiz 3 — Frequency counter

1. Write a function `WordCount(s string) map[string]int` that returns counts of words separated by spaces.
2. Test with `"go go gopher go"`.
3. Print top 2 words by count (just iterate and identify).

Skeleton:

```go
package main

import (
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    // implement
}

func main() {
    counts := WordCount("go go gopher go")
    fmt.Println(counts)
}
```

Use-case: text analytics / simple log analysis.

---

### Quiz 4 — Map with struct values

1. Define `type User struct { Name string; Email string }`.
2. Create `users := make(map[int]User)` (map user ID → User).
3. Add at least two users.
4. Update one user’s email (show how to read-modify-write).
5. Print all users.

Skeleton:

```go
package main

import "fmt"

type User struct {
    Name  string
    Email string
}

func main() {
    users := make(map[int]User)
    // steps...
}
```

Use-case: in-memory user store.

---

### Quiz 5 — Map of slices (groups)

1. Create `groups := make(map[string][]string)`.
2. Add members to group `"admins"` and `"devs"`.
3. Append another member to `"devs"`.
4. Iterate and print each group and members.

Skeleton:

```go
package main

import "fmt"

func main() {
    groups := make(map[string][]string)
    // steps...
}
```

Use-case: grouping users by roles.

---

## Advanced

### Quiz 6 — Nested maps & initialization helper

1. Create `permissions := make(map[string]map[string]bool)` where outer key is username, inner map is permission → bool.
2. Write helper `func SetPerm(m map[string]map[string]bool, user, perm string, value bool)` that ensures inner map is initialized before setting.
3. Set a few permissions and print them.

Skeleton:

```go
package main

import "fmt"

func SetPerm(m map[string]map[string]bool, user, perm string, value bool) {
    // ensure inner map exists, then set
}

func main() {
    perms := make(map[string]map[string]bool)
    // use SetPerm and print results
}
```

Use-case: per-user permission matrix in a service.

---

### Quiz 7 — Map concurrency warning & safe access

1. Create a read-heavy simulation where multiple goroutines read from a shared map concurrently.
2. Demonstrate that writing to the map concurrently can panic.
3. Refactor to use `sync.RWMutex` to safely support concurrent reads and occasional writes.

Skeleton:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    m := make(map[int]string)
    // populate m

    // start many goroutines that read m

    // optionally start goroutine that writes to m (show panic without mutex, fix with mutex)
}
```

Use-case: shared in-memory cache in concurrent server.

---

### Quiz 8 — Graph adjacency list

1. Use a map to represent an undirected graph: `graph := make(map[string][]string)`.
2. Implement `AddEdge(g map[string][]string, a, b string)` that adds an edge both ways.
3. Implement `Neighbors(g map[string][]string, node string) []string`.
4. Build a small graph, print neighbors of a node, and perform a simple BFS to find if two nodes are connected.

Skeleton:

```go
package main

import "fmt"

func AddEdge(g map[string][]string, a, b string) {
    // add both directions
}

func Neighbors(g map[string][]string, node string) []string {
    return g[node]
}

func main() {
    g := make(map[string][]string)
    // build graph, print neighbors, BFS simple connectivity
}
```

Use-case: routing, network topology, recommendation engines.

---

### Quiz 9 — Map key constraints & custom key

1. Define `type Point struct { X, Y int }` and use `map[Point]string` (color).
2. Insert a few points and read them.
3. Try to use a slice as a key and observe the compiler error — explain why slices cannot be map keys.

Skeleton:

```go
package main

import "fmt"

type Point struct { X, Y int }

func main() {
    m := make(map[Point]string)
    m[Point{1,2}] = "red"
    fmt.Println(m[Point{1,2}])
    // try map[[]int]string // should not compile
}
```

Use-case: spatial indexing, quick lookups for coordinate-based data.

---

### Quiz 10 — Large-scale memoization

1. Implement a Fibonacci with memoization using `map[int]int` to cache results: `func Fib(n int, memo map[int]int) int`.
2. Measure time with and without memoization for n=40.
3. Discuss memory/time tradeoffs.

Skeleton:

```go
package main

import (
    "fmt"
    "time"
)

func Fib(n int, memo map[int]int) int {
    // implement with memo
}

func main() {
    memo := make(map[int]int)
    start := time.Now()
    fmt.Println(Fib(40, memo))
    fmt.Println("Elapsed:", time.Since(start))
}
```

Use-case: dynamic programming optimization, expensive computation caching.

---

# Extra tips & gotchas

- Maps are reference types: passing to functions gives access to same underlying data.
- Reading from a nil map returns zero value; writing to a nil map causes panic. Always initialize (`make`) before write.
- Maps are not safe for concurrent writes — use sync.Mutex or sync.RWMutex or `sync.Map` for concurrent access.
- Map iteration order is random; do not rely on ordering. If order matters, use a slice of keys and sort it.

---

If you want, I can:

- Provide **a small test file** for one of the quizzes (e.g., WordCount) so you can run `go test`.
- Or produce a **cheat sheet** listing functions and idioms for maps.

Which would you like next?
