### Hello, 世界

Welcome to the **Go Programming Language Tour** — an interactive guide designed to help you learn Go step by step.

---

### Navigating the Tour

The tour is divided into several modules.
You can access them by clicking **“A Tour of Go”** at the top-left of the page.

To view the **table of contents** at any time, click the **menu** icon in the top-right corner.

Each module contains a series of **slides** and **hands-on exercises**.
Use the navigation options below to move through the tour:

- **Previous** or **PageUp** → Go to the previous page
- **Next** or **PageDown** → Move to the next page

---

### Running and Editing Code

The tour is interactive — try it now!
Click **Run** (or press **Shift + Enter**) to compile and execute the program on a remote server.
The output will appear directly below the code editor.

Each example demonstrates a specific feature or concept in Go.
These programs are designed as **starting points** for your own experimentation — feel free to **edit** the code and **run it again**.

You can also:

- Click **Format** (shortcut: **Ctrl + Enter**) to format your code using the `gofmt` tool.
- Toggle syntax highlighting by clicking the **Syntax** button.

When you’re ready to continue, click the **right arrow** below or press **PageDown**.

---

### Example Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

---

### Go in Your Language

The tour is available in multiple languages:

- Português do Brasil
- Català
- 中文（简体）
- Česky
- Bahasa Indonesia
- 日本語
- Polski
- Español
- ภาษาไทย
- Українською

Click **Next** or press **PageDown** to continue.

---

### Go Offline (Optional)

You can also take this tour **offline**.
The offline version runs locally on your computer and doesn’t require internet access.

To set it up:

1. **Install Go** on your system.
2. Run the following command:

   ```bash
   go install golang.org/x/website/tour@latest
   ```

   This installs the tour binary in your `GOPATH/bin` directory.

Then, launch the tour program to open a browser window with your **local version** of the tour.

(You can, of course, continue using this online version instead.)

---

### The Go Playground

This tour is powered by the **Go Playground**, a web service hosted on **golang.org**.

Here’s how it works:

- The Playground receives your Go program, compiles, links, and runs it securely inside a sandbox.
- It then returns the output to your browser.

#### Playground Limitations:

- The time inside the playground always starts at **2009-11-10 23:00:00 UTC** — a fixed reference that helps ensure consistent program output.
- Programs have **execution time**, **CPU**, and **memory limits**.
- Programs **cannot access external network hosts**.
- The Playground always uses the **latest stable release** of Go.

To learn more, read _Inside the Go Playground_.

---

### Example

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Welcome to the playground!")
    fmt.Println("The time is", time.Now())
}
```

---

### Congratulations

You’ve completed the **first module** of the Go Tour!

Click **A Tour of Go** at the top to explore more topics,
or proceed directly to the **next lesson** to continue your journey.
