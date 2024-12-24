```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int) // Initialize the map
    m["a"] = 1
    fmt.Println(m["a"])
}
```