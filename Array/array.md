In **Go**, an **array** is a **fixed-size sequence of elements of the same type**. Unlike slices, the **size of an array is part of its type**, so arrays in Go are less flexible, but they are very useful for understanding Go’s type system.

---

## 🔹 Basic Syntax

```go
var arr [size]Type
```

* `size` → number of elements
* `Type` → type of each element

### Example:

```go
package main

import "fmt"

func main() {
    var numbers [5]int // array of 5 integers

    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30

    fmt.Println(numbers)     // [10 20 30 0 0]
    fmt.Println(numbers[2])  // 30
    fmt.Println(len(numbers)) // 5
}
```

✅ Output:

```
[10 20 30 0 0]
30
5
```

---

## 🔹 Array Initialization

1. **All zeros (default)**

```go
var arr [3]int // [0 0 0]
```

2. **With values**

```go
arr := [3]int{1, 2, 3}  // [1 2 3]
```

3. **Let Go count the length**

```go
arr := [...]int{4, 5, 6, 7} // length 4
```

---

## 🔹 Looping Through Arrays

```go
arr := [3]string{"Go", "Python", "Java"}

for i := 0; i < len(arr); i++ {
    fmt.Println(i, arr[i])
}

// Using range
for index, value := range arr {
    fmt.Println(index, value)
}
```

✅ Output:

```
0 Go
1 Python
2 Java
```

---

## 🔹 Multidimensional Arrays

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

fmt.Println(matrix)
fmt.Println(matrix[1][2]) // 6
```

---

## 🔹 Key Points About Arrays

1. Arrays **have fixed size**
2. Size is **part of the type**:

```go
var a [3]int
var b [4]int
// a and b are different types
```

3. Arrays **are value types**: assigning one array to another copies all elements

```go
a := [3]int{1,2,3}
b := a
b[0] = 100
fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

> For dynamic-length collections, Go uses **slices** (more flexible than arrays).

---
