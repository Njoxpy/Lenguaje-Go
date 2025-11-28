## 1. Slices are Reference Types

Remember: **slices point to an underlying array**. Modifying a slice can affect the original array.

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // [2 3 4]
slice[0] = 20
fmt.Println(arr)  // [1 20 3 4 5]
```

Changing `slice[0]` also changes `arr[1]` because the slice shares the same array.

---

## 2. Slice Length and Capacity

Every slice has:

- **Length** (`len`) → number of accessible elements
- **Capacity** (`cap`) → max elements before a new array is allocated

```go
s := []int{1, 2, 3, 4, 5}
fmt.Println(len(s)) // 5
fmt.Println(cap(s)) // 5

s = s[:3]           // slice first 3 elements
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 5 (capacity is still based on underlying array)
```

---

## 3. Appending to a Slice

Use `append()` to add elements. If capacity is exceeded, Go allocates a **new underlying array**.

```go
s := []int{1, 2, 3}
fmt.Println(len(s), cap(s)) // 3 3

s = append(s, 4)
fmt.Println(s)               // [1 2 3 4]
fmt.Println(len(s), cap(s))  // 4 6 (capacity increased automatically)
```

> Important: always assign back to the slice (`s = append(s, ...)`) because append may return a **new slice**.

---

## 4. Copying Slices

Sometimes you want to **avoid modifying the original slice**.

```go
original := []int{1, 2, 3}
copySlice := make([]int, len(original))
copy(copySlice, original)

copySlice[0] = 100
fmt.Println(original)  // [1 2 3]
fmt.Println(copySlice) // [100 2 3]
```

`copy()` copies elements from one slice to another.

---

## 5. Slicing a Slice

Slices can be sliced multiple times:

```go
s := []int{10, 20, 30, 40, 50}
s1 := s[1:4] // [20 30 40]
s2 := s[:3]  // [10 20 30]
s3 := s[2:]  // [30 40 50]

fmt.Println(s1, s2, s3)
```

---

## 6. Common Slice Interview Questions

- What happens if you append beyond the slice’s capacity?
- How do slices differ from arrays in memory?
- Show how modifying a slice can affect its underlying array.
- Write a function that removes an element at a given index from a slice **in place**.

---

## Mini Challenge

1. Create a slice of integers: `[10, 20, 30, 40]`
2. Append `50` and `60`
3. Slice it to `[20, 30, 40, 50]`
4. Modify the first element of the new slice to `99`
5. Print both the original and new slice — see what changed

---

## 1. `make` Syntax

```go
slice := make([]Type, length, capacity)
```

- `Type` → the type of elements (`int`, `string`, etc.)
- `length` → number of elements initially accessible
- `capacity` → maximum number of elements before Go allocates a new underlying array

> If you omit `capacity`, it defaults to `length`.

---

## 2. Example with Integers

```go
numbers := make([]int, 3, 5) // length = 3, capacity = 5
fmt.Println(numbers)         // [0 0 0] → zero values
fmt.Println(len(numbers))    // 3
fmt.Println(cap(numbers))    // 5
```

- `numbers` has space for **5 elements**, but currently only **3 are accessible**.

---

## 3. Example with Strings

```go
names := make([]string, 2)
fmt.Println(names)      // ["" ""] → zero value for strings
fmt.Println(len(names)) // 2
fmt.Println(cap(names)) // 2 (capacity defaults to length)
```

---

## 4. Appending to a Slice Created with `make`

```go
nums := make([]int, 2, 4)
nums[0] = 10
nums[1] = 20

nums = append(nums, 30, 40) // adding elements within capacity
fmt.Println(nums)            // [10 20 30 40]
fmt.Println(len(nums))       // 4
fmt.Println(cap(nums))       // 4

nums = append(nums, 50)      // exceeds capacity, Go allocates a new array
fmt.Println(nums)            // [10 20 30 40 50]
fmt.Println(len(nums))       // 5
fmt.Println(cap(nums))       // 8 (capacity doubled automatically)
```

---

### Key Points

- `make` creates a slice with **pre-allocated memory**, which is useful for performance if you know the approximate size.
- Length ≤ Capacity
- `append` can grow the slice beyond capacity; Go handles reallocation automatically.
- Uninitialized elements are set to the **zero value** of their type (`0` for numbers, `""` for strings, `false` for booleans).

---

# BEGINNER

**1. Create a slice with make**

- Create a slice of integers with **length 3** and **capacity 5**.
- Assign values 10, 20, 30 to the slice.
- Print the slice, its length, and capacity.

**2. Append elements**

- Append 40 and 50 to the slice from exercise 1.
- Print the slice, its length, and capacity.
- Append 60 and observe what happens to the capacity.

**3. Zero values**

- Create a slice of strings with length 4 using `make`.
- Print the slice and explain the output.

---

# INTERMEDIATE

**4. Slicing a slice**

- Create a slice `nums := make([]int, 5, 10)` and assign values 1 to 5.
- Slice it to get `[2,3,4]` using slicing syntax.
- Modify the first element of the new slice to 99.
- Print both the original slice and the new slice. What changed and why?

**5. Function that modifies a slice**

- Write a function `DoubleValues(nums []int)` that doubles each element of a slice.
- Pass a slice created with `make` to it.
- Print the slice before and after the function call.

**6. Copying slices**

- Create a slice with `make` and copy it to a new slice using `copy()`.
- Modify the new slice and print both slices. Does the original change? Explain.

---

# ADVANCED

**7. Removing an element in-place**

- Create a slice of integers with length 5 using `make`.
- Write a function `RemoveIndex(s []int, index int) []int` that removes an element at a given index without creating a new slice.
- Test it by removing the 2nd element. Print the resulting slice.

**8. Slice growth and capacity**

- Create a slice `s := make([]int, 2, 4)` and append elements to exceed the capacity.
- Print the slice, its length, and capacity after each append.
- Explain how Go handles slice growth internally.

**9. Multi-dimensional slice with make**

- Create a 2D slice (matrix) with 3 rows and 4 columns using `make`.
- Assign values such that `matrix[i][j] = i*4 + j + 1`.
- Print the matrix in a readable format.

**10. Combined pointer + slice exercise**

- Create a struct `Product` with fields `Name` and `Stock`.
- Create a slice of `Product` pointers using `make`.
- Write a function `Restock(products []*Product, n int)` that increases stock for all products.
- Test by printing before and after restocking.

---
