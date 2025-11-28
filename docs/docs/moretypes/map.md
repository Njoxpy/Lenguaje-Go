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
