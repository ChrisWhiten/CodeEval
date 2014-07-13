package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
    )

func main() {
    file, _ := os.Open(os.Args[1])
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(strings.ToLower(line))
    }
}
