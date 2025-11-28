## 1. Install Delve

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

- Make sure your `$GOPATH/bin` or `$HOME/go/bin` is in your PATH.
- Check installation:

```bash
dlv version
```

## 2. Start a debug session

Assume you have a Go program `main.go`:

```go
package main

import "fmt"

func Add(x, y int) int {
    return x + y
}

func main() {
    sum := Add(10, 15)
    fmt.Println("Sum:", sum)
}
```

### Start debugging

```bash
dlv debug
```

- This **compiles your program with debugging info** and starts an interactive session.

## 3. Basic `dlv` commands

### 3.1 Breakpoints

Set a breakpoint at a line:

```bash
(dlv) break main.go:6
```

- Or at a function:

```bash
(dlv) break main.Add
```

### 3.2 Run the program

```bash
(dlv) continue
```

- Runs until it hits a breakpoint.

### 3.3 Inspect variables

When stopped at a breakpoint:

```bash
(dlv) print x
(dlv) print y
(dlv) print sum
```

- You can also evaluate expressions:

```bash
(dlv) print x + y
```

### 3.4 Step through code

- Step into functions:

```bash
(dlv) step
```

- Step over current line:

```bash
(dlv) next
```

- Step out of the current function:

```bash
(dlv) stepout
```

### 3.5 List code around current line

```bash
(dlv) list
```

### 3.6 Watch variables

```bash
(dlv) watch sum
```

- Stops execution when `sum` changes.

### 3.7 Backtrace

Show the call stack:

```bash
(dlv) bt
```

## 4. Exit debug session

```bash
(dlv) quit
```

## 5. Debugging tests

You can also debug tests:

```bash
dlv test
```

- Or debug a specific test:

```bash
dlv test -- -test.run TestAdd
```

- Use the same commands (`break`, `print`, `next`, etc.) inside your test.

## 6. Tips & best practices

1. **Use breakpoints on functions rather than line numbers** when refactoring often.
2. **Inspect complex types**: slices, structs, maps using `print`.
3. **Step over vs step into**: `next` skips function internals; `step` enters.
4. Combine with **VS Code or GoLand** for GUI debugging; Delve works behind the scenes.

### Example session:

```bash
$ dlv debug
(dlv) break main.Add
(dlv) continue
(dlv) print x
(dlv) next
(dlv) step
(dlv) print sum
(dlv) continue
(dlv) quit
```

## References

- [Delve is a debugger for the Go programming language.](https://github.com/go-delve/delve)
