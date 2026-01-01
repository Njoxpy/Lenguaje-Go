# **1️⃣ Exercise: Users and Roles**

- Create a `User` struct with `Name`, `Age`, `Role`
- Implement `UserActions` interface (`Login()`, `Logout()`)
- Add a method `IsAdmin()` that **returns true if Role is "Admin"**
- Write a function `PerformAdminAction(u UserActions)` that only runs `Logout()` if the user is admin

**Goal:** Practice **interfaces + methods + role-based logic**

---

# **2️⃣ Exercise: Products and Inventory**

- Create `Product` struct with `ProductName`, `Quantity`, `Price`
- Implement `ProductActions` interface (`GetPrice()`, `AddProduct()`)
- Write a function `Restock(p ProductActions, n int)` that increases quantity

**Goal:** Practice **interfaces + struct manipulation**

---

# **3️⃣ Exercise: Generic Repository**

- Create a generic `Repository[T any]`
- Add methods:

  - `Add(item T)`
  - `GetAll() []T`
  - `FindByName(name string) (T, error)` (return error if not found)

**Goal:** Practice **generics + type-safe storage + error handling**

---

# **4️⃣ Exercise: Validation with custom errors**

- Write `ValidateUser(u User) error` → returns error if name or age invalid
- Write `ValidateProduct(p Product) error` → returns error if quantity ≤ 0 or price ≤ 0

**Goal:** Practice **error interface + custom types**

---

# **5️⃣ Combine everything**

- Create repositories for users and products
- Validate them before adding to repo
- Use functions that accept **interfaces** to perform actions on all users/products
- Handle errors gracefully

**Goal:** This is a **mini backend simulation** with Go idiomatic patterns

---

# 🔑 **Tips while practicing**

1. Start **small**: just `User` struct and `UserActions` interface
2. Add **methods** and test them individually
3. Implement the **generic repository** after interfaces are solid
4. Add **error handling** last, to make sure validations work
