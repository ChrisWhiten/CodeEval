/*
Computing network reach based on levenshtein distance.
TODO: parallelize.
*/
package main

import (
    "fmt"
    "bytes"
    "os"
    "bufio"
    )

var Visited =  make(map[string]int)
var Input []string
var Dictionary []string

func dist1Levenshtein(A string, B string) (bool) {
    if _, ok := Visited[B]; ok {
        // already visited. escape.
        return false
    }

    // check if distance is 1.
    lenA := len(A)
    lenB := len(B)
    check := lenA - lenB
    if check < -1 || check > 1 {
        return false;
    }
    if check == -1 {
        // B is longer. try deleting chars from B until matches A.
        for i := 0; i < lenB; i++ {
            var buffer bytes.Buffer
            for j := 0; j < i; j++ {
                buffer.WriteString(B[j:j+1])
            }
            for j := i + 1; j < lenB; j++ {
                buffer.WriteString(B[j:j+1])
            }
            if buffer.String() == A {
                Visited[B] = 1
                return true
            }
        }
    }
    if check == 1 {
        // A is longer. try deleting chars from A until it matches B.
        for i := 0; i < lenA; i++ {
            var buffer bytes.Buffer
            for j := 0; j < i; j++ {
                buffer.WriteString(A[j:j+1])
            }
            for j := i + 1; j < lenA; j++ {
                buffer.WriteString(A[j:j+1])
            }
            if buffer.String() == B {
                Visited[B] = 1
                return true
            }
        }
    }

    if check == 0 {
        // Same length. try removing char from both until they match.
        for i := 0; i < lenA; i++ {
            var bufferA bytes.Buffer
            var bufferB bytes.Buffer
            for j := 0; j < i; j++ {
                bufferA.WriteString(A[j:j+1])
                bufferB.WriteString(B[j:j+1])
            }
            for j := i + 1; j < lenA; j++ {
                bufferA.WriteString(A[j:j+1])
                bufferB.WriteString(B[j:j+1])
            }
            if bufferA.String() == bufferB.String() {
                Visited[B] = 1
                return true
            }
        }
    }
    return false;
}

func loadData(f string) {
    readingInput := true
    file, _ := os.Open(f)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if readingInput && line == "END OF INPUT" {
            readingInput = false
        } else if readingInput {
            Input = append(Input, line)
        } else {
            Dictionary = append(Dictionary, line)
        }
    }
}

func networkReach(input string) (int) {
    private_count := 0
    for _, elem := range Dictionary {
        if dist1Levenshtein(input, elem) {
            private_count += 1
            private_count += networkReach(elem)
        }
    }
    return private_count
}

func main() {
    loadData(os.Args[1])
    for _, elem := range Input {
        fmt.Println(networkReach(elem))
        Visited = make(map[string]int)
    }
}
