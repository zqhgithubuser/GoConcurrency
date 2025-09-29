package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "sync"
)

func main() {
    wg := sync.WaitGroup{}
    wg.Add(1)
    go fileSearch(os.Args[1], os.Args[2], &wg)
    wg.Wait()
}

func fileSearch(dir string, filename string, wg *sync.WaitGroup) {
    files, _ := os.ReadDir(dir)
    for _, file := range files {
        filePath := filepath.Join(dir, file.Name())
        if strings.Contains(file.Name(), filename) {
            fmt.Println(filePath)
        }
        if file.IsDir() {
            wg.Add(1)
            go fileSearch(filePath, filename, wg)
        }
    }
    wg.Done()
}
