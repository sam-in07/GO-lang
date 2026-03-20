In **Go**, a **pointer** is a variable that **stores the memory address of another variable**. Pointers let you **directly access or modify a value in memory**, which is very useful for efficiency and for functions that need to modify their arguments.

---

## 🔹 Basic Pointer Syntax

```go
var ptr *int  // ptr is a pointer to an int
```

* `*Type` → type of the value the pointer points to
* `&` → **address-of operator**
* `*` → **dereference operator** (access the value at the address)

---

## 🔹 Example: Basic Pointer

```go
package main

import "fmt"

func main() {
    x := 42
    var p *int = &x // p stores the address of x

    fmt.Println("x =", x)      // 42
    fmt.Println("p =", p)      // memory address of x
    fmt.Println("*p =", *p)    // 42 (value at address)

    *p = 100 // change x through the pointer
    fmt.Println("x after *p =", x) // 100
}
```

✅ Output:

```
x = 42
p = 0xc0000140a0   // example memory address
*p = 42
x after *p = 100
```

---

## 🔹 Pointers and Functions

Pointers are often used to **modify variables inside a function**.

```go
func increment(val *int) {
    *val = *val + 1
}

func main() {
    a := 10
    increment(&a)
    fmt.Println(a) // 11
}
```

> Without pointers, `increment(a)` would only modify a copy, not the original variable.

---

## 🔹 Pointer to Struct

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := &Person{Name: "Alice", Age: 25} // pointer to struct
    fmt.Println(p.Age) // 25

    p.Age = 30 // can modify through pointer
    fmt.Println(p.Age) // 30
}
```

> Go lets you **automatically dereference** pointers to structs when accessing fields.

---

## 🔹 Key Points

1. `&` → get the **address** of a variable
2. `*` → get or set the **value at an address**
3. Pointers can improve performance (pass large structs by pointer)
4. Pointers are required to **modify values inside functions**

---

### 🔹 Quick Tip: Pointer vs Value Receiver in Structs

* **Value receiver**: gets a copy → changes **don’t affect** the original
* **Pointer receiver**: gets a pointer → changes **modify** the original

Example:

```go
func (p *Person) Birthday() {
    p.Age++
}
```

Here, `*Person` pointer allows the method to **update the Age field**.

---