🔥 Let’s go bro — time to *build your muscle memory* 💪
Here are some **pointer-to-struct exercises**, from beginner to pro-level, all realistic for backend Go devs.

---

## 🧩 **LEVEL 1 — Basics (Pointer & Struct Access)**

### 🧠 Exercise 1: Simple Pointer Update

Create a struct called `Student` with fields `Name` and `Marks`.

1. Create one variable `s1 := Student{"Ali", 70}`
2. Make a pointer to it (`s2 := &s1`)
3. Update marks to 90 using the pointer.
4. Print both `s1.Marks` and `s2.Marks` — they should be the same.

👉 **Goal:** understand how pointers share the same memory.

---

### 🧠 Exercise 2: Address Comparison

Create a struct `Car` with fields `Model` and `Year`.

1. Make one car variable.
2. Create a pointer to it.
3. Print `&car` and the pointer variable itself.
   They should show the same address — prove it in the output.

👉 **Goal:** verify that a pointer stores the address of the struct.

---

## ⚙️ **LEVEL 2 — Functions & Pointers**

### 🧠 Exercise 3: Passing Pointers to Functions

Make a struct:

```go
type Account struct {
    Balance float64
}
```

Write a function:

```go
func deposit(a *Account, amount float64)
```

It should add `amount` to the account’s balance.

1. Create `acc := Account{Balance: 1000}`
2. Call `deposit(&acc, 500)`
3. Print balance → should be 1500.

👉 **Goal:** practice passing struct pointers to modify data in functions.

---

### 🧠 Exercise 4: Withdraw Function

Using the same `Account` struct,
write a function `withdraw(a *Account, amount float64)`
that subtracts from balance only if there’s enough money.
Print “Insufficient funds” otherwise.

👉 **Goal:** conditionally modify struct fields via pointer.

---

## 💼 **LEVEL 3 — Realistic Struct + Pointer Use**

### 🧠 Exercise 5: Inventory System

Define a struct:

```go
type Product struct {
    Name  string
    Stock int
}
```

Create:

* A slice of 3 products
* A function `restock(p *Product, n int)` that adds `n` to stock
* Loop through all products, restock each by 10

Print before and after.

👉 **Goal:** work with slices + pointers to structs inside loops.

---

### 🧠 Exercise 6: User Profile Updater

Define:

```go
type User struct {
    Name  string
    Email string
}
```

1. Create a pointer function `updateEmail(u *User, newEmail string)`
2. Create a few users.
3. Call the function to change one’s email.

👉 **Goal:** simulate updating user data in a backend-style function.

---

## 🔥 BONUS (For 10/10)

### 🧠 Exercise 7: Nested Pointer Update

Define:

```go
type Address struct {
    City string
}
type Employee struct {
    Name string
    Addr Address
}
```

Try to update the employee’s city *via pointer* to the whole struct.

👉 **Hint:**

```go
func updateCity(e *Employee, newCity string) {
    e.Addr.City = newCity
}
```

---

### 🧩 Optional Challenge (Hardcore)

Make a simple mini system:

```go
type Customer struct {
    Name string
    Points int
}

func addPoints(c *Customer, pts int) {
    // ...
}

func main() {
    // multiple customers
}
```

Use pointers so that your functions can **modify points directly**, not copies.

---

Would you like me to turn these exercises into a **mini Go practice sheet** (with file templates and expected outputs) so you can code them one by one?
