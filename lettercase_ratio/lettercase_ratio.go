package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "unicode"
    )

func main() {
    content, _ := ioutil.ReadFile(os.Args[1])
    lines := strings.Split(string(content), "\n")
    for _, l := range lines {
        // empty line gets ignored.
        if len(l) == 0 { continue }
        // now get percentages and format them to 2 decimal places. [todo]
        upper := 0.0
        lower := 0.0
        unknown := 0
        for _, c := range l {
            if unicode.IsLower(c) {
                lower += 1
            } else if unicode.IsUpper(c) {
                upper += 1
            } else {
                unknown += 1
            }
        }
        upper_ratio := (100 * upper)/(upper + lower)
        lower_ratio := (100 * lower)/(upper + lower)
        fmt.Printf("Upper: %.2f Lower: %.2f\n", upper_ratio, lower_ratio)
    }
}
