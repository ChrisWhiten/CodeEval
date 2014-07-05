package main

import (
    "fmt"
    "bytes"
    "strconv"
    "strings"
    "io/ioutil"
    "os"
    )

func fizzbuzz(A int, B int, N int) (string) {
    if A < 1 || A > 20 {
        return ""
    }
    if (B < 1 || B > 20) {
        return ""
    }
    if (N < 21 || N > 100) {
        return ""
    }

    var buffer bytes.Buffer
    var byteCount int = 0
    for i := 1; i <= N; i++ {
        if i % A == 0 && i % B == 0 {
            bytesWritten, _ := buffer.WriteString("FB ")
            byteCount += bytesWritten
        } else if i % A == 0 {
            bytesWritten, _ := buffer.WriteString("F ")
            byteCount += bytesWritten
        } else if i % B == 0 {
            bytesWritten, _ := buffer.WriteString("B ")
            byteCount += bytesWritten
        } else {
            bytesWritten, _ := buffer.WriteString(strconv.Itoa(i))
            byteCount += bytesWritten
            bytesWritten, _ = buffer.WriteString(" ")
            byteCount += bytesWritten
        }
    }
    // remove the bonus space at the end.
    buffer.Truncate(byteCount - 1)
    return buffer.String()
}
func main() {
    // we can optimize file reads probably. [todo]
    content, _ := ioutil.ReadFile(os.Args[1])
    lines := strings.Split(string(content), "\n")
    for _, l := range lines {
        // empty line gets ignored.
        if len(l) == 0 { continue }
        numbers := strings.Fields(l)
        A, _ := strconv.Atoi(numbers[0])
        B, _ := strconv.Atoi(numbers[1])
        N, _ := strconv.Atoi(numbers[2])
        fmt.Println(fizzbuzz(A, B, N))
    }
}
