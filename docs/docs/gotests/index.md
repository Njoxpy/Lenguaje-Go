# GO TESTS ROADMAP

### Style: identical to _Learn Go With Tests_

### Coverage: all topics you learned

### Structure: multi-folder

### Difficulty: basic (but real)

---

# 1. Project Folder Structure

Create this structure:

```
go-basics-tests/
    functions/
        add.go
        add_test.go
    variables/
        convert.go
        convert_test.go
    flow/
        loops.go
        loops_test.go
    slices/
        slice_ops.go
        slice_ops_test.go
    maps/
        map_ops.go
        map_ops_test.go
    structs/
        user.go
        user_test.go
    pointers/
        counter.go
        counter_test.go
    closures/
        gen.go
        gen_test.go
```

Each folder covers a topic.

---

# 2. What to Test (Guided Checklist)

Below are the **required tests**, but written as **specifications**.
You will implement them yourself.

---

# functions/

### Test 1: Add two integers

- Write a function `Add(a, b int) int`
- Test cases:

  - positive numbers
  - negative numbers
  - zero

### Test 2: Repeat a string N times

- Function: `Repeat(s string, n int) string`
- Table-driven test.

---

# variables/

### Test 1: Celsius to Fahrenheit conversion

- Function: `CtoF(c float64) float64`
- Test:

  - known input/output pairs.

### Test 2: Boolean toggler

- Function: `Toggle(b bool) bool`
- Write two test cases.

---

# flow/

### Test 1: Sum numbers 1 to N

- Function: `SumTo(n int) int`
- Test:

  - 1
  - 5
  - 10

### Test 2: Find the first even number in a slice

- Function: `FirstEven([]int) (int, bool)`
- Test:

  - slice with evens
  - slice with no evens

---

# slices/

### Test 1: AppendSafe

- Function: `AppendSafe(s []int, v int) []int`
- Behavior:

  - original slice must not mutate
  - new slice contains old values + v

Test:

- original and returned slice should differ.

### Test 2: SumSlice

- Function: `SumSlice([]int) int`
- Table test.

---

# maps/

### Test 1: Basic map get

- Function: `GetAge(m map[string]int, name string) (int, bool)`
- Test for:

  - existing key
  - missing key

### Test 2: UpdateStock

- Function: `UpdateStock(m map[string]int, item string, qty int)`
- Test:

  - increasing quantity
  - creating new key
  - overwriting existing

---

# structs/

### Define:

```
type Book struct {
    Title  string
    Author string
    Price  float64
}
```

### Test 1: DiscountPrice(b Book, percent float64) float64

- Test:

  - 10% discount
  - 0% discount
  - 100% discount

### Test 2: RenameAuthor(b \*Book, newAuthor string)

- Test:

  - uses pointer (value should change)

---

# pointers/

### Define struct Counter with field `n int`

### Test 1: Increment

- Method: `func (c *Counter) Inc()`
- Test:

  - call Inc twice → value is 2
  - pointer receiver required

### Test 2: Reset

- Method: `func (c *Counter) Reset()`
- Test:

  - sets counter to 0

---

# closures/

### Test 1: NewIDGenerator

Function returns a closure:

```
func NewIDGen() func() int
```

Every call increases the ID.

Test:

- call 3 times → 1,2,3

### Test 2: Adder

- Function: `NewAdder(start int) func(int) int`
- Test:

  - closure maintains internal state.

---

# 3. Rules (important)

1. Use _Learn Go With Tests_ style:

   - `testing.T`
   - subtests (`t.Run`)
   - table-driven tests
   - helper functions (if needed)

2. Create implementation only **after** test fails.

3. Run tests all the time:

   ```
   go test ./...
   ```

4. Don’t copy online answers.
   Make mistakes. Fix them. That’s how you become badass.
