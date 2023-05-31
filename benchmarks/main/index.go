package main

import (
   "binarysearch/benchmarks"
   "fmt"
   "os"
   "path"
)

func main() {
	dir, _ := os.Getwd()
	dir = path.Join(dir, "benchmarks")
	benchmarks.Index(dir, benchmarks.Template(os.Stdin))
	fmt.Println("file://" + dir)
}
