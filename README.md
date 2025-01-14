# Unexpected Zero Value from Missing Map Keys in Go

This example demonstrates a common issue when working with maps in Go: accessing a non-existent key returns the zero value for the key's type. This might lead to unexpected behavior and hidden bugs if not handled properly.

## The Bug

The provided `bug.go` file shows a simple map access. If you try to access keys that are not present in the map, instead of raising an error, Go returns the zero value for the map's value type (in this case, 0 for `int`). This behavior is not always apparent to newcomers.

## Solution

The solution file, `bugSolution.go`, demonstrates the correct approach.  Always check if a key exists in a map using the `value, ok := map[key]` syntax. The `ok` boolean indicates whether the key was found. This ensures that your code handles missing keys gracefully and avoids potential errors.