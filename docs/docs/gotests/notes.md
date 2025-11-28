### 1. Test Naming Pattern

- All test functions start with `Test…` → `TestAdd`, `TestHello`.
- Test files always end with `_test.go`.

**Pattern to notice:** Go automatically detects `Test*` functions. This is consistent across all packages.

---

### 2. Assertion Pattern

- Store output in `got`.
- Define expected output as `want`.
- Compare: `if got != want { t.Errorf(...) }`.

**Pattern:**

```go
got := Function(input)
want := expected
if got != want {
    t.Errorf("got %v want %v", got, want)
}
```

Everything you test in Go follows this basic template, whether it’s strings, ints, slices, or structs.

---

### 3. Helper Pattern

- If checks are repeated, extract to helper function:

```go
func assertEqual(t testing.TB, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}
```

- Use `t.Helper()` so the test reports the **call site** instead of the helper.

---

### 4. Subtest Pattern

- For multiple scenarios:

```go
t.Run("description", func(t *testing.T) {
    got := Function(input)
    want := expected
    if got != want {
        t.Errorf("got %v want %v", got, want)
    }
})
```

- You’ll see this pattern a lot when testing multiple inputs.

---

### 5. Common types & verbs

- `%d` → integers
- `%q` → strings (quotes)
- `%v` → any type (generic, structs, slices)

**Pattern:** Always match type with the correct verb for readability.

---

### 6. TDD Cycle Pattern

1. Write a **failing test**.
2. Make it **compile**.
3. Write **minimum code** to pass.
4. **Refactor** safely with tests.

This cycle is **the same pattern** for any feature you develop in Go.
