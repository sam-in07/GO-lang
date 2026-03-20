In Go, a **receiver function** (more commonly called a **method**) is a function that is **associated with a specific type**, usually a **struct**. The receiver allows the function to **access or modify the fields** of that type, similar to how `this` works in other languages like Java or C++.

---

## 🔹 Syntax

```go
func (receiver Type) MethodName(parameters) returnType {
    // method body
}
```

* `receiver` → a variable name you choose (conventionally one or two letters)
* `Type` → the type this method is associated with
* Can be **value receiver** or **pointer receiver**

---

## 🔹 Value Receiver vs Pointer Receiver

### 1. **Value Receiver**

The method gets a **copy** of the struct. Changes inside the method **don’t affect the original struct**.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Value receiver
func (p Person) CelebrateBirthday() {
    p.Age += 1
    fmt.Println("Happy Birthday,", p.Name, "Age:", p.Age)
}

func main() {
    p1 := Person{Name: "Alice", Age: 25}
    p1.CelebrateBirthday() // Age increases only inside the method
    fmt.Println("Outside Age:", p1.Age) // Still 25
}
```

✅ Output:

```
Happy Birthday, Alice Age: 26
Outside Age: 25
```

---

### 2. **Pointer Receiver**

The method gets a **pointer to the struct**, so changes **affect the original struct**.

```go
// Pointer receiver
func (p *Person) CelebrateBirthdayPointer() {
    p.Age += 1
    fmt.Println("Happy Birthday,", p.Name, "Age:", p.Age)
}

func main() {
    p2 := Person{Name: "Bob", Age: 30}
    p2.CelebrateBirthdayPointer()
    fmt.Println("Outside Age:", p2.Age) // Now Age = 31
}
```

✅ Output:

```
Happy Birthday, Bob Age: 31
Outside Age: 31
```

---

## 🔹 Why Use Receiver Functions?

1. **Encapsulation** → Keep data and behavior together
2. **Methods on structs** → Makes code more readable
3. **Pointer receivers** → Efficient for large structs or if you need to modify fields
4. **Interface implementation** → Only methods can satisfy interfaces

---

## 🔹 Example: Bank Account

```go
type BankAccount struct {
    Owner   string
    Balance float64
}

// Deposit method
func (a *BankAccount) Deposit(amount float64) {
    a.Balance += amount
}

// Withdraw method
func (a *BankAccount) Withdraw(amount float64) {
    if amount <= a.Balance {
        a.Balance -= amount
    }
}

func main() {
    account := &BankAccount{Owner: "Alice"}
    account.Deposit(500)
    account.Withdraw(200)
    fmt.Println("Balance:", account.Balance)
}
```

✅ Output:

```
Balance: 300
```

* `Deposit` and `Withdraw` are **receiver functions** attached to `BankAccount`
* They can **modify the account's balance** directly because of the pointer receiver

---

