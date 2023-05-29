package main

import (
   "fmt"
   "os"
   "path"
   "trees/benchmarks"
)

func main() {
	dir, _ := os.Getwd()
	dir = path.Join(dir, "benchmarks")
	benchmarks.Index(dir, benchmarks.Template(os.Stdin))
	fmt.Println("file://" + dir)
}
