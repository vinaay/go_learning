#for loop
Go only has one looping construct, the for loop. It can be used in various forms: traditional loops, range-based loops (to iterate over slices, maps, and so on), and infinite loops.

The range based loop provides a simplified way for iteration on slices, maps, and others. It makes it easier to access each element directly without needing to manually handle index or length checks. The loop automatically provides both the index and the value during each iteration, improving readability and reducing the chance of off-by-one errors or other indexing issues.



`go run main.go`