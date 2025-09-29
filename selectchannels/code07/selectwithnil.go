package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    sales := generateAmounts(50)
    expenses := generateAmounts(40)
    endOfDayAmount := 0
    for sales != nil || expenses != nil {
        select {
        case sale, more := <-sales:
            if more {
                fmt.Println("Sale of:", sale)
                endOfDayAmount += sale
            } else {
                sales = nil
            }
        case expense, more := <-expenses:
            if more {
                fmt.Println("Expense of:", expense)
                endOfDayAmount -= expense
            } else {
                expenses = nil
            }
        }
    }
    fmt.Println("End of day profit and loss:", endOfDayAmount)
}

func generateAmounts(n int) <-chan int {
    amounts := make(chan int)
    go func() {
        defer close(amounts)
        for i := 0; i < n; i++ {
            amounts <- rand.Intn(100) + 1
            time.Sleep(100 * time.Millisecond)
        }
    }()
    return amounts
}
