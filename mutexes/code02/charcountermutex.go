package code02

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "sync"
    "time"
)

const AllLetters = "abcdefghijklmnopqrstuvwxyz"

func CountLetters(url string, frequency []int, mutex *sync.Mutex) {
    // 并行下载所有文档
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        panic("Server returning error status code: " + resp.Status)
    }

    body, _ := io.ReadAll(resp.Body)

    // 串行处理每个文档
    mutex.Lock()
    for _, b := range body {
        c := strings.ToLower(string(b))
        cIndex := strings.Index(AllLetters, c)
        if cIndex > 0 {
            frequency[cIndex]++
        }
    }
    mutex.Unlock()

    fmt.Println("Completed:", url, time.Now().Format("15:04:05"))
}
