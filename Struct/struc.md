In **Go**, a **struct** is a **composite data type** that groups together variables (called **fields**) under a single name. Structs are like **blueprints for objects**, but without the full object-oriented machinery—perfect for organizing data.

Think of a struct as a **custom type** for grouping related data.

---

## 🔹 Basic Syntax

```go
package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of the struct
    p1 := Person{Name: "Alice", Age: 25}

    fmt.Println(p1)
    fmt.Println("Name:", p1.Name)
    fmt.Println("Age:", p1.Age)
}
```

### ✅ Output:

```
{Alice 25}
Name: Alice
Age: 25
```

---

## 🔹 Key Points About Structs

1. **Fields can be of any type**

```go
type Book struct {
    Title  string
    Author string
    Pages  int
    Price  float64
}
```

2. **Struct values are accessed with dot `.` notation**

```go
b := Book{Title: "Go in Action", Pages: 300}
fmt.Println(b.Title)
```

3. **Structs can be initialized in multiple ways**

```go
// 1. Using field names
p := Person{Name: "Bob", Age: 30}

// 2. Without field names (order matters)
p := Person{"Bob", 30}

// 3. Pointer to struct
p := &Person{Name: "Carol", Age: 22}
p.Age = 23 // modifies the original struct
```

---

## 🔹 Structs With Methods

You can attach **methods** to structs in Go:

```go
type Rectangle struct {
    Width, Height float64
}

// Method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area())
}
```

✅ Output:

```
Area: 50
```

> `r Rectangle` is called a **receiver**, similar to `this` in other languages.

---

## 🔹 Structs vs Closures

* **Struct** → holds **data** and optionally **methods**
* **Closure** → holds **data** in captured variables and **functions**

Both can maintain state, but **structs are more explicit and typed**.

---

## 🔹 Real-world Example

```go
type BankAccount struct {
    Owner   string
    Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
    a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) {
    if amount <= a.Balance {
        a.Balance -= amount
    }
}

func main() {
    acc := &BankAccount{Owner: "Alice"}
    acc.Deposit(500)
    acc.Withdraw(200)
    fmt.Println("Balance:", acc.Balance)
}
```

✅ Output:

```
Balance: 300
```

* Here, `BankAccount` **encapsulates state**
* Methods manipulate that state safely

---

