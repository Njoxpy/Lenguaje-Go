# Beginner → Intermediate → Advanced Go Projects

### **1. Command-line Calculator (Beginner)**

- Features:

  - Supports addition, subtraction, multiplication, division
  - Uses functions, function values, and method receivers

- What you’ll learn:

  - Function values, method receivers, closures for custom operations

- Extensions:

  - Add factorial, power, and Fibonacci operations

---

### **2. Fibonacci & Prime Number Generator (Beginner → Intermediate)**

- Features:

  - Generate first N Fibonacci numbers
  - Check if a number is prime
  - Generate all primes up to N

- What you’ll learn:

  - Loops, slices, maps, closures, function values
  - Memoization with maps

- Extensions:

  - Write a **memoized Fibonacci** closure
  - Count prime frequencies in a given range

---

### **3. Stateful Counter with Closures (Intermediate)**

- Features:

  - Create multiple counters with different step sizes
  - Use closures to increment

- What you’ll learn:

  - Structs, pointer receivers, closures
  - Encapsulation of state

- Extensions:

  - Store counters in a slice and iterate them dynamically

---

### **4. Word Frequency Analyzer (Intermediate)**

- Features:

  - Reads a text file (or a string)
  - Counts frequency of each word

- What you’ll learn:

  - Maps, slices, strings, functions
  - Function as argument (e.g., pass a filter function)

- Extensions:

  - Return top N words
  - Use closures to filter stopwords

---

### **5. Stack & Queue Implementations (Intermediate)**

- Features:

  - Implement stack with push/pop
  - Implement queue with enqueue/dequeue
  - Use slices and pointers

- What you’ll learn:

  - Pointers, slices, methods

- Extensions:

  - Implement a stack of structs (e.g., Task struct)

---

### **6. Matrix Operations (Intermediate → Advanced)**

- Features:

  - Create 2D slices for matrices
  - Implement addition, multiplication, and transpose

- What you’ll learn:

  - Multi-dimensional slices, loops, methods

- Extensions:

  - Implement determinant calculation using recursion

---

### **7. Simple Banking System (Advanced)**

- Features:

  - Account struct with balance
  - Methods: Deposit, Withdraw, Transfer
  - Closures for transaction logs or statement generation

- What you’ll learn:

  - Structs, pointer receivers, closures, maps for accounts

- Extensions:

  - Slice of pointers to accounts
  - Track transaction history per account using slices

---

### **8. Graph Algorithms (Advanced)**

- Features:

  - Represent a graph using map[string][]string (adjacency list)
  - Implement BFS and DFS

- What you’ll learn:

  - Maps, slices, recursion, structs

- Extensions:

  - Track visited nodes with maps
  - Implement shortest path (Dijkstra’s basic version)

---

### **9. Memoization / Dynamic Programming Library (Advanced)**

- Features:

  - Create reusable closure-based memoization for expensive functions
  - Examples: Fibonacci, factorial, combinatorics

- What you’ll learn:

  - Closures, maps, function values, recursion

- Extensions:

  - Generic memoization function that works for different types (Go 1.18+ generics)

---

### **10. CLI Number Puzzle Solver (Advanced)**

- Features:

  - Solve small number puzzles (like 8-puzzle or Sudoku mini version)
  - Represent the board as a slice of slices
  - Use functions and closures for move validation

- What you’ll learn:

  - Slices, structs, maps, methods, recursion

- Extensions:

  - Track moves using slice of structs
  - Implement simple backtracking algorithm
