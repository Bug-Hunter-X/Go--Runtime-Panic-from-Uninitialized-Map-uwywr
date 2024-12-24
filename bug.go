```go
package main

import "fmt"

func main() {
    var m map[string]int
    m["a"] = 1 // This will cause a runtime panic if m is nil
    fmt.Println(m["a"])
}
```