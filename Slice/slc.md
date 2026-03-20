In **Go**, a **slice** is a **dynamic, flexible view into an array**. Unlike arrays, slices **can grow and shrink**, making them much more commonly used in Go programs. You can think of a slice as a **lightweight wrapper around an array**.

---

## 🔹 Key Points About Slices

1. Slices are **references to an underlying array**.
2. Slices **do not store data themselves**; they describe a segment of an array.
3. Length and capacity are important:

   * `len(slice)` → number of elements in the slice
   * `cap(slice)` → maximum elements before it needs to allocate a new array

---

## 🔹 Creating Slices

### 1. From an array

```go id="sl1"
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4] // slice from index 1 to 3
fmt.Println(s) // [2 3 4]
```

### 2. Using slice literal

```go id="sl2"
s := []int{10, 20, 30}
fmt.Println(s) // [10 20 30]
```

### 3. Using `make()`

```go id="sl3"
s := make([]int, 3, 5) // length 3, capacity 5
fmt.Println(s)           // [0 0 0]
fmt.Println(len(s), cap(s)) // 3 5
```

---

## 🔹 Appending to a Slice

```go id="sl4"
s := []int{1, 2, 3}
s = append(s, 4, 5)
fmt.Println(s) // [1 2 3 4 5]
```

> `append()` may **allocate a new underlying array** if the capacity is exceeded.

---

## 🔹 Looping Through Slices

```go id="sl5"
s := []string{"Go", "Python", "Java"}

for i := 0; i < len(s); i++ {
    fmt.Println(i, s[i])
}

// Using range
for i, v := range s {
    fmt.Println(i, v)
}
```

---

## 🔹 Slices Are Reference Types

```go id="sl6"
arr := [5]int{1, 2, 3, 4, 5}
s1 := arr[1:4] // [2 3 4]
s2 := s1[1:3]  // [3 4]

s2[0] = 100
fmt.Println(s1) // [2 100 4]
fmt.Println(arr) // [1 2 100 4 5]
```

> Changing a slice **changes the underlying array**, affecting other slices pointing to it.

---

## 🔹 Multidimensional Slice

```go id="sl7"
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
fmt.Println(matrix[1][2]) // 6
```

---

## 🔹 Slice vs Array

| Feature        | Array      | Slice                  |
| -------------- | ---------- | ---------------------- |
| Size           | Fixed      | Dynamic                |
| Type           | `[n]T`     | `[]T`                  |
| Reference      | Value type | Reference type         |
| Built-in funcs | len() only | len(), cap(), append() |

---

Slices are **used almost everywhere in Go**, and they’re often combined with **structs, functions, and pointers** for real-world programs.

---

