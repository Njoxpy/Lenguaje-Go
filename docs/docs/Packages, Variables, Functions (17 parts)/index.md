## Packages in Go

Every Go program is composed of **packages**.
Execution always begins in the **`main`** package.

This example uses the packages with import paths **`"fmt"`** and **`"math/rand"`**.

By convention, the **package name** matches the **last element** of its import path.
For example, files in the `"math/rand"` package start with the declaration:

```go
package rand
```

### Example:

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}
```

---

## Imports

In Go, imports can be grouped into a **factored import statement**, enclosed in parentheses.
This is the preferred style for readability and organization.

### Example (factored import):

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

**Output:**

```
Now you have 2.6457513110645907 problems.
Program exited.
```

Alternatively, you can use multiple single-line import statements, though it’s less common:

```go
import "fmt"
import "math"
```

---

## Exported Names

In Go, **a name is exported** if it begins with a **capital letter**.

For example:

- `Pizza` — exported
- `Pi` — exported from the `math` package
- `pizza` and `pi` — _not exported_ (private to the package)

When you import a package, you can access only its **exported names**.
Unexported names (those starting with a lowercase letter) are **not accessible** from outside the package.

### Example (with error):

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.pi) // Error: cannot refer to unexported name math.pi
}
```

### Fixed version:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi) // Works correctly
}
```
