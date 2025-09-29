package main

import (
    "fmt"
    "github.com/zqhgithubuser/GoConcurrency/mutexes/code02"
    "sync"
)

func main() {
    wg := sync.WaitGroup{}
    wg.Add(31)
    mutex := sync.Mutex{}
    var frequency = make([]int, 26)
    for i := 1000; i <= 1030; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go func() {
            code02.CountLetters(url, frequency, &mutex)
            defer wg.Done()
        }()
    }
    wg.Wait()

    mutex.Lock()
    for i, c := range code02.AllLetters {
        fmt.Printf("%c-%d ", c, frequency[i])
    }
    mutex.Unlock()
}
