package main

import (
    "fmt"
    "sync"
)

//func main() {
//    chA := make(chan int)
//    chB := make(chan int)
//    go func() {
//        for i := 0; i<=100; i+=2 {
//            <- chA
//            fmt.Println("thread A :", i)
//            chB <- 1
//        }
//    }()
//    go func() {
//        for i := 1; i<=100; i+=2 {
//            <- chB
//            fmt.Println("thread B :", i)
//            chA <- 1
//        }
//    }()
//    chA <- 1
//    time.Sleep(100*time.Second)
//
//}


func main() {
    a := sync.WaitGroup{}
    a.Add(2)
    ch := make(chan int)

    go func() {
        for i := 1; i <= 100; i += 1 {
            <-ch
            if i%2 == 0 {
                fmt.Println(i)
            }
        }
        a.Done()
    }()

    go func() {
        for i := 1; i <= 100; i += 1 {
            ch <- 0
            if i%2 == 1 {
                fmt.Println(i)
            }
        }
        a.Done()
    }()

    a.Wait()
    // 死等
    // select {}
}