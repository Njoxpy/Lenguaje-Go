# 1. What are Methods in Go?

In Go, a **method** is a **function with a receiver**. The receiver can be:

- **Value receiver**: method works on a **copy** of the struct
- **Pointer receiver**: method works on the **original struct**

Syntax:

```go
func (receiver Type) MethodName(params) returnType {
    // code
}
```

- `receiver` is usually a variable representing the struct instance.
- `Type` is the struct type (or any custom type).

---

### Example — Basic Method

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

// Value receiver
func (u User) Greet() string {
    return "Hello, " + u.Name
}

// Pointer receiver
func (u *User) HaveBirthday() {
    u.Age += 1
}

func main() {
    user := User{Name: "Alice", Age: 25}

    fmt.Println(user.Greet()) // Hello, Alice

    user.HaveBirthday()
    fmt.Println(user.Age)     // 26
}
```

**Key points:**

- Value receiver: `Greet()` does not modify `user`
- Pointer receiver: `HaveBirthday()` **modifies the original struct**

---

# 2. Real-World Uses of Methods

Methods are extremely useful because they **associate behavior with data**. This is similar to OOP, but Go keeps it simple.

### Examples in real projects:

1. **Backend / Business Logic**

   - User management: `user.ResetPassword()`, `user.IsActive()`
   - Banking system: `account.Deposit()`, `account.Withdraw()`
   - Inventory system: `product.IncreaseStock(amount)`

2. **Encapsulation**

   - Struct fields can be private (`lowercase`) and methods control access:

     ```go
     type Account struct {
         balance int
     }

     func (a *Account) Deposit(amount int) {
         if amount > 0 {
             a.balance += amount
         }
     }

     func (a *Account) Balance() int {
         return a.balance
     }
     ```

3. **Stateful Closures**

   - Methods can return closures that capture struct state:

     ```go
     func (a *Account) TransactionLogger() func(amount int) int {
         return func(amount int) int {
             a.balance += amount
             return a.balance
         }
     }
     ```

4. **Collections of Methods**

   - Methods allow struct-based APIs in backend projects:

     - `Cart.AddItem()`, `Cart.RemoveItem()`, `Cart.Total()`
     - Makes code more readable and maintainable.

---

# 3. Importance of Methods in Backend Projects

1. **Organization of code**

   - Methods tie **data + behavior together**, reducing global state
   - Easier to maintain and test

2. **Encapsulation and Safety**

   - Pointer receivers allow controlled mutation of structs
   - Fields can remain private, exposed via methods only

3. **Reusability**

   - One method can be used across the backend wherever that struct is used

4. **Readability**

   - Instead of `UpdateUserAge(&user, 25)`, you write `user.UpdateAge(25)` — more intuitive

5. **Testing**

   - Methods can be tested independently
   - You can combine methods with closures for stateful testing

---

# 4. Disadvantages / Things to Watch Out For

1. **Copying large structs**

   - Using value receivers on large structs **copies the whole struct**, which is inefficient
   - Solution: use pointer receivers for large structs

2. **Mixing pointer and value receivers**

   - Can lead to confusion
   - Rule of thumb: use pointer receiver if method **modifies the struct** or the struct is large; otherwise, value receiver is fine

3. **Interface complexity**

   - Too many small methods can lead to large interfaces
   - Keep interfaces minimal (Go encourages small, composable interfaces)

4. **Concurrent access**

   - Methods do **not automatically handle concurrency**
   - If multiple goroutines call pointer-receiver methods that mutate the struct, you must use `sync.Mutex` or `sync.RWMutex`

---

# 5. Summary

| Feature                  | Value Receiver                        | Pointer Receiver                    |
| ------------------------ | ------------------------------------- | ----------------------------------- |
| Modifies original struct | No                                    | Yes                                 |
| Struct copying cost      | Copies entire struct                  | No copy, uses pointer               |
| Use case                 | Read-only methods                     | Methods that mutate or large struct |
| Safety in concurrency    | Same as pointer (must lock if needed) | Same, must lock for safety          |

**Backend importance:**

- Methods are core to building maintainable, encapsulated, and testable business logic.
- They reduce global state and allow you to write reusable APIs.

---

If you want, I can **give a set of step-by-step method exercises** that include:

- Value vs pointer receiver
- Methods returning closures
- Using methods in mini backend-style structs like `User`, `Account`, `Cart`
