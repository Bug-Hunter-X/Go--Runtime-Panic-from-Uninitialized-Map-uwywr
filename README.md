# Go: Uninitialized Map Access

This repository demonstrates a common runtime error in Go: accessing an uninitialized map.  Attempting to assign or read from an uninitialized map will result in a runtime panic. The `bug.go` file contains the erroneous code, and `bugSolution.go` shows the corrected version.

## Bug

The `bug.go` file contains code that attempts to access a map before it's initialized. This leads to a panic because the map is `nil`. 

## Solution

The `bugSolution.go` file shows the corrected code. We explicitly initialize the map before accessing it, preventing the panic.  Always ensure maps are initialized before use.