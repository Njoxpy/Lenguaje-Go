## Go Pointers

Go supports **pointers**, which hold the **memory address** of a value.

* The type `*T` represents a pointer to a value of type `T`.
* Its **zero value** is `nil`.

```go
var p *int
```

### Getting a Pointer

Use the `&` operator to get the memory address of a variable.

```go
i := 42
p = &i
```

### Using a Pointer

The `*` operator dereferences a pointer — it accesses the value stored at that address.

```go
fmt.Println(*p) // read i through pointer p
*p = 21         // set i through pointer p
```

This is called **dereferencing** or **indirection**.

> Unlike C, Go **does not allow pointer arithmetic**.

#### Example

```go
package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i         // point to i
    fmt.Println(*p) // read i through pointer
    *p = 21         // set i through pointer
    fmt.Println(i)  // new value of i

    p = &j          // point to j
    *p = *p / 37    // modify j through pointer
    fmt.Println(j)  // new value of j
}
```

---

## Structs

A **struct** is a collection of fields (data grouped together).

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
```

---

## Struct Fields

Access struct fields using the **dot (.) operator**.

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```

---

## Pointers to Structs

Struct fields can also be accessed via **struct pointers**.

Instead of writing `(*p).X`, Go allows shorthand access using `p.X`.

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v
    p.X = 1e9
    fmt.Println(v)
}
```

---

## Struct Literals

A **struct literal** creates and initializes a struct.

You can:

* Specify all fields in order.
* Specify only some fields by name (unordered).
* Use `&` to get a pointer to the new struct.

```go
package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}  // type: Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0, Y:0
    p  = &Vertex{1, 2} // type: *Vertex
)

func main() {
    fmt.Println(v1, p, v2, v3)
}
```

---

## Arrays

An **array** is a fixed-size sequence of elements of the same type.

```go
var a [10]int
```

Here, `a` is an array of 10 integers.

> Array length is part of its type, so arrays **cannot be resized** — but Go provides **slices** for more flexible use.

#### Example

```go
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"

    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
```

## What is a Pointer?

A **pointer** is a variable that stores the **memory address** of another variable.

Think of it like:

> A house address (pointer) that tells you where the actual item (value) is stored.

---

## Basic Example

```go
package main

import "fmt"

func main() {
    x := 10        // x is a normal integer variable
    p := &x        // p is a pointer to x (holds the address of x)

    fmt.Println("x =", x)     // value of x
    fmt.Println("p =", p)     // address of x
    fmt.Println("*p =", *p)   // value stored at that address
}
```

### Output (example)

```
x = 10
p = 0xc000014090
*p = 10
```

* `&x` → gives the **address** of `x`.
* `*p` → gets the **value** stored at that address.

---

## Changing Values via Pointers

If you modify a value using its pointer, it changes the original variable.

```go
*p = 20
fmt.Println(x)  // Output: 20
```

Because `p` points to `x`, updating `*p` means you’re updating `x`.

---

## Declaring Pointers Directly

```go
var p *int // p is a pointer to int, but it’s nil for now
fmt.Println(p) // Output: <nil>
```

You can assign it later:

```go
x := 42
p = &x
```

---

## Functions and Pointers

Go **passes arguments by value**, meaning it copies data.
To modify a variable inside a function, use a pointer.

```go
func update(num *int) {
    *num = *num + 10
}

func main() {
    x := 5
    update(&x)
    fmt.Println(x) // Output: 15
}
```

Using `&x` passes the **address**, not the value — letting the function modify `x`.

---

## Why Pointers Matter (Real Backend Use)

1. **Performance:** Avoid copying big structs or slices.
2. **Mutability:** Allow functions to modify shared data.
3. **Nil checks:** Optional fields in structs (similar to nullable references).

Example with a struct:

```go
type User struct {
    Name string
    Age  int
}

func updateUser(u *User) {
    u.Age += 1
}

func main() {
    user := User{Name: "Njox", Age: 21}
    updateUser(&user)
    fmt.Println(user.Age) // Output: 22
}
```

---

## Pointers and Nil

A pointer with no address points to `nil`.

```go
var p *int
if p == nil {
    fmt.Println("Pointer is nil")
}
```

You must allocate before using it:

```go
p = new(int)
*p = 100
fmt.Println(*p) // Output: 100
```

---

## Recap Summary

| Concept      | Symbol      | Meaning                     |
| ------------ | ----------- | --------------------------- |
| `&x`         | Address-of  | Get address of variable     |
| `*p`         | Dereference | Get or set value at pointer |
| `var p *int` | Declaration | p can point to an int       |
| `new(T)`     | Allocation  | Creates pointer to new T    |
| `nil`        | Zero value  | Uninitialized pointer       |

---

## Mini Challenge

Try this out 

```go
package main

import "fmt"

func main() {
    a := 5
    b := &a
    fmt.Println(a, *b)

    *b = 10
    fmt.Println(a, *b)

    c := new(int)
    *c = 15
    fmt.Println(*c)
}
```

Can you explain what prints at each line?

---

### Summary

| Symbol | Name                          | Meaning                                           | Example   | Description                                                |
| ------ | ----------------------------- | ------------------------------------------------- | --------- | ---------------------------------------------------------- |
| `&`    | **Address-of / Referencing**  | Gives you the **memory address** of a variable    | `p := &x` | “Point `p` to the address of `x`.”                         |
| `*`    | **Dereference / Indirection** | Accesses the **value** stored at a memory address | `*p = 10` | “Go to the address stored in `p` and set its value to 10.” |

---

### Example to See Both in Action

```go
x := 5
p := &x   // referencing
fmt.Println(p)   // prints memory address, e.g. 0xc0000180b0
fmt.Println(*p)  // dereferencing → prints 5

*p = 20          // modify the value at that address
fmt.Println(x)   // prints 20
```

---

### Visual Memory View

```
x:  [ 5 ]           ← value
p:  [ address of x ] → 0xc0000180b0
*p: dereferences p → gives 5
```

So:

* `&x` = “Where is x?”
* `*p` = “What’s stored there?”


🔥 Perfect timing bro — you’ve already mastered **structs**, so now let’s take it **one level higher:**
👉 **Pointers to Structs** — how Go lets you modify struct data directly without copying.

---

## 🧩 1. Why We Need Pointers to Structs

By default, **structs are value types**.
That means assigning or passing them **copies** the entire struct.

Example 👇

```go
type User struct {
    Name string
    Age  int
}

func main() {
    u1 := User{"Njox", 22}
    u2 := u1      // full copy
    u2.Age = 30

    fmt.Println(u1.Age) // 22 ❌ not changed
    fmt.Println(u2.Age) // 30
}
```

So `u2` is a **separate copy** — changing it doesn’t affect `u1`.

---

## 🧠 2. Fixing That with Pointers

If we want both variables to point to the **same memory**, we use a **pointer to the struct**.

```go
type User struct {
    Name string
    Age  int
}

func main() {
    u1 := User{"Njox", 22}
    u2 := &u1        // u2 is a pointer to u1

    u2.Age = 30      // modify through the pointer

    fmt.Println(u1.Age) // 30 ✅ changed
    fmt.Println(u2.Age) // 30 ✅ same memory
}
```

### 🧩 Explanation:

* `&u1` → gives the address of `u1`
* `u2` → stores that address (pointer)
* `u2.Age` → Go automatically dereferences the pointer for you (shortcut for `(*u2).Age`)

---

## ⚙️ 3. Creating a Pointer Struct Directly

You can create a struct pointer without an existing variable using `new()` or `&{}` syntax.

```go
u := &User{Name: "Lee", Age: 25}
fmt.Println(u)       // prints pointer address
fmt.Println(*u)      // prints full struct value
fmt.Println(u.Name)  // access field directly
```

✅ Go lets you use `u.Name` even though `u` is a pointer — it automatically dereferences for you.

---

## 🧱 4. Passing Pointer Structs to Functions

This is common in backend apps — you pass a pointer so the function modifies the original struct.

```go
func makeOlder(u *User) {
    u.Age++
}

func main() {
    user := User{"Njox", 22}
    makeOlder(&user)
    fmt.Println(user.Age) // 23 ✅ modified
}
```

* We pass `&user` (the address)
* The function uses `*User` (pointer receiver)
* The original struct changes directly

---

## 🧠 5. Pointer Shortcut Recap

| Expression     | Meaning                                          |
| -------------- | ------------------------------------------------ |
| `&x`           | Get the address of `x`                           |
| `*p`           | Get the value stored at pointer `p`              |
| `p := &User{}` | Pointer to new struct                            |
| `p.Name`       | Go auto-dereferences (you can use this directly) |

---

## ⚡ Example: Full Demo

```go
package main

import "fmt"

type Product struct {
    Name  string
    Price float64
}

func discount(p *Product, percent float64) {
    p.Price = p.Price - (p.Price * percent / 100)
}

func main() {
    prod := Product{"Laptop", 1000}
    discount(&prod, 10)
    fmt.Println(prod) // {Laptop 900}
}
```

✅ The function modifies the original `prod` because it receives a **pointer**.

---

## 🔥 Summary

| Concept        | Without Pointer         | With Pointer           |
| -------------- | ----------------------- | ---------------------- |
| Memory         | Copies struct           | References same struct |
| Function edits | Doesn’t affect original | Updates original       |
| Common use     | Reading data            | Updating data          |

---

So now you can:

* Create structs ✅
* Use pointers to structs ✅
* Modify data directly ✅

---

Next step (and it’s where Go shines in backend dev):
👉 **Pointer Receivers in Struct Methods** — attaching functions to structs so they can mutate data directly (like classes in OOP).

Do you want to move to that next?
