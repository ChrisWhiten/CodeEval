package main

import (
    "fmt"
    "strconv"
    "os"
    "bufio"
    )


func main() {
    running_sum := 0

    file, _ := os.Open(os.Args[1])
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        number, _ := strconv.Atoi(scanner.Text())
        running_sum += number
    }
    fmt.Print(running_sum)
}
