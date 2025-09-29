package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    money := 100
    mutex := sync.Mutex{}
    go stingy(&money, &mutex)
    go spendy(&money, &mutex)
    time.Sleep(2 * time.Second) // 确保所有goroutine执行完成
    mutex.Lock()
    fmt.Println("Money in bank account: ", money)
    mutex.Unlock()
}

func stingy(money *int, mutex *sync.Mutex) {
    for i := 0; i < 1000000; i++ {
        // 保护money变量
        mutex.Lock()
        *money += 10
        mutex.Unlock()
    }
    fmt.Println("Stingy Done")
}

func spendy(money *int, mutex *sync.Mutex) {
    for i := 0; i < 1000000; i++ {
        // 保护money变量
        mutex.Lock()
        *money -= 10
        mutex.Unlock()
    }
    fmt.Println("Spendy Done")
}
