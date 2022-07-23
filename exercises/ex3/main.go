package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)


func main() {
    flag.Parse()
    root := flag.Arg(0)
    err := filepath.WalkDir(root, visit)
    fmt.Printf("filepath.WalkDir() returned %v\n", err)
}

// walkdir function

func visit(path string, di fs.DirEntry, err error) error { 
    fmt.Printf("Visited: %s\n", path)
    return nil
}

// run : go run main.go dir1

// output : 

// Visited: dir1
// Visited: dir1/dir2
// Visited: dir1/dir2/dir2/file2.txt
// Visited: dir1/file1.txt
// filepath.WalkDir() returned <nil>