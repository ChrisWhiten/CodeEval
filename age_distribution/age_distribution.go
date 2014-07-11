package main

import (
    "fmt"
    "strconv"
    "os"
    "bufio"
    )


func main() {
    file, _ := os.Open(os.Args[1])
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        number, _ := strconv.Atoi(scanner.Text())
        if number < 0 || number > 100 {
            fmt.Println("This program is for humans")
        } else if number > -1 && number < 3 {
            fmt.Println("Home")
        } else if number > 2 && number < 5 {
            fmt.Println("Preschool")
        } else if number > 4 && number < 12 {
            fmt.Println("Elementary school")
        } else if number > 11 && number < 15 {
            fmt.Println("Middle school")
        } else if number > 14 && number < 19 {
            fmt.Println("High school")
        } else if number > 18 && number < 23 {
            fmt.Println("College")
        } else if number > 22 && number < 66 {
            fmt.Println("Work")
        } else if number > 65 && number < 101 {
            fmt.Println("Retirement")
        }
    }
}

/*

CHALLENGE DESCRIPTION:

In this challenge, we are analyzing the ages of people to determine their status.
If the age is from 0 to 2 the child should be with parents at home, print : 'Home'
If the age is from 3 to 4 the child should visit preschool, print : 'Preschool'
If the age is from 5 to 11 the child should visit elementary school, print : 'Elementary school'
From 12 to 14: 'Middle school'
From 15 to 18: 'High school'
From 19 to 22: 'College'
From 23 to 65: 'Work'
From 66 to 100: 'Retirement'
If the age of the person less than 0 or more than 100 - it might be an alien - print: "This program is for humans"
INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Each line of input contains one integer - age of the person:


1
2
0
19
OUTPUT SAMPLE:

For each line of input print out where the person is:


1
2
Home
College
Submit your solution in a file (some file name).(py2| c| cpp| java| rb| pl| php| tcl| clj| js| scala| cs| m| py3| hs| go| bash| lua) or use the online editor.

*/
