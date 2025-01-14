```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["a"] = 1

    value, ok := m["b"]
    if ok {
        fmt.Println("Value of b:", value) // This won't print unless "b" is a key
    } else {
        fmt.Println("Key 'b' not found")
    }

    value, ok = m["c"]
    if ok {
        fmt.Println("Value of c:", value)
    } else {
        fmt.Println("Key 'c' not found")
    }
}
```